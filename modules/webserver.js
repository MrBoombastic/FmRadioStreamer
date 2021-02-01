const express = require("express"),
    app = express(),
    {YTSearcher} = require('ytsearcher'),
    config = require("../config.json"),
    ytsearcher = new YTSearcher(config.apikey),
    ytdl = require('ytdl-core'),
    fs = require('fs'),
    helpers = require("./helpers"),
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
            app.get("/mng", (req, res) => {
                switch (req.query.action) {
                    case 'loudstop':
                        helpers.playPiFmADV(config);
                        break;
                    case 'superstop':
                        helpers.killPiFmADV();
                        break;
                    case 'Papayas':
                        console.log('Mangoes and papayas are $2.79 a pound.');
                        // expected output: "Mangoes and papayas are $2.79 a pound."
                        break;
                    default:
                        console.log(`Sorry, we are out of ${expr}.`);
                }

                res.render('dash.ejs', {list: fs.readdirSync('./music/')});
            });

            app.get("/yt/:song", async (req, res) => {
                ffmpegorytdlWorking = true;
                const result = await ytsearcher.search(req.params.song, {type: 'video'});
                const music = await result.first;
                const safeSongName = music.title.replace(/[^a-zA-Z0-9\s]/g, "").replace(/ /g, "-");
                const video = ytdl(music.url, {quality: "highestaudio", filter: "audioonly"});
                video.pipe(fs.createWriteStream(`ytdl-temp/${safeSongName}`));
                video.on('end', async () => {
                    const conversion = await helpers.convertToWav(safeSongName);
                    if (conversion) {
                        res.json({
                            status: "done",
                            desc: safeSongName
                        });
                    } else return res.status(500).json({
                        status: "error",
                        desc: "FFmpeg returned error: " + e
                    });
                });
            });

            app.get("/list", (req, res) => {
                let finallist = "";
                const musiclist = fs.readdirSync('./music/');
                musiclist.forEach(song => finallist += (song + ", "));
                res.json({status: "done", desc: finallist.slice(0, -2)});
            });
            app.get("/play", async (req, res) => {
                if (!req.query.song) return res.status(404).json({status: "failed", desc: "not found"});
                await helpers.playPiFmADV(config, req.query.song);
                res.json({status: "done", desc: "If file exist, it will be played."});
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
            await helpers.playPiFmADV(config);

            app.listen(config.port, () => console.log("Webserver is up!"));
        };
    }
};