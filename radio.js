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
	exec(`sudo pkill -2 pi_fm_adv`)
	process.exit();
})
process.title = "fmradiostreamer"
const config = require("./config.json")
const { YTSearcher } = require('ytsearcher');
const ytsearcher = new YTSearcher(config.apikey);
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
var spawn = require('child_process').spawn
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
const ytdl = require('ytdl-core');
const { Progress } = require('express-progressbar');
function timeString2ms(a, b, c) {// time(HH:MM:SS.mss)
	return c = 0,
		a = a.split('.'),
		!a[1] || (c += a[1] * 1),
		a = a[0].split(':'), b = a.length,
		c += (b == 3 ? a[0] * 3600 + a[1] * 60 + a[2] * 1 : b == 2 ? a[0] * 60 + a[1] * 1 : s = a[0] * 1) * 1e3,
		c
}
const sleeptime = 500
function sleep(time) {
	return new Promise((resolve) => {
		setTimeout(resolve, time)
	})
}
oled = new oled(i2cBus, opts);
const font = require('oled-font-5x7');
let freq = Number(config.freq).toFixed(1)
let multiplier = 0.1
function save(setting, value) {
	fs.readFile('./config.json', 'utf8', (err, data) => {
		if (err) throw err;
		data = JSON.parse(data)
		data[setting] = Math.round(value)
		fs.writeFile('./config.json', JSON.stringify(data, null, 4), (err) => {
			if (err) throw err;
			oled.setCursor(1, 40);
			oled.writeString(font, 2, "OK");
			setTimeout(function () { updateScreen() }, 2000)
		})
	})
}
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
	save('freq', Math.round(freq))
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
app.get(prefix + "yt/:song", (req, res) => {
	const safeSongName = req.params.song.replace(/[^a-zA-Z0-9\s]/g, "").replace(/ /g, "-")
	ytsearcher.search(req.params.song, { type: 'video' })
		.then(result => {
			const music = result.first
			ytdl.getInfo(music.id, (err) => {
				if (err) throw err;
				const p = new Progress(res);
				const video = ytdl(music.url, { format: 'aac' })
				video.pipe(fs.createWriteStream(`./music/${safeSongName}.aac`))
				video.on("progress", (chunkLength, downloaded, total) => {
					const percent = downloaded / total * 50;
					p.update(percent, { stage: 'downloading' });
				});
				video.on('end', () => {
					p.update(100, {stage: 'done'})
				})
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
	exec(`sudo pkill -2 pi_fm_rds`, () => {
		exec(`mkfifo rds_ctl`)
		const execWithStd = spawn('sudo', [`sox -t aac ../music/${req.params.song}.aac -t wav -  | sudo ./pi_fm_rds`,
			"-ps", config.PS,
			"-rt", config.RT,
			"-freq", config.freq,
			"-audio", "-",
			"-ctl", "rds_ctl"])
		//In case of debugging, you can uncomment this safely:
		//execWithStd.stdout.on('data', function (data) { console.log('stdout: ' + data.toString()); });
		execWithStd.stderr.on('data', function (data) { console.log('stderr: ' + data.toString()); });
		res.send("ok")
	})
})
app.get(prefix + "change/:setting/:value", (req, res) => {
	const setting = req.params.setting.toString()
	const value = req.params.value.toString()
	if (setting !== "PS" && setting !== "RT") return res.status(405).send("Allowed settings: PT, RT, TA, PTY.")
	else {save(setting, value);res.send(setting.toString() + ", " + value.toString());}
})
/*
app.get(prefix + "set/:setting/:value", (req, res) => {
	if (setting != "PS" && setting != "RT" && setting != "TA" && setting != "PTY") return res.status(405).send("Allowed settings: PT, RT, TA, PTY.")
	else {save(setting, value);res.send(setting.toString() + ", " + value.toString())}
})*/
app.listen(config.port, () => console.log(`FmRadioStreamer working!\nPort: ${config.port}\nFreq: ${Number(config.freq).toFixed(1)}\nPS: ${config.PS}\nRT: ${config.RT}\nTA: ${config.TA}\nPTY: ${config.PTY}`))