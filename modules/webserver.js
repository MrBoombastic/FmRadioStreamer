const express = require("express"),
    app = express(),
    {YTSearcher} = require('ytsearcher'),
    config = require("../config.json"),
    ytsearcher = new YTSearcher(config.apikey),
    ytdl = require('ytdl-core'),
    fs = require('fs'),
    spawn = require('child_process').spawn,
    {exec} = require("child_process"),
    led = require('./led');

app.set('view engine', 'ejs');

module.exports = class webserver {
    constructor() {
        this.run = async function () {
            app.get("/", (req, res) => {
                res.render('dash.ejs', {list: fs.readdirSync('./music/')});
            });

            app.get("/yt/:song", async (req, res) => {
                ffmpegorytdlWorking = true;
                const result = await ytsearcher.search(req.params.song, {type: 'video'});
                const music = await result.first;
                const safeSongName = music.title.replace(/[^a-zA-Z0-9\s]/g, "").replace(/ /g, "-");
                const video = ytdl(music.url, {quality: "highestaudio", filter: "audioonly"});
                video.pipe(fs.createWriteStream(`ytdl-temp/${safeSongName}`));
                video.on('end', () => {
                    new led().workingLed();
                    const ffmpegProgress = spawn('ffmpeg', ['-i', `ytdl-temp/${safeSongName}`, `music/${safeSongName}.wav`]);
                    //In case of debugging, uncomment this. Why FFmpeg produces data on stderr? idk
                    //ffmpegProgress.stdout.on('data', function (data) { console.log('stdout: ' + data.toString()); });
                    //ffmpegProgress.stderr.on('data', function (data) {console.log('stderr or ffmpeg: ' + data.toString());});
                    ffmpegProgress.on('exit', function (code) {
                        ffmpegorytdlWorking = false;
                        if (code !== 0) {
                            console.error("FFmpeg returned error: " + code);
                            return res.status(500).json({
                                status: "error",
                                desc: "FFmpeg returned error: " + code
                            });
                        }
                        exec('sudo rm -r ./ytdl-temp/*');
                        res.json({
                            status: "done",
                            desc: safeSongName
                        });
                    });
                });
            });

            app.get("/list", (req, res) => {
                let finallist = "";
                const musiclist = fs.readdirSync('./music/');
                musiclist.forEach(song => finallist += (song + ", "));
                res.json({status: "done", desc: finallist.slice(0, -2)});
            });
            app.get("/play", (req, res) => {
                if (!req.query.song) return res.status(404).json({status: "failed", desc: "not found"});
                exec(`sudo pkill -2 pi_fm_adv`, () => {
                    exec(`mkfifo rds_ctl`);
                    const execWithStd = spawn(`sudo`, [
                        'core/pi_fm_adv',
                        `--ps "${config.PS}"`,
                        `--rt "${config.RT}"`,
                        '--freq', freq,
                        '--ctl', 'rds_ctl',
                        '--power', config.power,
                        '--audio', `"./music/${req.query.song}.wav"`], {shell: true});
                    //In case of debugging, you can uncomment this safely:
                    //execWithStd.stdout.on('data', function (data) { console.log('stdout: ' + data.toString()); });
                    execWithStd.stderr.on('data', function (data) {
                        console.log('stderr: ' + data.toString());
                    });
                    res.send("ok");
                });
            });
            /*
            app.get("change/:setting/:value", (req, res) => {
                const setting = req.params.setting.toString();
                const value = req.params.value.toString();
                if (setting !== "PS" && setting !== "RT") return res.status(405).send("Allowed settings: PT, RT, TA, PTY.");
                else {
                    helpers.save(setting, value);
                    res.send(setting.toString() + ", " + value.toString());
                }
            });*/
            exec(`sudo pkill -2 pi_fm_adv`, () => {
                spawn(`sudo`, [
                        'core/pi_fm_adv',
                        `--ps "${config.PS}"`,
                        `--rt "${config.RT}"`,
                        '--freq', freq,
                        '--ctl', 'rds_ctl',
                        '--power', config.power
                    ],
                    {shell: true});
            });
            let humanConfig = "";
            for (const key in config) humanConfig += `     ${key}: ${config[key]}\n`;

            app.listen(config.port, () => console.log(`FmRadioStreamer working!\nUsing configuration:\n${humanConfig.slice(0, -1)}`));
        };
    }
};