package tools

import (
	"github.com/MrBoombastic/FmRadioStreamer/pkg/logs"
	"net"
)

// localIP stores local IP of the device
var localIP net.IP

// fetchLocalIP fetches current local IP.
func fetchLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80") // It will not actually connect
	if err != nil {
		logs.FmRadStrWarn("Failed to get local IP address! Falling back to localhost...")
		return net.ParseIP("127.0.0.1")
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP
}

// GetLocalIP returns saved local IP or fetches a new one.
func GetLocalIP() string {
	if localIP != nil {
		return localIP.String()
	}
	ip := fetchLocalIP()
	localIP = ip
	return ip.String()

}
