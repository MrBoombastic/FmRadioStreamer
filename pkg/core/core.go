package core

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/rt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"github.com/pbar1/pkill-go"
	"os"
	"path/filepath"
	"strings"
)

// GenerateOptions generates options to pass to PiFmAdv
func GenerateOptions(params tools.Params) []string {
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
	// When playing file, we need to add which one. Stream needs "-" every time, and silence needs nothing.
	if params.Type == tools.FileType {
		options = append(options, "--audio", fmt.Sprintf("./music/%v", params.Audio))
	} else if params.Type == tools.StreamType {
		options = append(options, "--audio", "-")
	}
	return options
}

// run starts playing music, silence or stream via PiFmAdv
func run(name string, params tools.Params) error {
	// Support for "dynamic" RDS - getting current playing file name - only in FileType mode!
	if params.Type == tools.FileType {
		extension := filepath.Ext(params.Audio)
		audio := strings.TrimSuffix(params.Audio, extension)
		rt.Secondary = strings.Replace(audio, "./music/", "", 1)
		if len(rt.Secondary) > 64 {
			rt.Secondary = rt.Secondary[0:63]
		}
		rt.Rotating = true
	} else {
		rt.Rotating = false
	}

	// Actual playing audio starts here!
	if params.Type == tools.FileType || params.Type == tools.SilenceType {
		logs.FmRadStrInfo(fmt.Sprintf("Executing %v as a child process. Output below:", name))
		err := tools.ExecCommand(name, config.GetVerbose(), GenerateOptions(params)...)
		if err != nil {
			return err
		}
	}
	// Using workaround when playing stream
	if params.Type == tools.StreamType {
		textoptions := fmt.Sprintf("sox %v -t wav - | sudo core/pi_fm_adv %v", params.Audio, strings.Join(GenerateOptions(tools.Params{Type: tools.StreamType}), " "))

		// Go is doing some weird things when using "|" in exec.Command, so we will run command through temp script
		err := tools.TextToFile(textoptions, "temp.sh")
		if err != nil {
			return err
		}
		logs.FmRadStrInfo(fmt.Sprintf("Executing streaming shell script as a child process. Output below:"))
		err = tools.ExecCommand("/bin/sh", config.GetVerbose(), "temp.sh")
		if err != nil {
			return err
		}
	}
	return nil
}

// Play kills old instance of PiFmAdv and launches new one using new parameters
func Play(params tools.Params) error {
	// Make sure that previous playback is stopped
	_, err := pkill.Pkill("pi_fm_adv", os.Interrupt)
	if err != nil {
		return err
	}
	go func() {
		err := run("core/pi_fm_adv", params)
		if err != nil {
			errorText := err.Error()
			if errorText == "exit status 2" || errorText == "signal: interrupt" {
				logs.PiFmAdvInfo("Expected PiFmAdv kill")
			} else {
				errorString := "Unexpected PiFmAdv error."
				if !config.GetVerbose() {
					errorString += " Use verbose option next time to get more information."
				}
				logs.PiFmAdvError(fmt.Sprintf("%v %v", errorString, err))
			}
		}
	}()
	return nil
}
