const fs = require("fs");
let screen;
if (!oledNotSupported) screen = require("./oled");

module.exports = {
    save: async function (setting, value) {
        const config = fs.readFileSync('./config.json', 'utf-8');
        let data = JSON.parse(config.toString());
        data[setting] = value;
        fs.writeFileSync('./config.json', JSON.stringify(data, null, 4), 'utf-8');
        if (!oledNotSupported) new screen().miniMessage("OK!");
    }
};