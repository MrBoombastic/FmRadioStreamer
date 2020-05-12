# FmRadioStreamer
[BETA] RPi based FM streamer with OLED, buttons and LEDs support.
Using PiFmAdv for better quality sound.

## Installation:
Go to project directory, type `chmod +x install.sh`, then `sudo ./install.sh`. Installer will use NPM, but you can change it to Yarn or PNPM. Then edit `config.json` file and enter [YT API key](https://developers.google.com/youtube/v3/getting-started).
Also, you have to add `gpu_freq=250` in `/boot/config.txt`.

## Running:
Rename `config.json.example` to `config.json`. Then type `sudo node index.js`. You can change RDS in `config.json` file. Go to API Docs section for more.

## Functions:
- Show RDS on screen
- LEDs blinking when all is OK
- Yellow LED blinks if frequency is out of limit
- Blue LED blinks if music is being dowloaded from YouTube or converted
- API - download music from YT and play
- Change and save frequency with buttons

## API Docs
- /fmradiostreamer/yt/(song) - Downloads song from YT and save with a filename of requested text.
- /fmradiostreamer/list - Returns song list.
- ~~/fmradiostreamer/change/(setting)/(value) - Changes settings. Saves to config, so changes are effective when you restart application or change music.~~ disabled, need rewrite.
- /fmradiostreamer/play/(song) - Starts song.

## Dependencies note
This project uses PiFmAdv, FFmpeg, libsndfile1-dev and other stuff listed in packages.json.

## Hardware
SSD1306 screen, 4 THT buttons, ~400 Ohm resistors, ~20k Ohm resistors, leds, female-male and male-male cables. Tested on Raspberry Pi Zero W with goldpins soldered by me. :D

## What if I don't have that hardware?
The minimum requirement is RPi. FmRadioStreamer SHOULD work without them.

## Gallery
In the `docs` directory there are pictures of first version of this project. Watch on your own risk!
![Image](docs/hwv2_1.jpg?raw=true "Image")
![Image](docs/hwv2_2.jpg?raw=true "Image")

## What is NOT working a.k.a. bugs
- ~~Optimisation sucks~~ maybe not
- ~~Files are in a mess~~ definitly not

## What is in progress?
- ~~Android app~~ not too much time
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