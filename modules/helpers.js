const fs = require("fs"),
    screen = require("../config.json").screen ? require('./oled') : false,
    spawn = require('child_process').spawn,
    {exec} = require("child_process"),
    ffmpeg = require("fluent-ffmpeg"),
    {YTSearcher} = require('ytsearcher'),
    ytsearcher = new YTSearcher(require("../config.json").apikey),
    ytdl = require('ytdl-core');

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
        const PiFmAdv = spawn(`sudo`, options, {shell: true});
        //In case of debugging, you can uncomment this safely:
        //PiFmAdv.stdout.on('data', function (data) { console.log('stdout: ' + data.toString()); });
        PiFmAdv.stderr.on('data', (data) => console.error('stderr: ' + data.toString()));
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
                .on('progress', (progress) => new screen().miniMessage(Math.round(progress.percent).toString(), true))
                .on('end', () => {
                    ffmpegorytdlWorking = false;
                    resolve(true);
                    exec('sudo rm -r ./ytdl-temp/*');
                })
                .save(`./music/${song}.wav`);
        });
    },
    getYT: async function (song, searchOnly = false) {
        //return console.log(screen)
        return new Promise(async (resolve, reject) => {
            ffmpegorytdlWorking = true;
            const result = await ytsearcher.search(song, {type: 'video'});
            const music = await result.first;
            if (searchOnly) {
                ffmpegorytdlWorking = false;
                return resolve(music);
            }
            const safeSongName = music.title.replace(/[^a-zA-Z0-9\s]/g, "").replace(/ /g, "-");
            const video = ytdl(music.url, {quality: "highestaudio", filter: "audioonly"});
            video.pipe(fs.createWriteStream(`ytdl-temp/${safeSongName}`));
            video.on('end', async () => {
                const conversion = await this.convertToWav(safeSongName);
                if (conversion) resolve(safeSongName);
                else return reject(conversion);
            });
        });
    }
};