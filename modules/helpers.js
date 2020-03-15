const fs = require("fs")
const screen = require("./oled")

module.exports = {
    save: async function (setting, value) {
        const config = fs.readFileSync('./config.json', 'utf-8')
        let data = JSON.parse(config.toString())
        data[setting] = Math.round(value)
        fs.writeFileSync('./config.json', JSON.stringify(data, null, 4), 'utf-8')
        new screen().miniMessage("OK!")
    }
}