package logs

import "github.com/BOOMfinity/golog"

var FmRadStr golog.Logger
var PiFmAdv golog.Logger

func init() {
	FmRadStr = golog.New("FmRadioStreamer")
	PiFmAdv = golog.New("PiFmAdv")
}

func FmRadStrInfo(info interface{}) {
	FmRadStr.Info().Send("%v", info)
}

func FmRadStrFatal(warn interface{}) {
	FmRadStr.Fatal().Send("%v", warn)
}

func FmRadStrError(err interface{}) {
	FmRadStr.Error().Send("%v", err)
}

func PiFmAdvInfo(info interface{}) {
	PiFmAdv.Info().Send("%v", info)
}

func PiFmAdvWarn(warn interface{}) {
	PiFmAdv.Warn().Send("%v", warn)
}

func PiFmAdvError(err interface{}) {
	PiFmAdv.Error().Send("%v", err)
}

// PiFmAdvHandleError uses PiFmAdvError and returns true, if app should return
func PiFmAdvHandleError(err error) bool {
	if err != nil {
		PiFmAdvError(err)
		return true
	}
	return false
}
