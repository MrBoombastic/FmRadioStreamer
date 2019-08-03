process.on('SIGINT', () => {
	oled.update()
	oled.clearDisplay();
	oled.turnOffDisplay();
	allLightsOff();
	buttonLow.unexport();
	buttonHigh.unexport();
	led1.unexport();
	led2.unexport();
	led3.unexport();
	led4.unexport();
	ledWarn.unexport();
	process.exit();
})
process.title = "fmradiostreamer"
const config = require("./config.json")
const { YTSearcher } = require('ytsearcher');
const ytsearcher = new YTSearcher(config.apikey);
const ytdl = require('ytdl-core');
const fs = require('fs')
const express = require("express");
const app = express();
const i2c = require('i2c-bus-i2c-1')
const i2cBus = i2c.openSync(0)
let oled = require('oled-i2c-bus');
const SIZE_X = 128,
	SIZE_Y = 64;
const opts = {
	width: SIZE_X,
	height: SIZE_Y,
	address: 0x3C
};
const { exec } = require("child_process");

const Gpio = require('onoff').Gpio;
const led4 = new Gpio(26, 'out');
const led3 = new Gpio(13, 'out');
const led2 = new Gpio(6, 'out');
const led1 = new Gpio(19, 'out');
const ledWarn = new Gpio(5, 'out');
const buttonLow = new Gpio(21, 'in', 'rising', { debounceTimeout: 40 });
const buttonHigh = new Gpio(20, 'in', 'rising', { debounceTimeout: 40 });
const buttonSet = new Gpio(16, 'in', 'rising', { debounceTimeout: 40 });
const buttonMultiplier = new Gpio(12, 'in', 'rising', { debounceTimeout: 40 });

const sleeptime = 500
const sleep = (howLong) => {
	return new Promise((resolve) => {
		setTimeout(resolve, howLong)
	})
}
oled = new oled(i2cBus, opts);
const font = require('oled-font-5x7');
let freq = Number(config.freq).toFixed(1)
let multiplier = 0.1

async function runLights() {
	while (true) {
		// led1
		led4.writeSync(0)
		led1.writeSync(1)
		await sleep(sleeptime)
		// led 2
		led1.writeSync(0)
		led2.writeSync(1)
		await sleep(sleeptime)
		// led3
		led2.writeSync(0)
		led3.writeSync(1)
		await sleep(sleeptime)
		// led4
		led3.writeSync(0)
		led4.writeSync(1)
		await sleep(sleeptime)
	}
}
async function allLightsOff() {
	led1.writeSync(0)
	led2.writeSync(0)
	led3.writeSync(0)
	led4.writeSync(0)
	ledWarn.writeSync(0);
}
oled.clearDisplay();
oled.turnOnDisplay();
function insert(str, index, value) {
	return str.substr(0, index) + value + str.substr(index);
}
let RT = config.RT
if (RT.length <= 19) { }
else if (RT.length > 20 && RT.length <= 39) { RT = insert(RT, 20, "\n") }
else if (RT.length > 39 && RT.length <= 58) { RT = insert(RT, 20, "\n"); RT = insert(RT, 41, "\n") }
else { RT = insert(RT, 20, "\n"); RT = insert(RT, 41, "\n"); RT = insert(RT, 62, "\n") }
function updateScreen() {
	oled.clearDisplay()
	oled.setCursor(100, 1);
	oled.writeString(font, 1, "x" + multiplier.toString());
	oled.setCursor(1, 1);
	oled.writeString(font, 1, config.PS);
	oled.setCursor(1, 15);
	oled.writeString(font, 1, RT);
	oled.setCursor(40, 40);
	oled.writeString(font, 2, Number(freq).toFixed(1) + " FM");
}
updateScreen()
allLightsOff()
runLights()
buttonLow.watch(async (err) => {
	if (err) throw err;
	if (Number(freq) - multiplier <= 87.2) {
		ledWarn.writeSync(0)
		ledWarn.writeSync(1)
		await sleep(sleeptime)
		ledWarn.writeSync(0)
		return;
	} else { freq = Number(freq) - multiplier; updateScreen() }
})
buttonHigh.watch(async (err) => {
	if (err) throw err;
	if (Number(freq) + multiplier >= 108.9) {
		ledWarn.writeSync(0)
		ledWarn.writeSync(1)
		await sleep(sleeptime)
		ledWarn.writeSync(0)
		return;
	} else { freq = Number(freq) + multiplier; updateScreen() }
})
buttonSet.watch(async (err) => {
	if (err) throw err;
	fs.readFile('./config.json', 'utf8', (err, data) => {
		if (err) throw err;
		data = JSON.parse(data)
		data.freq = Math.round(freq)
		fs.writeFile('./config.json', JSON.stringify(data), (err) => {
			if (err) throw err;
			oled.setCursor(1, 40);
			oled.writeString(font, 2, "OK");
			setTimeout(function () { updateScreen() }, 2000)
		})
	})
})
buttonMultiplier.watch(async (err) => {
	if (err) throw err;
	if (multiplier == 0.1) multiplier = 0.5
	else if (multiplier == 0.5) multiplier = 1
	else if (multiplier == 1) multiplier = 2
	else if (multiplier == 2) multiplier = 5
	else multiplier = 0.1
	updateScreen()
})


const prefix = "/fmradiostreamer/"
app.post(prefix + "yt/:song", (req, res) => {
	ytsearcher.search(req.params.song, { type: 'video' })
		.then(result => {
			const music = result.first
			ytdl.getInfo(music.id, (err, info) => {
				if (err) throw err;
				const audioFormats = ytdl.filterFormats(info.formats, 'audioonly');
				// eslint-disable-next-line no-unused-vars
				ytdl(music.url, { filter: (format) => audioFormats[0] }).pipe(fs.createWriteStream(`./music/${search}.wav`)).then(res.send('ok'))
			});
		});
})
app.get(prefix + "list", (req, res) => {
	let finallist = ""
	const musiclist = fs.readdirSync('./music/')
	musiclist.forEach(function (song) {
		finallist = finallist + song + ", "
	})
	res.json(JSON.parse(`{"list": "${finallist.slice(0, -2)}"}`))
})
app.get(prefix + "play/:song", (req, res) => {
	exec(`sudo pkill -2 pi_fm_adv`, () => {
		exec(`sudo ./radioPlay.sh '${config.PS}' '${config.RT}' ${config.freq} '${req.params.song}'`)
		res.send("ok")
	})
})
app.listen(1080, () => console.log(`FmRadioStreamer working!\nPort: 1080\nFreq: ${Number(config.freq).toFixed(1)}\nPS: ${config.PS}\nRT: ${config.RT}`))