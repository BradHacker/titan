package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"time"

	"github.com/BradHacker/titan/ent"
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


// SendInstruction sends an instruction struct via TCP over a connection
func SendInstruction(ctx context.Context, instruction models.Instruction, dbInstruction *ent.Instruction, connection net.Conn) (err error) {
	instruction.SentAt = time.Now()
	dataBytes, jsonErr := models.EncodeInstruction(instruction)
	if jsonErr != nil {
		return jsonErr
	}
	_, err = connection.Write(append(dataBytes, []byte{0xff}...))
	dbInstruction.Update().SetSentAt(instruction.SentAt).Save(ctx)
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

	receivedBeacon, err := models.DecodeBeacon(buffer)
	if err != nil {
		fmt.Printf("Couldn't decode beacon: %v", err)
		return
	}
	fmt.Printf(receivedBeacon.String())
	return nil
}

// WaitForBeaconHeartbeat waits for a Beacon to send it's heartbeat after initializing a connection
func WaitForBeaconHeartbeat(connection net.Conn) (heartbeat *models.Heartbeat, err error) {
	buffer, err := bufio.NewReader(connection).ReadBytes(0xff)
	if err != nil {
		fmt.Println("Client dropped the connection")
		return
	}

	if len(buffer) > 0 {
    buffer = buffer[:len(buffer)-1]
	}

	heartbeat, err = models.DecodeHeartbeat(buffer)
	return
}