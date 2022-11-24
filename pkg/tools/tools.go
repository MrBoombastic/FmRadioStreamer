package tools

import (
	"encoding/json"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/stianeikeland/go-rpio/v4"
	"io"
	"net/http"
	urltool "net/url"
	"time"
)

// InitGPIO opens connection for LEDs and buttons
func InitGPIO() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	return nil
}

// StopGPIO closes connection for LEDs and buttons. This function crashes RPi, therefore is not used internally.
func StopGPIO() error {
	err := rpio.Close()
	if err != nil {
		return err
	}
	return nil
}

// SearchYouTube queries YouTube API and returns first found video
func SearchYouTube(query string, apikey string) (YouTubeAPIResult, error) {
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/search?key=%v&q=%v&part=snippet&maxResults=1&type=video", apikey, urltool.QueryEscape(query))
	client := http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return YouTubeAPIResult{}, err
	}
	res, getErr := client.Do(req)
	if getErr != nil {
		return YouTubeAPIResult{}, getErr
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return YouTubeAPIResult{}, readErr
	}

	result := YouTubeAPIResult{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		return YouTubeAPIResult{}, jsonErr
	}

	return result, nil
}

// ConfigToMap converts given config to map. Needs mutex locked!
func ConfigToMap(cfg *config.SafeConfig) map[string]interface{} {
	var data map[string]interface{}
	raw, _ := json.Marshal(cfg.Config)
	_ = json.Unmarshal(raw, &data)
	return data
}
