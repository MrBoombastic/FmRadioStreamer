const i2c = require('i2c-bus')
const i2cBus = i2c.openSync(1)
let oled = require('oled-i2c-bus');
const opts = {width: 128, height: 64, address: 0x3C};
oled = new oled(i2cBus, opts);
const font = require('oled-font-5x7');
const config = require("../config.json")
let RT = config.RT
function insert(str, index, value) {
    return str.substr(0, index) + value + str.substr(index);
}
if (RT.length <= 19) { } //for code readability
else if (RT.length > 20 && RT.length <= 39) { RT = insert(RT, 20, "\n") }
else if (RT.length > 39 && RT.length <= 58) { RT = insert(RT, 20, "\n"); RT = insert(RT, 41, "\n") }
else { RT = insert(RT, 20, "\n"); RT = insert(RT, 41, "\n"); RT = insert(RT, 62, "\n") }

module.exports = class screen {
    constructor() {
        this.run = async function () {
            oled.clearDisplay();
            oled.turnOnDisplay();
        }
        this.updateScreen = function () {
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
            this.access = function () {
                return oled;
            }
            this.stop = function () {
                oled.update()
                oled.clearDisplay();
                oled.turnOffDisplay();
            }
            this.miniMessage = function (message) {
                oled.setCursor(1, 40);
                oled.writeString(font, 2, message);
                setTimeout(function () {
                    new screen().updateScreen()
                }, 2000)
            }
        }
}