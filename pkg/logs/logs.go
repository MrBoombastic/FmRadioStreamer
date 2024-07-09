package logs

import "github.com/BOOMfinity/golog"

var FmRadStrLog golog.Logger
var PiFmAdvLog golog.Logger

func init() {
	FmRadStrLog = golog.New("FmRadioStreamer")
	PiFmAdvLog = golog.New("PiFmAdv")
}

func FmRadStrInfo(info interface{}) {
	FmRadStrLog.Info().Send("%v", info)
}

func FmRadStrWarn(err interface{}) {
	FmRadStrLog.Error().Send("%v", err)
}

func FmRadStrFatal(warn interface{}) {
	FmRadStrLog.Fatal().Send("%v", warn)
}

func FmRadStrError(err interface{}) {
	FmRadStrLog.Error().Send("%v", err)
}

func PiFmAdvInfo(info interface{}) {
	PiFmAdvLog.Info().Send("%v", info)
}

func PiFmAdvError(err interface{}) {
	PiFmAdvLog.Error().Send("%v", err)
}
