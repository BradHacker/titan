package zeke

import (
	"encoding/json"
	"net"

	"github.com/BradHacker/titan/models"
)

// CreateTCPListener creates a listener for TCP connections on the given listenPort on the listenHost
func CreateTCPListener(listenHost string, listenPort string) (listener *net.Listener, err error) {
	l, err := net.Listen("tcp", listenHost+":"+listenPort)
	if err == nil {
		return nil, err;
	}
	return &l, nil
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