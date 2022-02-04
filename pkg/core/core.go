package core

import (
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"io"
	"log"
	"os/exec"
	"syscall"
)

var PiFmAdv *exec.Cmd

func GenerateOptions(audio string) []string {
	cfg := config.Get()
	options := []string{
		"--ps", cfg.PS,
		"--rt", cfg.RT,
		"--freq", fmt.Sprintf("%f", cfg.Frequency),
		"--power", fmt.Sprintf("%v", cfg.Power),
		"--pi", cfg.PI,
		"--pty", fmt.Sprintf("%v", cfg.PTY),
		"--mpx", fmt.Sprintf("%v", cfg.Mpx),
	}
	if audio != "" {
		options = append(options, "--audio", fmt.Sprintf("./music/%v", audio))
	}
	return options
}
func run(name string, args []string) error {
	PiFmAdv = exec.Command(name, args...)
	PiFmAdv.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGINT,
	}
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
	fmt.Printf("%s\n", cmderr)
	/*
		cmdout, _ := io.ReadAll(stdout)
		fmt.Printf("%s\n", cmdout)
	*/
	return nil
}
func Kill() {
	if PiFmAdv != nil {
		PiFmAdv.Process.Kill()
		PiFmAdv = nil
	}
}
func SuperKill() {
	cmd := exec.Command("pkill", "-2", "pi_fm_adv")
	cmd.Start()

	/*	stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		cmdout, _ := io.ReadAll(stdout)
		fmt.Printf("%s\n", cmdout)*/

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
