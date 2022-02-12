package tools

import (
	"encoding/json"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/stianeikeland/go-rpio/v4"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	urltool "net/url"
	"os"
	"time"
)

var LocalIP net.IP

func RefreshLocalIP() {
	conn, err := net.Dial("udp", "8.8.8.8:80") // It will not actually connect
	if err != nil {
		log.Println("WARNING: Failed to get local IP address! Falling back to localhost...")
		LocalIP = net.ParseIP("127.0.0.1")
	}
	defer conn.Close()
	LocalIP = conn.LocalAddr().(*net.UDPAddr).IP
}

func InitGPIO() error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	return nil
}

func StopGPIO() error {
	err := rpio.Close()
	if err != nil {
		return err
	}
	return nil
}

func SearchYouTube(query string) (YouTubeAPIResult, error) {
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/search?key=%v&q=%v&part=snippet&maxResults=1&type=video", config.GetYouTubeAPIKey(), urltool.QueryEscape(query))
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
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return YouTubeAPIResult{}, readErr
	}

	result := YouTubeAPIResult{}
	//fmt.Println(string(body))
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		return YouTubeAPIResult{}, jsonErr
	}

	return result, nil
}

func CheckRoot() {
	if os.Geteuid() != 0 {
		log.Println("WARNING: Not running as root! Exiting...")
		os.Exit(0)
	}
}
