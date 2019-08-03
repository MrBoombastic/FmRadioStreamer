const config = require("./config.json")
const { exec } = require("child_process");
exec(`sudo ./radio.sh '${config.PS}' '${config.RT}' ${config.freq} & sudo node radio.js`, (out, stdout, stderr) => { console.log(out, stdout,stderr) })
process.on('SIGINT', () => {
    exec(`sudo pkill -2 pi_fm_adv`)
    exec(`sudo pkill -2 fmradiostreamer`)
    process.exit();
})