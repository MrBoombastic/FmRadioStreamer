# FmRadioStreamer

Raspberry Pi based FM streamer with OLED, buttons and LEDs support.

[![Go Reference](https://pkg.go.dev/badge/github.com/MrBoombastic/FmRadioStreamer.svg)](https://pkg.go.dev/github.com/MrBoombastic/FmRadioStreamer)
[![CodeFactor](https://www.codefactor.io/repository/github/mrboombastic/fmradiostreamer/badge)](https://www.codefactor.io/repository/github/mrboombastic/fmradiostreamer)
[![Go Report Card](https://goreportcard.com/badge/github.com/MrBoombastic/FmRadioStreamer)](https://goreportcard.com/report/github.com/MrBoombastic/FmRadioStreamer)

## Building (optional)

Clone this repository, run `go get` and then `go build`. Of course, you have to install `Go` first.

If you want to build for other device (e.g. build on Windows and run on RaspberyPi), use cross-compilation!

Environmental variables for RPi Zero: `GOOS=linux;GOARCH=arm;GOARM=5`. For other RPis, `GOARM` field may be unnecessary.

## Installation

Build or copy binary from `releases` tab. Copy `install.sh` to your directory, type `chmod +x install.sh`,
then `sudo ./install.sh`.

> [!CAUTION]
> This script may install some unwanted dependencies. Review the script if you want.

Copy `config.json.example`, rename to `config.json` aned edit.

Here you can get your [YouTube API Key](https://developers.google.com/youtube/v3/getting-started).

> [!IMPORTANT]  
> You have to add `gpu_freq=250` in `/boot/firmware/config.txt` and also enable I2C in `raspi-config`.

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
![Image](docs/dashboard_new.png?raw=true "Dashboard screenshot")
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

## Which format should I choose?

Downloading and converting benchmark:

- Song URL: https://www.youtube.com/watch?v=SaoBbCC66I4
- Results:

| Format | Converting time | File size |
|--------|-----------------|-----------|
| MP3    | 6:17            | 12 431 KB |
| WAV    | 4:20 hehe       | 6 499 KB  |
| Opus   | 3:32            | 69 891 KB |

For pretty good audio quality, low file size and best conversion time, use Opus. If for
some reason you can't use Opus format and have outdated `libsndfile1-dev` package, use WAV.
Use MP3 if you don't want to re-encode your existing music library.

## MP3 support (pre-bookworm)

1. Update your RPi (optional, but recommended).

```bash
sudo bash -c 'for i in update {,dist-}upgrade auto{remove,clean}; do apt-get $i -y; done'
```

2. Edit `/etc/apt/sources.list` file using your favourite editor, for example `sudo vim /etc/apt/sources.list`. In my
   case, the file looks like this:

```bash
deb http://raspbian.raspberrypi.org/raspbian/ bullseye main contrib non-free rpi
# Uncomment line below then 'apt-get update' to enable 'apt-get source'
#deb-src http://raspbian.raspberrypi.org/raspbian/ bullseye main contrib non-free rpi
```

Replace `bullseye` with `bookworm`.

```bash
deb http://raspbian.raspberrypi.org/raspbian/ bookworm main contrib non-free rpi
```

3. Run `sudo apt update`. It wil say that 69420 packages can be updated - ignore that.
4. Run `sudo apt install libsndfile1-dev`. It will update only this package and its dependencies.
5. Confirm that `libsndfile1-dev` is now version 1.1.0 or later. Run `sudo apt-cache policy libsndfile1-dev`. My output:

```bash
libsndfile1-dev:
  Installed: 1.1.0-2 
(...)
```

5. Repeat step 1, but now revert `bookworm` to `bullseye`. Done. 🎉

## JSON settings

| Option            | Description                                                                  | Type (additional info) | Default                                  | Notes                                                                                                                      |
|-------------------|------------------------------------------------------------------------------|------------------------|------------------------------------------|----------------------------------------------------------------------------------------------------------------------------|
| freq              | Frequency in MHz                                                             | number (float64)       | 108.0                                    |                                                                                                                            |
| format            | Audio format for saved files                                                 | string                 | opus                                     | formats limited by `libsndfile1-dev` and by its version, more [here](https://libsndfile.github.io/libsndfile/formats.html) |                       | 
| multiplier        | Frequency multiplier (used when using physical buttons)                      | number (float64)       | 0.1                                      |                                                                                                                            |
| PS                | RDS station name                                                             | string (len: 8)        | FmRadStr                                 |                                                                                                                            |
| RT                | RDS station text                                                             | string (len: 64)       | RPi based radio streamer. It is working! |                                                                                                                            |
| PI                | RDS station ID                                                               | string (len: 4, hex)   | FFFF                                     |                                                                                                                            |
| TP                | RDS traffic programme                                                        | string                 | (empty string)                           | I don't know the type of this setting.                                                                                     |
| PTY               | RDS programme type                                                           | number (uint)          | 0                                        | Warning: different meaning in EU and US! Range: 0-31.                                                                      |
| ytApiKey          | YouTube API V3 Key                                                           | string                 | (empty string)                           |                                                                                                                            |
| port              | Web dashboard port                                                           | number (uint16)        | 80                                       |                                                                                                                            |
| power             | RaspberryPi output power                                                     | number (uint8)         | 5                                        | Refer to PiFmAdv for more info.                                                                                            |
| mpx               | Mpx output power ("volume")                                                  | number (uint)          | 30                                       | See above.                                                                                                                 |
| preemph           | pre-emphasis                                                                 | string                 | eu                                       | Possible values: "eu", "us".                                                                                               |
| antennaGPIO       | GPIO antenna header                                                          | number (uint8)         | 4                                        | Possible values: 4, 20, 32, 34                                                                                             |
| ssd1306           | OLED screen type SSD1306 enabled                                             | bool                   | true                                     |                                                                                                                            |
| dynamicRT         | Switching RT between that saved in config and current playing audio filename | bool                   | true                                     |                                                                                                                            |
| dynamicRTInterval | Dynamic RT switching interval in seconds                                     | number (uint)          | 20                                       |                                                                                                                            |
| verbose           | Additional logs in console                                                   | bool                   | false                                    |                                                                                                                            |
| ytdlp             | Check if yt-dlp is installed and install, if not                             | bool                   | true                                     | If you have yt-dlp installed and always up-to-date, set to false.                                                          |






