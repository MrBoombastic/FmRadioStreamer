const express = require("express");
const app = express();
const prefix = "/fmradiostreamer/";
const {YTSearcher} = require('ytsearcher');
const config = require("../config.json");
const ytsearcher = new YTSearcher(config.apikey);
const ytdl = require('ytdl-core');
const fs = require('fs');
const spawn = require('child_process').spawn;
const {exec} = require("child_process");
const helpers = require('./helpers');
const led = require('./led');

module.exports = class webserver {
    constructor() {
        this.run = async function () {
            app.get(prefix + "yt/:song", (req, res) => {
                const safeSongName = req.params.song.replace(/[^a-zA-Z0-9\s]/g, "").replace(/ /g, "-");
                ytsearcher.search(req.params.song, {type: 'video'})
                    .then(result => {
                        const music = result.first;
                        ytdl.getInfo(music.id, (err) => {
                            if (err) throw err;
                            ffmpegorytdlWorking = true;
                            const video = ytdl(music.url, {format: 'aac'});
                            video.pipe(fs.createWriteStream(`ytdl-temp/${safeSongName}.aac`));
                            video.on('end', () => {
                                new led().workingLed();
                                const ffmpegProgress = spawn('ffmpeg', ['-i', `ytdl-temp/${safeSongName}.aac`, `music/${safeSongName}.wav`]);
                                //In case of debugging, uncomment this. Why FFmpeg produces data on stderr? idk
                                //ffmpegProgress.stdout.on('data', function (data) { console.log('stdout: ' + data.toString()); });
                                //ffmpegProgress.stderr.on('data', function (data) { console.log('stderr or ffmpeg: ' + data.toString()); });
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
                exec(`sudo pkill -2 pi_fm_rds`, () => {
                    exec(`mkfifo rds_ctl`);
                    const execWithStd = spawn(`sudo`, [
                            'src/pi_fm_rds',
                            `-ps "${config.PS}"`,
                            `-rt "${config.RT}"`,
                            '-freq', config.freq,
                            '-ctl', 'rds_ctl',
                            '-cutoff', `${config.quality}`,
                            '-audio', `"./music/${req.params.song}.wav"`],
                        {shell: true});
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
            app.listen(config.port, () => console.log(`FmRadioStreamer working!\nPort: ${config.port}\nFreq: ${Number(config.freq).toFixed(1)}\nPS: ${config.PS}\nRT: ${config.RT}`));
        };
    }
};