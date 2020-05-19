const Gpio = require('onoff').Gpio;
const helpers = require('./helpers');
const buttonLow = new Gpio(21, 'in', 'rising', {debounceTimeout: 40});
const buttonHigh = new Gpio(20, 'in', 'rising', {debounceTimeout: 40});
const buttonSet = new Gpio(16, 'in', 'rising', {debounceTimeout: 40});
const buttonMultiplier = new Gpio(12, 'in', 'rising', {debounceTimeout: 40});
const led = require("./led");
let screen
if(!oledNotSupported) { screen = require("./oled"); }


module.exports = class buttons {
    constructor() {
        this.run = async function () {
            buttonLow.watch(async (err) => {
                if (err) throw err;
                if (Number(freq) - multiplier <= 87.2) {
                    if(!oledNotSupported) { new screen().miniMessage("MIN"); }
                    await new led().ledWarnBlink();
                } else {
                    freq = Number(freq) - multiplier;
                    if(!oledNotSupported) { new screen().updateScreen(); }
                }
            });
            buttonHigh.watch(async (err) => {
                if (err) throw err;
                if (Number(freq) + multiplier >= 108.9) {
                    if(!oledNotSupported) { new screen().miniMessage("MAX"); }
                    await new led().ledWarnBlink();
                } else {
                    freq = Number(freq) + multiplier;
                    if(!oledNotSupported) { new screen().updateScreen(); }
                }
            });
            buttonSet.watch(async (err) => {
                if (err) throw err;
                await helpers.save('freq', Number((Math.round(freq * 10) / 10).toFixed(1)));
            });
            buttonMultiplier.watch(async (err) => {
                if (err) throw err;
                if (multiplier === 0.1) multiplier = 0.5;
                else if (multiplier === 0.5) multiplier = 1;
                else if (multiplier === 1) multiplier = 2;
                else if (multiplier === 2) multiplier = 5;
                else multiplier = 0.1;
                if(!oledNotSupported) { new screen().updateScreen(); }
            });
        };
        this.unexport = function () {
            buttonLow.unexport();
            buttonHigh.unexport();
            buttonMultiplier.unexport();
            buttonSet.unexport();
        };
    }
};