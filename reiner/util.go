package main

import (
	"net"
	"os"
	"strings"
	"time"

	"github.com/BradHacker/titan/models"
	"github.com/denisbrodbeck/machineid"
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

func GenerateHeartbeat(connection net.Conn) (heartbeat models.Heartbeat) {
	var agent models.Agent

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	} else {
		agent.Hostname = hostname
	}
	localAaddressParts := strings.Split(connection.LocalAddr().String(), ":")
	agent.IP = localAaddressParts[0]
	agent.PID = GetProcessID()
	agent.Port = localAaddressParts[1]
	macId, err := machineid.ID();
	if err != nil {
		agent.UUID = "unknown"
	} else {
		agent.UUID = macId
	}

	heartbeat.SentAt = time.Now()
	heartbeat.Agent = agent
	return
}