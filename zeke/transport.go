package main

import (
	"encoding/json"
	"net"

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
	dataBytes, jsonErr := EncodeInstruction(instruction)
	if jsonErr != nil {
		return jsonErr
	}
	_, err = connection.Write(dataBytes)
	return
}

// DecodeBeacon decodes an incoming data packet into tho Beacon struct
func DecodeBeacon(data []byte) (beacon *models.Beacon, err error) {
	err = json.Unmarshal(data, &beacon)
	return
}