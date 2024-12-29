package core

import (
	"fmt"
	"net"
	"os"
)

func GetServerAddress() (string, error) {
	envAddress := os.Getenv("SERVER_ADDRESS")
	if envAddress != "" {
		return envAddress, nil
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", fmt.Errorf("failed to get server IP address: %w", err)
	}

	for _, address := range addrs {
		// Kiểm tra loại địa chỉ
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String(), nil
		}
	}

	return "", fmt.Errorf("could not determine server IP address")
}
