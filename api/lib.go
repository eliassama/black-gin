package api

import (
	"fmt"
	"net"
)

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}
	return "", fmt.Errorf("no local IP found")
}

func checkPortAvailable(address string) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	_ = ln.Close()
	return nil
}
