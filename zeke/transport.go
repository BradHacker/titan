package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/BradHacker/titan/models"
)

// CreateTCPListener creates a listener for TCP connections on the given listenPort on the listenHost
func CreateTCPListener(listenHost string, listenPort string) (listener net.Listener, err error) {
	listener, err = net.Listen("tcp", listenHost+":"+listenPort)
	return
}

// CloseListener closes a net listener
func CloseListener(listener net.Listener) (err error) {
	err = listener.Close()
	return
}

// CloseConnection closes an existing net.Conn to the C2
func CloseConnection(connection net.Conn) (err error) {
	err = connection.Close()
	return
}

// AcceptConnection accepts an incoming connection from a Beacon
func AcceptConnection(listener net.Listener) (conn net.Conn, err error) {
	conn, err = listener.Accept()
	return
}

// EncodeInstruction encodes an instruction as a JSON object in the form of a byte array
func EncodeInstruction(instruction models.Instruction) (instructionBytes []byte, err error) {
	instructionBytes, err = json.Marshal(instruction)
	return
}

// SendInstruction sends an instruction struct via TCP over a connection
func SendInstruction(instruction models.Instruction, connection net.Conn) (err error) {
	instruction.SentAt = time.Now()
	dataBytes, jsonErr := EncodeInstruction(instruction)
	if jsonErr != nil {
		return jsonErr
	}
	_, err = connection.Write(append(dataBytes, []byte{0xff}...))
	return
}

// DecodeBeacon decodes an incoming data packet into tho Beacon struct
func DecodeBeacon(data []byte) (beacon *models.Beacon, err error) {
	err = json.Unmarshal(data, &beacon)
	return
}

// WaitForBeaconReturn waits for a Beacon response after sending an instruction to the Beacon
func WaitForBeaconReturn(connection net.Conn) (err error) {
	buffer, err := bufio.NewReader(connection).ReadBytes(0xff)
	if err != nil {
		fmt.Println("Client dropped the connection")
		return
	}

	if len(buffer) > 0 {
    buffer = buffer[:len(buffer)-1]
	}

	receivedBeacon, err := DecodeBeacon(buffer)
	if err != nil {
		fmt.Printf("Couldn't decode beacon: %v", err)
		return
	}
	fmt.Printf(receivedBeacon.String())
	return nil
}