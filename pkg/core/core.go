package core

import (
	"errors"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
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
	formattedError := strings.TrimSpace(fmt.Sprintf("%s", cmderr))
	if formattedError != "" {
		if strings.Contains(formattedError, "WARNING:") {
			formattedError = strings.Replace(formattedError, "WARNING: ", "", 1)
			logs.PiFmAdvWarn(formattedError)
		} else {
			logs.PiFmAdvError(formattedError)
		}
	}
	// Stdout commented out for clearer command output, safe to undo!
	/*
		cmdout, _ := io.ReadAll(stdout)
		logs.PiFmAdv.Info().Send("%v", cmdout)
	*/
	return nil
}

// Kill stops PiFmAdv using pkill and SIGINT
func Kill() error {
	cmd := exec.Command("sudo", "pkill", "-2", "pi_fm_adv")
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("pkill: %v", err)
	}
	err = cmd.Wait()
	// Preventing RPi overloading
	if err != nil {
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
			log.Fatalln(err)
			return
		}
	}(file)

	_, err = file.WriteString(textoptions)
	if err != nil {
		return err
	}
	cmd, err := exec.Command("/bin/sh", "temp.sh").Output()
	if err != nil {
		if err.Error() != "exit status 2" {
			return err
		}
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
func RotateRT() error {
	err := os.Remove("rds_ctl")
	if err != nil {
		return errors.New("cannot remove rds_ctl pipe file, maybe it's missing")
	}
	err = unix.Mkfifo("rds_ctl", 0666)
	if err != nil {
		return fmt.Errorf("cannot create pipe file: %v", err)
	}
	f, err := os.OpenFile("rds_ctl", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return fmt.Errorf("cannot open pipe file: %v", err)
	}
	for {
		if currentRTState == 0 {
			_, err := f.WriteString("AB ON\n")
			if err != nil {
				return errors.New("cannot update dynamic RT text (A/B flag)")
			}
			_, err = f.WriteString(fmt.Sprintf("RT %s", config.GetRT()))
			if err != nil {
				return errors.New("cannot update dynamic RT text")
			}
			currentRTState++
		} else {
			_, err := f.WriteString("AB OFF\n")
			if err != nil {
				return errors.New("ERROR: Cannot update dynamic RT text (A/B flag)")
			}
			_, err = f.WriteString(fmt.Sprintf("RT %s", alternateRT))
			if err != nil {
				return errors.New("ERROR: Cannot update dynamic RT text")
			}
			currentRTState--
		}
		time.Sleep(time.Second * time.Duration(config.GetDynamicRTInterval()))
	}
}
