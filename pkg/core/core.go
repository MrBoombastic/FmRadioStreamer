package core

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"io"
	"log"
	"os/exec"
)

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
	}
	if audio != "" {
		options = append(options, "--audio", fmt.Sprintf("./music/%v", audio))
	}
	return options
}
func run(name string, args []string) error {
	PiFmAdv := exec.Command(name, args...)
	stderr, err := PiFmAdv.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	// Stdout commented out for clearer command output, safe to undo!
	/*
		stdout, err := PiFmAdv.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
	*/
	err = PiFmAdv.Start()
	if err != nil {
		return err
	}
	cmderr, _ := io.ReadAll(stderr)
	if fmt.Sprintf("%s", cmderr) != "" {
		fmt.Printf("Pi_Fm_Adv error:\n%s", cmderr)
	}
	/*
		cmdout, _ := io.ReadAll(stdout)
		fmt.Printf("%s\n", cmdout)
	*/
	return nil
}

func Kill() {
	cmd := exec.Command("pkill", "-2", "pi_fm_adv")
	cmd.Start()
	cmd.Wait()
}
func Play(audio string) {
	// Make sure that previous playback is stopped
	Kill()
	go func() {
		options := GenerateOptions(audio)
		err := run("core/pi_fm_adv", options)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
