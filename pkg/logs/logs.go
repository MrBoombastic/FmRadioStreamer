package logs

import "github.com/BOOMfinity/golog"

var FmRadStr golog.Logger
var PiFmAdv golog.Logger
var Pkill golog.Logger

func init() {
	FmRadStr = golog.New("FmRadioStreamer")
	PiFmAdv = golog.New("PiFmAdv")
	Pkill = golog.New("pkill")
}

func FmRadStrInfo(info interface{}) {
	FmRadStr.Info().Send("%v", info)
}

func FmRadStrWarn(warn interface{}) {
	FmRadStr.Warn().Send("%v", warn)
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

func PkillInfo(info interface{}) {
	PiFmAdv.Info().Send("%v", info)
}

func PkillWarn(warn interface{}) {
	PiFmAdv.Warn().Send("%v", warn)
}

func PkillError(err interface{}) {
	PiFmAdv.Error().Send("%v", err)
}
