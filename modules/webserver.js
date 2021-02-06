const express = require("express"),
    app = express(),
    config = require("../config.json"),
    fs = require('fs'),
    helpers = require("./helpers");

app.set('view engine', 'ejs');
app.use('/public', express.static('./public'));

module.exports = class webserver {
    constructor() {
        this.run = async function () {
            app.get("/", (req, res) => {
                res.render('index.ejs', {list: fs.readdirSync('./music/'), config});
            });
            app.get("/mng", async (req, res) => {
                switch (req.query.action) {
                    case 'loudstop':
                        await helpers.playPiFmADV();
                        res.sendStatus(200);
                        break;
                    case 'superstop':
                        await helpers.killPiFmADV();
                        res.sendStatus(200);
                        break;
                    case 'yt':
                        if (req.query.searchOnly) {
                            await helpers.getYT(req.query.song, true)
                                .then(data => res.json(data))
                                .catch(e => res.status(500).json({
                                    status: "failed",
                                    desc: e
                                }));
                        } else {
                            res.json({status: "done", desc: "Request understood. Processing now..."});
                            await helpers.getYT(req.query.song);
                        }
                        break;
                    case 'list':
                        res.json({status: "done", desc: fs.readdirSync('./music/')});
                        break;
                    case 'play':
                        if (!req.query.song) return res.status(404).json({status: "failed", desc: "not found"});
                        await helpers.playPiFmADV(req.query.song);
                        res.json({status: "done", desc: "If file exists, it will be played."});
                        break;
                    default:
                        return res.status(404).json({status: "failed", desc: "Action not found!"});
                }
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
            await helpers.playPiFmADV();

            app.listen(config.port, () => console.log("Webserver is up!"));
        };
    }
};