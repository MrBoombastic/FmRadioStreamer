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
	url2 "net/url"
	"time"
)

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80") // Will not actually connect
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
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
	fmt.Println("GPIO closed")
	return nil
}

func SearchYouTube(query string) YouTubeAPIResult {
	url := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/search?key=%v&q=%v&part=snippet&maxResults=1&type=video", config.GetYouTubeAPIKey(), url2.QueryEscape(query))
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	result := YouTubeAPIResult{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return result
}
