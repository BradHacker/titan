package main

import (
	"crypto/sha256"
	"fmt"
	"log"
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

	testBeacon := generateTestBeacon(instruction)
	time.Sleep(1000)

	err = SendBeacon(testBeacon, connection)
	if err != nil {
		log.Fatal(fmt.Errorf("Couldn't send beacon: %v", err))
	}
}

func generateTestBeacon(instruction models.Instruction) (beacon models.Beacon) {
	var testAgent models.Agent

	testAgent.Hostname = "Evil-Laptop1"
	testAgent.IP = "111.111.111.111"
	testAgent.PID = GetProcessID()
	testAgent.Port = "1337"
	uuidHasher := sha256.New()
	uuidHasher.Write([]byte("yeet"))
	testAgent.UUID = uuidHasher.Sum(nil)

	instruction.Action.Output = "bob"

	instruction.Agent = testAgent
	
	beacon.Action = instruction.Action
	beacon.Agent = testAgent
	beacon.Instruction = instruction
	beacon.ReceivedAt = time.Now()
	return
}