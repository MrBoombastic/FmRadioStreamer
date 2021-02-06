const fs = require("fs"),
    config = require("../config.json"),
    screen = require("../config.json").screen ? require('./oled') : false,
    spawn = require('child_process').spawn,
    {exec} = require("child_process"),
    ffmpeg = require("fluent-ffmpeg"),
    fetch = require("node-fetch"),
    ytdl = require('ytdl-core');
let localAddress = "";
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
    playPiFmADV: async function (song = false) {
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
                .on('progress', async (progress) => new screen().miniMessage(Math.round(progress.percent).toString(), true))
                .on('end', () => {
                    ffmpegorytdlWorking = false;
                    resolve(true);
                    exec('sudo rm -r ./ytdl-temp/*');
                })
                .save(`./music/${song}.wav`);
        });
    },
    ytSearch: async function (name) {
        let response = await fetch(`https://youtube.googleapis.com/youtube/v3/search?key=${config.apikey}&q=${encodeURIComponent(name)}&part=snippet&maxResults=1&type=video`);
        return await response.json();
    },
    getYT: function (song, searchOnly = false) {
        return new Promise(async (resolve, reject) => {
            const result = await this.ytSearch(song);
            if (!result || !result.items[0] || result.error) return reject("Error from YouTube. Possibly no results or being ratelimited!");
            const music = result.items[0].snippet;
            music.id = result.items[0].id.videoId;
            music.url = "https://youtu.be/" + result.items[0].id.videoId;
            if (searchOnly) return resolve(music);
            ffmpegorytdlWorking = true;
            const safeSongName = music.title.replace(/[^a-zA-Z0-9\s]/g, "").replace(/ /g, "-");
            const video = ytdl(music.url, {quality: "highestaudio", filter: "audioonly"});
            video.pipe(fs.createWriteStream(`ytdl-temp/${safeSongName}`));
            video.on('end', async () => {
                const conversion = await this.convertToWav(safeSongName);
                if (conversion) resolve(safeSongName);
                else reject(conversion);
                ffmpegorytdlWorking = false;
            });
        });
    },
    getWebserverAddr: function () {
        return localAddress;
    },
    fetchWebserverAddr: function () {
        return new Promise(async (resolve) => {
            exec("hostname -I | awk '{print $1}'", {shell: true})
                .stdout.on('data', (data) => {
                    const address = data.replace(/\n|\r/g, "") + ":" + config.port;
                    resolve(address);
                    localAddress = address;
                }
            );
        });
    }
};