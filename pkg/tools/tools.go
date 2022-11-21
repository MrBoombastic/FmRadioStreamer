package tools

import (
	"encoding/json"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"github.com/arduino/go-apt-client"
	"github.com/stianeikeland/go-rpio/v4"
	"io"
	"net"
	"net/http"
	urltool "net/url"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// LocalIP stores local IP of the device
var LocalIP net.IP

// RefreshLocalIP fetches current local IP and saves it to LocalIP variable
func RefreshLocalIP() {
	conn, err := net.Dial("udp", "8.8.8.8:80") // It will not actually connect
	if err != nil {
		logs.PiFmAdvWarn("Failed to get local IP address! Falling back to localhost...")
		LocalIP = net.ParseIP("127.0.0.1")
	}
	defer conn.Close()
	LocalIP = conn.LocalAddr().(*net.UDPAddr).IP
}

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

// CheckRoot checks if user has root permissions
func CheckRoot() bool {
	return os.Geteuid() == 0
}

func CheckLibsndfileVersion() (float64, error) {
	lib, err := apt.Search("libsndfile1-dev")
	if err != nil || len(lib) < 1 {
		return 0, err
	}
	ver := lib[0].Version
	floatVer, err := strconv.ParseFloat(ver[0:3], 64)
	if err != nil {
		return 0, err
	}
	return floatVer, nil
}

func TextToFile(text string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logs.FmRadStrError(err)
			return
		}
	}(file)

	_, err = file.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

func ExecCommand(name string, verbose bool, args ...string) error {
	cmd := exec.Command(name, args...)
	// Redirecting stdout and stdin from child process to master process
	if verbose {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func ConfigToMap(cfg *config.SafeConfig) map[string]interface{} {
	var data map[string]interface{}
	raw, _ := json.Marshal(cfg.Config)
	fmt.Println(raw)
	_ = json.Unmarshal(raw, &data)
	return data
}
