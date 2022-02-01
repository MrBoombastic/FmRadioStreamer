# Warning: rewrite in progress!

## Here is the latest "stable" version: [click](https://github.com/MrBoombastic/FmRadioStreamer/tree/bc3cae0455e9352db580cb24f56f9788ea515354)

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
This project uses PiFmAdv, FFmpeg, libsndfile1-dev, Start Bootstrap - Bare and other stuff listed in package.json.

## Hardware
SSD1306 screen, 4 THT buttons, ~400 Ohm resistors, ~20k Ohm resistors, leds, female-male and male-male cables. Tested on Raspberry Pi Zero W with goldpins soldered by me. :D

## What if I don't have that hardware?
The minimum requirement is RPi. FmRadioStreamer SHOULD work without them.

## Gallery
In the `docs` directory there are pictures of first version of this project. Watch on your own risk!
![Image](docs/hwv2_1.jpg?raw=true "Image")
![Image](docs/hwv2_2.jpg?raw=true "Image")
![Image](docs/webserver.png?raw=true "Image")

## What is NOT working a.k.a. bugs
- ~~Optimisation sucks~~ maybe not
- ~~Files are in a mess~~ definitly not

## What is in progress?
- ~~Android app~~ not much time
- More stuff to change/improve

## GPIO
- 1 - 3V3 Power
- 3 - GPIO 2 - screen
- 5 - GPIO 3 - screen
- 6 - GND - for screen
- 7 - GPIO 4 -  antenna
- 26 - GPIO 7 - BLUE WORKING LED
- 29 - GPIO 5 - YELLOW WARNING LED
- 31 - GPIO 6 - LED
- 32 - GPIO 12 - button
- 33 - GPIO 13 - LED
- 35 - GPIO 19 - LED
- 36 - GPIO 16 - button
- 37 - GPIO 26 - LED
- 38 - GPIO 20 - button
- 39 - GND - for LEDs
- 40 - GPIO 21 - button