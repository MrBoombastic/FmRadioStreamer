# FmRadioStreamer

[BETA] Raspberry Pi based FM streamer with OLED, buttons and LEDs support.

## Installation:

Go to project directory, type `chmod +x install.sh`, then `sudo ./install.sh`. **WARNING!!!** Installer will use Yarn to
install packages INSTEAD of NPM!

Edit `config.json` file and enter [YT API key](https://developers.google.com/youtube/v3/getting-started). Also, you have
to add `gpu_freq=250` in `/boot/config.txt`.

## Running:
Rename `config.json.example` to `config.json`. Then type `sudo node index.js`. You can change RDS in `config.json` file. Go to API Docs section for more.

## Functions:
- Show RDS on screen
- LEDs blinking when all is OK
- Yellow LED blinks if frequency is out of limit
- Blue LED blinks if music is being dowloaded from YouTube or converted
- API - download music from YT and play
- Change and save frequency with buttons
- Web dashboard

## API Docs
Removed, but it's easy to reverse-engineer. I hope so.

## Dependencies note
This project uses PiFmAdv, FFmpeg, libsndfile1-dev, youtube-dl and other stuff listed in go.mod.

## Hardware

SSD1306 screen, 4 THT buttons, ~400 Ohm resistors, ~20k Ohm resistors, LEDs, female-male and male-male wires. Tested on
Raspberry Pi Zero W Rev 1.1.

## What if I don't have that hardware?

The minimum requirement is RPi. FmRadioStreamer SHOULD work without any other peripherals.

## Gallery

In the `docs` directory there are pictures of first version of this project. Watch on your own risk!
![Image](docs/hwv2rev2_1.jpg?raw=true "Image")
![Image](docs/hwv2rev2_2.jpg?raw=true "Image")
![Image](docs/dashboard.png?raw=true "Dashboard screenshot")
![Image](docs/sdrsharp.png?raw=true "SDR# screenshot")

## GPIO

- 1 - 3V3 Power
- 3 - GPIO 2 - screen
- 5 - GPIO 3 - screen
- 6 - GND - for screen
- 7 - GPIO 4 - antenna
- 26 - GPIO 7 - blue LED (audio conversion in progress)
- 29 - GPIO 5 - yellow LED (frequency out of limits)
- 31 - GPIO 6 - green LED
- 32 - GPIO 12 - button (screen color inversion)
- 33 - GPIO 13 - green LED
- 35 - GPIO 19 - green LED
- 36 - GPIO 16 - button (frequency multiplier switch)
- 37 - GPIO 26 - green LED
- 38 - GPIO 20 - button (frequency up)
- 39 - GND - for LEDs
- 40 - GPIO 21 - button (frequency down)