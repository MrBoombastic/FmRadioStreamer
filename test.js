
const fs = require('fs')
fs.readFile('./config.json', 'utf8', (err, data) => {
	if (err) throw err;
	data = JSON.parse(data)
	data.freq = 100
	console.log(data)
	fs.writeFile('./config.json', JSON.stringify(data), (err) => {
		if (err) throw err;
		oled.setCursor(1, 40);
		oled.writeString(font, 2, "OK");
		setTimeout(function () { updateScreen() }, 2000)
	})
})