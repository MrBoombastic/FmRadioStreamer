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
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// LocalIP stores local IP of the devices
var LocalIP net.IP

// RefreshLocalIP fetches current local IP and saves it to LocalIP variable
func RefreshLocalIP() {
	conn, err := net.Dial("udp", "8.8.8.8:80") // It will not actually connect
	if err != nil {
		log.Println("WARNING: Failed to get local IP address! Falling back to localhost...")
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

// CheckRoot checks if user has root permissions. If not, exits application.
func CheckRoot() {
	if os.Geteuid() != 0 {
		log.Println("WARNING: Not running as root! Exiting...")
		os.Exit(0)
	}
}

func CheckLibsndfileVersion() (float64, error) {
	out, err := exec.Command("sudo", "apt-cache", "policy", "libsndfile1-dev").Output()
	if err != nil {
		return 0, err
	}
	// Trimming unnecessary output
	dirtyVersion := strings.Split(string(out), "\n")[1]
	dirtyVersion = strings.TrimSpace(dirtyVersion)
	dirtyVersion = strings.Split(dirtyVersion, ": ")[1]
	MMP := strings.Split(dirtyVersion, ".") //Major, Minor, Patches
	// Compiling Major and Minor to float
	version, err := strconv.ParseFloat(fmt.Sprintf("%v.%v", MMP[0], MMP[1]), 4)
	if err != nil {
		return 0, err
	}
	return version, nil
}
