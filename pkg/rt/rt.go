package rt

import (
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"time"
)

var currentRTState = 0

// Primary is used to display RT - both static and dynamic.
var Primary string

// Secondary is only used when dynamic RT is enabled.
var Secondary string

// Rotating switches dynamic RT on or off.
var Rotating = false

// Rotate enables switching RT between that saved in config and current playing audio filename.
func Rotate(interval uint) error {
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
		if Rotating {
			if currentRTState == 0 {
				_, err := f.WriteString("AB ON\n")
				if err != nil {
					return errors.New("cannot update dynamic RT text (A/B flag)")
				}
				_, err = f.WriteString(fmt.Sprintf("RT %s", Primary))
				if err != nil {
					return errors.New("cannot update dynamic RT text")
				}
				currentRTState++
			} else {
				_, err := f.WriteString("AB OFF\n")
				if err != nil {
					return errors.New("ERROR: Cannot update dynamic RT text (A/B flag)")
				}
				_, err = f.WriteString(fmt.Sprintf("RT %s", Secondary))
				if err != nil {
					return errors.New("ERROR: Cannot update dynamic RT text")
				}
				currentRTState--
			}
		}
		time.Sleep(time.Second * time.Duration(interval))
	}
}
