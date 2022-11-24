package tools

import (
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"github.com/arduino/go-apt-client"
	"os"
	"os/exec"
	"strconv"
)

// CheckRoot checks if user has root permissions.
func CheckRoot() bool {
	return os.Geteuid() == 0
}

// CheckLibsndfileVersion returns truncated libsndfile1-dev version.
func CheckLibsndfileVersion() (float64, error) {
	lib, err := apt.Search("libsndfile1-dev")
	if err != nil || len(lib) < 1 {
		return 0, err
	}
	ver := lib[0].Version
	floatVer, err := strconv.ParseFloat(ver[0:3], 64)
	if err != nil {
		return 0, err
	}
	return floatVer, nil
}

// TextToFile saves any given text to given file overwriting it.
func TextToFile(text string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logs.FmRadStrError(err)
			return
		}
	}(file)

	_, err = file.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

// ExecCommand is just exec.Command wrapper.
func ExecCommand(name string, verbose bool, args ...string) error {
	cmd := exec.Command(name, args...)
	// Redirecting stdout and stdin from child process to master process
	if verbose {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
