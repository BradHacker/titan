package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"time"

	"github.com/BradHacker/titan/models"
)

func main() {
	testBeacon := generateTestBeacon()

	connection, err := CreateTCPConnection("localhost", "1337")
	if err != nil {
		log.Fatal(fmt.Errorf("Couldn't establish connection: %v", err))
	}

	defer CloseConnection(connection)

	err = SendBeacon(testBeacon, connection)
	if err != nil {
		log.Fatal(fmt.Errorf("Couldn't send beacon: %v", err))
	}
}

func generateTestBeacon() (beacon models.Beacon) {
	var testAction models.Action
	var testInstruction models.Instruction
	var testAgent models.Agent

	testAgent.Hostname = "Evil-Laptop1"
	testAgent.IP = "111.111.111.111"
	testAgent.PID = GetProcessID()
	testAgent.Port = "1337"
	uuidHasher := sha256.New()
	uuidHasher.Write([]byte("yeet"))
	testAgent.UUID = uuidHasher.Sum(nil)

	testAction.ActionType = "EXEC"
	testAction.Cmd = "whoami"
	testAction.Output = "bob"

	testInstruction.Action = testAction
	testInstruction.Agent = testAgent
	testInstruction.SentAt = time.Now()
	
	beacon.Action = testAction
	beacon.Agent = testAgent
	beacon.Instruction = testInstruction
	beacon.ReceivedAt = time.Now()
	beacon.SentAt = time.Now()
	return
}