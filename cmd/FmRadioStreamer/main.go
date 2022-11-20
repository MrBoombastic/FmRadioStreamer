package main

import (
	"context"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/buttons"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/dashboard"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/rt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"github.com/pbar1/pkill-go"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Checking if running via sudo
	if !tools.CheckRoot() {
		logs.FmRadStrFatal("Not running as root! Exiting...")
	}
	// Checking libsndfile version
	libsndfileVersion, err := tools.CheckLibsndfileVersion()
	if err != nil {
		logs.FmRadStrFatal("Couldn't check libsndfile1-dev version. Possibly dependencies are not installed. Exiting...")
	}
	if libsndfileVersion >= 1.1 {
		logs.FmRadStrInfo("This system can play MP3, Opus and WAV files.")
	} else {
		logs.FmRadStrInfo("This system can play only Opus and WAV files. MP3 is not supported. Update libsndfile1-dev.")
	}
	// Exit handler
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	var wg sync.WaitGroup

	// Get current configuration
	cfg, err := config.Get()
	if err != nil {
		logs.FmRadStrFatal("Couldn't retrieve config. Exiting...")
	}
	// Get local IP
	tools.RefreshLocalIP()
	logs.FmRadStrInfo(fmt.Sprintf("Your local IP is: %v", tools.LocalIP))
	logs.FmRadStrInfo("Starting peripherals")

	// Init GPIO pins and leds
	err = tools.InitGPIO()
	if err != nil {
		log.Fatal(err)
	}
	leds.Init()
	wg.Add(1)
	go leds.QuadGreensLoop(&wg, ctx)
	wg.Add(1)
	go leds.BlueLedLoop(&wg, ctx)

	if cfg.SSD1306 {
		// Init screen
		wg.Add(1)
		go func() {
			err := ssd1306.Init(&wg, ctx)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	// Init buttons
	buttons.Init()
	wg.Add(1)
	go buttons.Listen(&wg, ctx)
	logs.FmRadStrInfo("Peripherals started")

	// Init Dynamic Radio Text
	if cfg.DynamicRT {
		rt.Primary, rt.Secondary = cfg.RT, cfg.RT
		go func() {
			err := rt.Rotate(config.GetDynamicRTInterval())
			if err != nil {
				logs.FmRadStrError(err)
			}
		}()
	}

	// Starting dashboard and core with no music
	go func() {
		err := dashboard.Init()
		if err != nil {
			log.Fatal(err)
		}
	}()

	logs.FmRadStrInfo("Starting core")
	go func() {
		err = core.Play(tools.Params{Type: tools.SilenceType})
		if err != nil {
			log.Fatal(err)
		}
	}()

	logs.FmRadStrInfo("Core started")
	logs.FmRadStrInfo("Ready!")

	wg.Wait()

	// Code below is initated AFTER Ctrl-C or other terminating signal is sent

	// This should be executed, but it works fine whithout it and invoking it crashes my RPi :(
	// tools.StopGPIO()
	fmt.Println() // Usually "^C" is printed in the console, so it will be more pretty to go to next line
	logs.FmRadStrInfo("Gracefully exiting")
	logs.FmRadStrInfo("Killing core")
	_, _ = pkill.Pkill("pi_fm_adv", os.Interrupt)
	logs.FmRadStrInfo("Gracefully exited. Bye!")
}
