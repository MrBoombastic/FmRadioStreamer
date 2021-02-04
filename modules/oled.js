const i2c = require('i2c-bus');
const i2cBus = i2c.openSync(1);
const bus = require('oled-i2c-bus');
const oled = new bus(i2cBus, {width: 128, height: 64, address: 0x3C});
const font = require('oled-font-5x7');
const config = require("../config.json");
let RT = config.RT;

module.exports = class screen {
    constructor() {
        this.run = async function () {
            oled.clearDisplay();
            oled.turnOnDisplay();
            await this.updateScreen();
        };
        this.updateScreen = async function () {
            oled.clearDisplay();
            oled.setCursor(100, 1);
            oled.writeString(font, 1, "x" + multiplier.toString());
            oled.setCursor(1, 1);
            oled.writeString(font, 1, config.PS);
            oled.setCursor(1, 15);
            oled.writeString(font, 1, RT, 1, true);
            oled.setCursor(40, 37);
            oled.writeString(font, 2, (Math.round(freq * 10) / 10).toFixed(1) + " FM");
            oled.setCursor(1, 57);
            oled.writeString(font, 1, await helpers.getWebserverAddr());
        };
        this.stop = function () {
            oled.update();
            oled.clearDisplay();
            oled.turnOffDisplay();
        };
        this.miniMessage = function (message, live = true) {
            if (message === "100") message = " "; //assuming this message is from FFmpeg
            oled.setCursor(1, 37);
            oled.writeString(font, 2, message);
            if (!live) setTimeout( function () {
                new screen().updateScreen();
            }, 2000);
        };
    }
};