process.title = "fmradiostreamer";
const config = require("./config.json");
global.multiplier = 0.1;
global.freq = Math.round(config.freq * 10) / 10;
global.ffmpegorytdlWorking = false;
let oled;
try {
    oled = require('./modules/oled');
    global.oledNotSupported = false;
} catch (e) {
    console.log("OLED not present!");
    global.oledNotSupported = true;
}
const led = require('./modules/led');
const webserver = require('./modules/webserver');
const buttons = require('./modules/buttons');
const {exec} = require("child_process");
new led().run();
new buttons().run();
new webserver().run();
if (!global.oledNotSupported) new oled().run();
process.on('SIGINT', () => {
    if (!global.oledNotSupported) new oled().stop();
    new buttons().unexport();
    new led().unexport();
    exec(`sudo pkill -2 pi_fm_adv`);
    process.exit();
});