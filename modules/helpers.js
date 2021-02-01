const fs = require("fs"),
    screen = require("../config.json").screen ? require('./oled') : false,
    spawn = require('child_process').spawn,
    {exec} = require("child_process"),
    ffmpeg = require("fluent-ffmpeg");

module.exports = {
    save: async function (setting, value) {
        const config = fs.readFileSync('./config.json', 'utf-8');
        let data = JSON.parse(config.toString());
        data[setting] = value;
        fs.writeFileSync('./config.json', JSON.stringify(data, null, 4), 'utf-8');
        if (config.screen) new screen().miniMessage("OK!");
    },
    killPiFmADV: function () {
        return new Promise((resolve) => {
            exec(`sudo pkill -2 pi_fm_adv`, () => resolve(true));
        });
    },
    playPiFmADV: async function (config, song = false) {
        await this.killPiFmADV();
        const options = [
            'core/pi_fm_adv',
            `--ps "${config.PS}"`,
            `--rt "${config.RT}"`,
            '--freq', freq,
            '--power', config.power];
        if (song) options.push('--audio', `"./music/${song}.wav"`);
        const execWithStd = spawn(`sudo`, options, {shell: true});
        //In case of debugging, you can uncomment this safely:
        //execWithStd.stdout.on('data', function (data) { console.log('stdout: ' + data.toString()); });
        execWithStd.stderr.on('data', function (data) {
            console.error('stderr: ' + data.toString());
        });
    },
    convertToWav: function (song) {
        return new Promise(async (resolve, reject) => {
            ffmpeg(`./ytdl-temp/${song}`)
                .toFormat('wav')
                .on('start', () => ffmpegorytdlWorking = true)
                .on('error', (err) => {
                    console.error("FFmpeg returned error: " + err.message);
                    reject(err.message);
                })
                //.on('progress', (progress) => console.log('Processing: ' + progress.percent)) //Uncomment for fun!
                .on('end', () => {
                    ffmpegorytdlWorking = false;
                    resolve(true);
                    exec('sudo rm -r ./ytdl-temp/*');
                })
                .save(`./music/${song}.wav`);
        });
    }
};