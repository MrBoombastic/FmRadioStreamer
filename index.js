process.title = "fmradiostreamer";
const config = require("./config.json");
global.multiplier = 0.1;
global.freq = Math.round(config.freq * 10) / 10;
global.ffmpegorytdlWorking = false;
global.helpers = require("./modules/helpers")
const oled = config.screen ? require('./modules/oled') : false;
const led = require('./modules/led');
const webserver = require('./modules/webserver');
const buttons = require('./modules/buttons');
const {exec} = require("child_process");
new led().run();
new led().workingLed();
new buttons().run();
new webserver().run();
if (config.screen) new oled().run();
process.on('SIGINT', () => {
    if (config.screen) new oled().stop();
    new buttons().unexport();
    new led().unexport();
    exec(`sudo pkill -2 pi_fm_adv`);
    process.exit();
});

let humanConfig = "";
for (const key in config) humanConfig += `     ${key}: ${config[key]}\n`;

console.log(`FmRadioStreamer working!\nUsing configuration:\n${humanConfig.slice(0, -1)}`);

