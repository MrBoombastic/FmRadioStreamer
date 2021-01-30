const express = require("express"),
    app = express(),
    prefix = "/fmradiostreamer/",
    {YTSearcher} = require('ytsearcher'),
    config = require("../config.json"),
    ytsearcher = new YTSearcher(config.apikey),
    ytdl = require('ytdl-core'),
    fs = require('fs'),
    spawn = require('child_process').spawn,
    {exec} = require("child_process"),
    helpers = require('./helpers'),
    led = require('./led');
/*const runConfig = [
    'core/pi_fm_adv',
    `--ps "${config.PS}"`,
    `--rt "${config.RT}"`,
    '--freq', freq,
    '--ctl', 'rds_ctl',
    '--power', config.power
];
*/
module.exports = class webserver {
    constructor() {
        this.run = async function () {
            app.get(prefix + "yt/:song", async (req, res) => {
                const safeSongName = req.params.song.replace(/[^a-zA-Z0-9\s]/g, "").replace(/ /g, "-");
                ffmpegorytdlWorking = true;
                const result = await ytsearcher.search(req.params.song, {type: 'video'});
                const music = await result.first;
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
                            console.log("FFmpeg returned error:" + code);
                            return res.send("FFmpeg error!");
                        }
                        exec('sudo rm -r ./ytdl-temp/*');
                        res.send("Done!");
                    });
                });
            });

            app.get(prefix + "list", (req, res) => {
                let finallist = "";
                const musiclist = fs.readdirSync('./music/');
                musiclist.forEach(function (song) {
                    finallist = finallist + song + ", ";
                });
                res.json(JSON.parse(`{"list": "${finallist.slice(0, -2)}"}`));
            });
            app.get(prefix + "play/:song", (req, res) => {
                exec(`sudo pkill -2 pi_fm_adv`, () => {
                    exec(`mkfifo rds_ctl`);
                    const execWithStd = spawn(`sudo`, [
                        'core/pi_fm_adv',
                        `--ps "${config.PS}"`,
                        `--rt "${config.RT}"`,
                        '--freq', freq,
                        '--ctl', 'rds_ctl',
                        '--power', config.power,
                        '--audio', `"./music/${req.params.song}.wav"`], {shell: true});
                    //In case of debugging, you can uncomment this safely:
                    //execWithStd.stdout.on('data', function (data) { console.log('stdout: ' + data.toString()); });
                    execWithStd.stderr.on('data', function (data) {
                        console.log('stderr: ' + data.toString());
                    });
                    res.send("ok");
                });
            });
            app.get(prefix + "change/:setting/:value", (req, res) => {
                const setting = req.params.setting.toString();
                const value = req.params.value.toString();
                if (setting !== "PS" && setting !== "RT") return res.status(405).send("Allowed settings: PT, RT, TA, PTY.");
                else {
                    helpers.save(setting, value);
                    res.send(setting.toString() + ", " + value.toString());
                }
            });
            /*
            app.get(prefix + "set/:setting/:value", (req, res) => {
                if (setting != "PS" && setting != "RT" && setting != "TA" && setting != "PTY") return res.status(405).send("Allowed settings: PT, RT, TA, PTY.")
                else {save(setting, value);res.send(setting.toString() + ", " + value.toString())}
            })*/
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
            app.listen(config.port, () => console.log(`FmRadioStreamer working!\nPort: ${config.port}\nFreq: ${freq}\nPS: ${config.PS}\nRT: ${config.RT}`));
        };
    }
};