package config

type Config struct {
	Frequency     float64 `json:"freq"`
	Multiplier    float64 `json:"multiplier"`
	PS            string  `json:"PS"`
	RT            string  `json:"RT"`
	PI            string  `json:"PI"`
	PTY           uint    `json:"PTY"`
	YouTubeAPIKey string  `json:"apikey"`
	Port          uint16  `json:"port"`
	Power         uint8   `json:"power"`
	Mpx           uint    `json:"mpx"`
	Preemph       string  `json:"preemph"`
	Screen        bool    `json:"screen"`
}
