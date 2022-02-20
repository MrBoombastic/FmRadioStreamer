package main

import (
	"context"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/buttons"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/core"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/dashboard"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/ssd1306"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	tools.CheckRoot()
	// Exit handler
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	var wg sync.WaitGroup

	// Get current configuration
	cfg := config.Get()
	// Get local IP
	tools.RefreshLocalIP()
	log.Println("Your local IP is:", tools.LocalIP)
	log.Println("Starting peripherals")

	// Init GPIO pins and leds
	tools.InitGPIO()
	leds.Init()
	wg.Add(1)
	go leds.QuadGreensLoop(&wg, ctx)
	wg.Add(1)
	go leds.BlueLedLoop(&wg, ctx)

	if cfg.SSD1306 {
		// Init screen
		wg.Add(1)
		go ssd1306.Init(&wg, ctx)
	}

	// Init buttons
	buttons.Init()
	wg.Add(1)
	go buttons.Listen(&wg, ctx)
	log.Println("Peripherals started")

	if cfg.DynamicRT {
		log.Println("Starting dynamic RDS control")
		go core.RotateRT()
		log.Println("Dynamic RDS control started")
	}

	// Starting dashboard and core with no music
	log.Println("Starting dashboard")
	go dashboard.Init()
	log.Println("Starting core")
	core.Play("")
	log.Println("Core started")
	log.Println("Starting procedure done!")

	wg.Wait()

	// Code below is initated AFTER Ctrl-C or other terminating signal is sent

	// This should be executed, but it works fine whithout it and invoking it crashes my RPi :(
	// tools.StopGPIO()
	fmt.Println() // Usually "^C" is printed in the console, so it will be more pretty to go to next line
	log.Println("Gracefully exiting")
	log.Println("Killing core")
	core.Kill()
	log.Println("Gracefully exited")
}
