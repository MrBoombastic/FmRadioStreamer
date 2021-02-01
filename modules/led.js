const ledgpio = require('onoff').Gpio;
const led1 = new ledgpio(19, 'out');
const led2 = new ledgpio(6, 'out');
const led3 = new ledgpio(13, 'out');
const led4 = new ledgpio(26, 'out');
const ledWarn = new ledgpio(5, 'out');
const ledWorking = new ledgpio(7, 'out');
const sleeptime = 500;

function sleep(time) {
    return new Promise((resolve) => {
        setTimeout(resolve, time);
    });
}

module.exports = class led {
    constructor() {
        this.run = function () {
            this.stop();
            new Promise(async () => {
                while (true) {
                    // led1
                    led4.writeSync(0);
                    led1.writeSync(1);
                    await sleep(sleeptime);
                    // led 2
                    led1.writeSync(0);
                    led2.writeSync(1);
                    await sleep(sleeptime);
                    // led3
                    led2.writeSync(0);
                    led3.writeSync(1);
                    await sleep(sleeptime);
                    // led4
                    led3.writeSync(0);
                    led4.writeSync(1);
                    await sleep(sleeptime);
                }
            });
        };
        this.stop = function () {
            led1.writeSync(0);
            led2.writeSync(0);
            led3.writeSync(0);
            led4.writeSync(0);
            ledWarn.writeSync(0);
            ledWorking.writeSync(0);
        };
        this.unexport = function () {
            this.stop();
            led1.unexport();
            led2.unexport();
            led3.unexport();
            led4.unexport();
            ledWarn.unexport();
            ledWorking.unexport();
        };
        this.ledWarnBlink = async function () {
            ledWarn.writeSync(0);
            ledWarn.writeSync(1);
            await sleep(2000);
            ledWarn.writeSync(0);
        };
        this.workingLed = async function () {
            while (true) {
                if (ffmpegorytdlWorking) {
                    ledWorking.writeSync(1);
                    await sleep(500);
                    ledWorking.writeSync(0);
                    await sleep(500);
                } else await sleep(1000);
            }
        };
    }
};