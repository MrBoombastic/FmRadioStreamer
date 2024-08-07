package config

import "sync"

// Config is config.json structure of all available settings.
type Config struct {
	Frequency         float64 `json:"freq"`
	Format            string  `json:"format"`
	Multiplier        float64 `json:"multiplier"`
	PS                string  `json:"PS"`
	RT                string  `json:"RT"`
	PI                string  `json:"PI"`
	TP                string  `json:"TP"`
	PTY               uint    `json:"PTY"`
	YouTubeAPIKey     string  `json:"ytApiKey"`
	Port              uint16  `json:"port"`
	Power             uint8   `json:"power"`
	Mpx               uint    `json:"mpx"`
	Preemph           string  `json:"preemph"`
	AntennaGPIO       uint8   `json:"antennaGPIO"`
	SSD1306           bool    `json:"ssd1306"`
	DynamicRT         bool    `json:"dynamicRT"`
	DynamicRTInterval uint    `json:"dynamicRTInterval"`
	Verbose           bool    `json:"verbose"`
	Ytdlp             bool    `json:"ytdlp"`
}

// SafeConfig is just Config combined with mutex.
type SafeConfig struct {
	sync.Mutex
	Config
}
