package tools

import (
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"net"
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
	return nil
}
