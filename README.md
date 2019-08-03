# FmRadioStreamer
[ALPHA] RPI based FM streamer with OLED, buttons and LEDs support

## Installation:
Go to project directory, type `chmod +x install.sh`, then `sudo ./install.sh`. Installer will use NPM, but you can change to Yarn or PNPM. Then edit `config.json` file and enter [YT API key](https://developers.google.com/youtube/v3/getting-started).

## Running:
Type `sudo node start.js`. You can change RDS in `config.json` file. Visit `rpi_ip:1080/fmradiostreamer/list` to show available music and play it with `rpi_ip:1080/fmradiostreamer/play/music_name`.

## Functions:
- Show RDS on screen
- LEDs blinking when all is OK
- Yellow LED blinks if frequency is out of limit
- API - download music from YT and play
- Change frequency with buttons

## Dependencies note
This project uses Pi_Fm_Adv and my own fork of i2c-bus: i2c-bus-i2c-1, and other stuff listed in packages.json.

## Gallery
Don't blame me pls.
![Image](docs/IMG_20190728_172941.jpg?raw=true "Image")
![Image](docs/IMG_20190728_172930.jpg?raw=true "Image")

## What is NOT working a.k.a. bugs
- Optimisation sucks
- Files are in a mess

## What is in progress?
- Android app
- More stuff to change

## GPIO
- 1 - 3V3 Power
- 3 - GPIO 2 - button
- 5 - GPIO 3 - button
- 6 - GND for screen
- 7 - GPIO 4 -  antenna
- 29 - GPIO 5 - LED YELLOW
- 31 - GPIO 6 - LED
- 32 - GPIO 12 - button
- 33 - GPIO 13 - LED
- 35 - GPIO 19 - LED
- 36 - GPIO 16 - button
- 37 - GPIO 26 - LED
- 38 - GPIO 20 - button
- 39 - GND - for LEDs
- 40 - GPIO 21 - button
