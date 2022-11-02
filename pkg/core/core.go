package core

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"golang.org/x/sys/unix"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// GenerateOptions generate options to pass to PiFmAdv
func GenerateOptions(audio string) []string {
	cfg := config.Get()
	options := []string{
		"--ps", cfg.PS,
		"--rt", cfg.RT,
		"--freq", fmt.Sprintf("%f", cfg.Frequency),
		"--power", fmt.Sprintf("%v", cfg.Power),
		"--pi", cfg.PI,
		"--tp", cfg.TP,
		"--pty", fmt.Sprintf("%v", cfg.PTY),
		"--preemph", cfg.Preemph,
		"--gpio", fmt.Sprintf("%v", cfg.AntennaGPIO),
		"--mpx", fmt.Sprintf("%v", cfg.Mpx),
		"--ctl", "rds_ctl",
	}
	if audio != "" && audio != "-" {
		options = append(options, "--audio", fmt.Sprintf("./music/%v", audio))
	} else if audio == "-" {
		options = append(options, "--audio", "-")
	}
	return options
}

// run starts playing music or silence via PiFmAdv
func run(name string, args []string) error {
	PiFmAdv := exec.Command(name, args...)
	stderr, err := PiFmAdv.StderrPipe()
	if err != nil {
		return err
	}
	// Stdout commented out for clearer command output, safe to undo!
	/*
		stdout, err := PiFmAdv.StdoutPipe()
		if err != nil {
			log.Println(err)
		}
	*/
	err = PiFmAdv.Start()
	if err != nil {
		return err
	}
	// Support for "dynamic" RDS - getting current playing file name
	if args[len(args)-2] == "--audio" {
		audio := args[len(args)-1]
		extension := filepath.Ext(audio)
		audio = strings.TrimSuffix(audio, extension)
		alternateRT = strings.Replace(audio, "./music/", "", 1)
		if len(alternateRT) > 64 {
			alternateRT = alternateRT[0:63]
		}
	}
	cmderr, _ := io.ReadAll(stderr)
	if fmt.Sprintf("%s", cmderr) != "" {
		log.Printf("PiFmAdv: %s", cmderr)
	}
	// Stdout commented out for clearer command output, safe to undo!
	/*
		cmdout, _ := io.ReadAll(stdout)
		fmt.Printf("PiFmADV: %s", cmdout)
	*/
	return nil
}

// Kill stops PiFmAdv using pkill and SIGINT
func Kill() error {
	cmd := exec.Command("pkill", "-2", "pi_fm_adv")
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("pkill: %v", err)
	}
	err = cmd.Wait()
	// Preventing RPi overloading
	if err != nil {
		fmt.Println(err)
		if err.Error() == "exit status 1" {
			return nil
		}
		return fmt.Errorf("pkill: %v", err)
	}
	return nil
}

// Play generates options with GenerateOptions function and launches PiFmAdv
func Play(audio string) error {
	// Make sure that previous playback is stopped
	err := Kill()
	if err != nil {
		return err
	}
	options := GenerateOptions(audio)
	err = run("core/pi_fm_adv", options)
	if err != nil {
		return err
	}
	return nil
}

// Sox generates options with GenerateOptions function, launches SoX, then pipes output to PiFmAdv
func Sox(path string) error {
	// Make sure that previous playback is stopped
	err := Kill()
	if err != nil {
		return err
	}
	options := GenerateOptions("-")
	textoptions := fmt.Sprintf("sox %v -t wav - | sudo core/pi_fm_adv %v", path, strings.Join(options, " "))

	// Go is doing some weird things when using "|" in exec.Command, so we will run command through temp script
	file, err := os.Create("temp.sh")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(file)

	_, err = file.WriteString(textoptions)
	if err != nil {
		return err
	}
	cmd, err := exec.Command("/bin/sh", "temp.sh").Output()
	if err != nil {
		return err
	}
	// AlLoCaTiOnS aRe BaD!
	_, err = os.Stdout.Write(cmd)
	if err != nil {
		return err
	}
	return nil
}

var alternateRT = config.GetRT()
var currentRTState = 0

// RotateRT enables switching RT between that saved in config and current playing audio filename
func RotateRT() {
	err := os.Remove("rds_ctl")
	if err != nil {
		log.Println("ERROR: Cannot remove rds_ctl pipe file. Missing?")
	}
	err = unix.Mkfifo("rds_ctl", 0666)
	if err != nil {
		log.Println("ERROR: Cannot create pipe file: ", err)
		return
	}
	f, err := os.OpenFile("rds_ctl", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println("ERROR: Cannot open pipe file: ", err)
		return
	}
	for {
		if currentRTState == 0 {
			_, err := f.WriteString("AB ON\n")
			if err != nil {
				log.Println("ERROR: Cannot update dynamic RT text (A/B flag)")
				break
			}
			_, err = f.WriteString(fmt.Sprintf("RT %s", config.GetRT()))
			if err != nil {
				log.Println("ERROR: Cannot update dynamic RT text")
				break
			}
			currentRTState++
		} else {
			_, err := f.WriteString("AB OFF\n")
			if err != nil {
				log.Println("ERROR: Cannot update dynamic RT text (A/B flag)")
				break
			}
			_, err = f.WriteString(fmt.Sprintf("RT %s", alternateRT))
			if err != nil {
				log.Println("ERROR: Cannot update dynamic RT text")
				break
			}
			currentRTState--
		}
		time.Sleep(time.Second * time.Duration(config.GetDynamicRTInterval()))
	}
}
