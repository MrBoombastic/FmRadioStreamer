package config

type Config struct {
	Frequency     float32 `json:"freq"`
	PS            string  `json:"PS"`
	RT            string  `json:"RT"`
	YouTubeAPIKey string  `json:"apikey"`
	Port          uint16  `json:"port"`
	Power         uint8   `json:"power"`
	Screen        bool    `json:"screen"`
}
