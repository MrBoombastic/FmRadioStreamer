# FmRadioStreamer

Raspberry Pi based FM streamer with OLED, buttons and LEDs support.

[![Go Reference](https://pkg.go.dev/badge/github.com/MrBoombastic/FmRadioStreamer.svg)](https://pkg.go.dev/github.com/MrBoombastic/FmRadioStreamer)
[![CodeFactor](https://www.codefactor.io/repository/github/mrboombastic/fmradiostreamer/badge)](https://www.codefactor.io/repository/github/mrboombastic/fmradiostreamer)
[![BCH compliance](https://bettercodehub.com/edge/badge/MrBoombastic/FmRadioStreamer?branch=master)](https://bettercodehub.com/)

## Installation

Go to project directory, type `chmod +x install.sh`, then `sudo ./install.sh`.

**WARNING!** This script will install some unnecessary dependencies like Python for `youtube-dl`. If you already have
working `youtube-dl` instance, edit the script (comment line #6).

Edit `config.json` file and enter your [YouTube API Key](https://developers.google.com/youtube/v3/getting-started).
Also, you have to add `gpu_freq=250` in `/boot/config.txt`.

## Running

Rename `config.json.example` to `config.json`. Then type `sudo node index.js`. You can change RDS in `config.json` file.
Go to API Docs section for more.

## Functions

- RDS and other options rendering on screen
- LEDs blinking when everything is OK
- Yellow LED blinking if frequency is out of limits
- Blue LED blinking if music is being processed (downloading, converting)
- Downloading music from YouTube
- Changing frequency with buttons
- Web dashboard

## API docs

None. Please reverse-engineer pkg/dashboard for that. :(

## Dependencies note

This project uses PiFmAdv, FFmpeg, libsndfile1-dev, youtube-dl and other stuff listed in go.mod.

## Optional hardware

SSD1306 screen, 4 THT buttons, ~400 Ohm resistors, ~20k Ohm resistors, LEDs, female-male and male-male wires. Tested on
Raspberry Pi Zero W Rev 1.1.

## What if I don't have that hardware?

The minimum requirement is the RaspberryPi. FmRadioStreamer SHOULD work without any other peripherals.

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

## Settings

- freq: frequency, default 108.0 MHz
- multiplier: frequency multiplier (when using physical buttons), default 0.1
- PS: RDS station name, default "FmRadStr"
- RT: RDS station description, default "RPi based radio streamer. It is working!"
- PI: RDS station ID in hex, default "FFFF"
- TP: RDS traffic programme, default "" (empty string)
- PTY: RDS programme type, default 0
- apikey: YouTube API V3 Key, default "" (empty string)
- port: dashboard port, default 80
- power: RaspberryPi output power, default 5
- mpx: mpx output power ("volume"), default 30
- preemph: pre-emphasis, "eu" or "us", default "eu"
- antennaGPIO: GPIO antenna header, default 4, available: 4, 20, 32, 34
- ssd1306: OLED screen model SSD1306, default true