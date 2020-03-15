const Gpio = require('onoff').Gpio;
const helpers = require('./helpers')
const buttonLow = new Gpio(21, 'in', 'rising', {debounceTimeout: 40});
const buttonHigh = new Gpio(20, 'in', 'rising', {debounceTimeout: 40});
const buttonSet = new Gpio(16, 'in', 'rising', {debounceTimeout: 40});
const buttonMultiplier = new Gpio(12, 'in', 'rising', {debounceTimeout: 40});
const ledLoop = require("./ledLoop")
const screen = require("./oled")


module.exports = class buttons {
    constructor() {
        this.run = async function () {
            buttonLow.watch(async (err) => {
                if (err) throw err;
                if (Number(freq) - multiplier <= 87.2) {
                    new screen().miniMessage("MIN")
                    await new ledLoop().ledWarnBlink()
                } else {
                    freq = Number(freq) - multiplier;
                    new screen().updateScreen()
                }
            })
            buttonHigh.watch(async (err) => {
                if (err) throw err;
                if (Number(freq) + multiplier >= 108.9) {
                    new screen().miniMessage("MAX")
                    await new ledLoop().ledWarnBlink()
                } else {
                    freq = Number(freq) + multiplier;
                    new screen().updateScreen()
                }
            })
            buttonSet.watch(async (err) => {
                if (err) throw err;
                await helpers.save('freq', Math.round(freq))
            })
            buttonMultiplier.watch(async (err) => {
                if (err) throw err;
                if (multiplier === 0.1) multiplier = 0.5
                else if (multiplier === 0.5) multiplier = 1
                else if (multiplier === 1) multiplier = 2
                else if (multiplier === 2) multiplier = 5
                else multiplier = 0.1
                new screen().updateScreen()
            })
        }
        this.stop = function () {
            buttonLow.unexport();
            buttonHigh.unexport();
            buttonMultiplier.unexport()
            buttonSet.unexport()
        }
    }
}