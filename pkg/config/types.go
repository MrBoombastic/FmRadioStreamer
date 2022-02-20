package config

// Config is config.json structure of all available settings
type Config struct {
	Frequency     float64 `json:"freq"`
	Multiplier    float64 `json:"multiplier"`
	PS            string  `json:"PS"`
	RT            string  `json:"RT"`
	PI            string  `json:"PI"`
	TP            string  `json:"TP"`
	PTY           uint    `json:"PTY"`
	YouTubeAPIKey string  `json:"apikey"`
	Port          uint16  `json:"port"`
	Power         uint8   `json:"power"`
	Mpx           uint    `json:"mpx"`
	Preemph       string  `json:"preemph"`
	AntennaGPIO   uint8   `json:"antennaGPIO"`
	SSD1306       bool    `json:"ssd1306"`
	DynamicRT     bool    `json:"dynamicRT"`
}
