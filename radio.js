process.title = "fmradiostreamer"
const config = require("./config.json")
global.multiplier = 0.1
global.freq = Number(config.freq).toFixed(1)
global.ffmpegorytdlWorking = false
const ledLoop = require('./modules/ledLoop');
const webserver = require('./modules/webserver');
const oled = require('./modules/oled')
const buttons = require('./modules/buttons')
const { exec } = require("child_process");
new oled().run()
new oled().updateScreen()
new ledLoop().stop();
new ledLoop().run();
new buttons().run()
new webserver().run()
process.on('SIGINT', () => {
	new oled().stop()
	new ledLoop().stop()
	new buttons().stop()
	new ledLoop().unexport()
	exec(`sudo pkill -2 pi_fm_adv`)
	process.exit();
})