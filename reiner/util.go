package main

import (
	"net"
	"os"
)

// GetProcessID returns the current PID of the Beacon
func GetProcessID() (pid int) {
	pid = os.Getpid()
	return
}

// GetMACAddrs gets the MAC address of the default NIC on the host
func GetMACAddrs() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
			return nil, err
	}
	var as []string
	for _, ifa := range interfaces {
			a := ifa.HardwareAddr.String()
			if a != "" {
					as = append(as, a)
			}
	}
	return as, nil
}