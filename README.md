# FmRadioStreamer

Raspberry Pi based FM streamer with OLED, buttons and LEDs support.

[![Go Reference](https://pkg.go.dev/badge/github.com/MrBoombastic/FmRadioStreamer.svg)](https://pkg.go.dev/github.com/MrBoombastic/FmRadioStreamer)
[![CodeFactor](https://www.codefactor.io/repository/github/mrboombastic/fmradiostreamer/badge)](https://www.codefactor.io/repository/github/mrboombastic/fmradiostreamer)
[![BCH compliance](https://bettercodehub.com/edge/badge/MrBoombastic/FmRadioStreamer?branch=master)](https://bettercodehub.com/)

## Building

Clone this repository, run `go get` and then `go build`. Of course, you have to install `Go` first.

If you want to build for other device (e.g. build on Windows and run on RaspberyPi), use cross-compilation!

Environmental variables for RPi Zero: `GOOS=linux;GOARCH=arm;GOARM=5`. For other RPis, `GOARM` may be unnecessary.

## Installation

Build or copy binary from `releases` tab. Copy `install.sh` to your directory, type `chmod +x install.sh`,
then `sudo ./install.sh`.

**WARNING!** This script will install some unnecessary dependencies like Python for `youtube-dl`. If you already have
working `youtube-dl` instance, edit the script (comment line #6).

Copy `config.json.example`, rename to `config.json` aned edit. Here you can get
your [YouTube API Key](https://developers.google.com/youtube/v3/getting-started). Also, you have to add `gpu_freq=250`
in `/boot/config.txt`.

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

## JSON settings

- dynamicRTInterval: dynamic RT switching interval in seconds, default 20

| Option            | Description                                                                  | Type (additional info) | Default                                  | Notes                                                 |
|-------------------|------------------------------------------------------------------------------|------------------------|------------------------------------------|-------------------------------------------------------|
| freq              | Frequency in MHz                                                             | number (float64)       | 108.0                                    |                                                       |
| multiplier        | Frequency multiplier (used when using physical buttons)                      | number (float64)       | 0.1                                      |                                                       |
| PS                | RDS station name                                                             | string (len: 8)        | FmRadStr                                 |                                                       |
| RT                | RDS station text                                                             | string (len: 64)       | RPi based radio streamer. It is working! |                                                       |
| PI                | RDS station ID                                                               | string (len: 4, hex)   | FFFF                                     |                                                       |
| TP                | RDS traffic programme                                                        | string                 | (empty string)                           | I don't know the type of this setting.                |
| PTY               | RDS programme type                                                           | number (uint)          | 0                                        | Warning: different meaning in EU and US! Range: 0-31. |
| ytApiKey          | YouTube API V3 Key                                                           | string                 | (empty string)                           |                                                       |
| port              | Web dashboard port                                                           | number (uint16)        | 80                                       |                                                       |
| power             | RaspberryPi output power                                                     | number (uint8)         | 5                                        | Refer to PiFmAdv for more info.                       |
| mpx               | Mpx output power ("volume")                                                  | number (uint)          | 30                                       | See above.                                            |
| preemph           | pre-emphasis                                                                 | string                 | eu                                       | Possible values: "eu", "us".                          |
| antennaGPIO       | GPIO antenna header                                                          | number (uint8)         | 4                                        | Possible values: 4, 20, 32, 34                        |
| ssd1306           | OLED screen type SSD1306 enabled                                             | bool                   | true                                     |                                                       |
| dynamicRT         | Switching RT between that saved in config and current playing audio filename | bool                   | true                                     |                                                       |
| dynamicRTInterval | Dynamic RT switching interval in seconds                                     | number (uint)          | 20                                       |                                                       | 





