package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ActionType is the kind of action to be executed
type ActionType string

const (
	// Execute is an action that does RCE
	Execute ActionType = "EXEC"
)

// Action is an action to be executed by the agent
type Action struct {
	ActionType ActionType `json:"actionType"`
	Cmd string `json:"cmd"`
	Args []string `json:"args"`
	Output string `json:"output"`
}

// Agent represents the agent running on the target
type Agent struct {
	UUID string `json:"uuid"`
	Hostname string `json:"hostname"`
	IP string `json:"ip"`
	Port string `json:"port"`
	PID int `json:"pid"`
}

type Heartbeat struct {
	SentAt time.Time `json:"sentAt"`
	Agent Agent `json:"agent"`
}

// Instruction is a command sent by the C2 to the agent
type Instruction struct {
	Agent Agent `json:"agent"`
	Action Action `json:"action"`
	SentAt time.Time `json:"sentAt"`
}

// Beacon is the object sent back from the agent to the C2 when beaconing back
type Beacon struct {
	SentAt time.Time `json:"sentAt"`
	ReceivedAt time.Time `json:"receivedAt"`
	Instruction Instruction `json:"instruction"`
}

func (b Beacon) String() string {
	return "----Beacon----" +
		"\nAgent" + 
		"\n-----" + 
		"\nHost\t" +
		b.Instruction.Agent.Hostname +
		"\nIP\t" +
		b.Instruction.Agent.IP +
		"\nPort\t" +
		b.Instruction.Agent.Port +
		"\nPID\t" +
		fmt.Sprintf("%d", b.Instruction.Agent.PID) +
		"\nInstruction" + 
		"\n-----------" +
		"\nAction\t" +
		string(b.Instruction.Action.ActionType) +
		"\nCmd\t" +
		b.Instruction.Action.Cmd +
		"\nArgs\t" +
		strings.Join(b.Instruction.Action.Args, " ") +
		"\nOutput" +
		"\n------\n" +
		b.Instruction.Action.Output +
		"\nMeta" +
		"\n----" +
		"\nSent\t" +
		b.SentAt.String() +
		"\nRecvd\t" +
		b.ReceivedAt.String() +
		"\n"
}

// EncodeBeacon encodes a beacon response as a byte array
func EncodeBeacon(beacon Beacon) (beaconBytes []byte, err error) {
	beaconBytes, err = json.Marshal(beacon)
	return
}

// DecodeBeacon decodes an incoming data packet into tho Beacon struct
func DecodeBeacon(data []byte) (beacon *Beacon, err error) {
	err = json.Unmarshal(data, &beacon)
	return
}

// EncodeInstruction encodes an instruction as a JSON object in the form of a byte array
func EncodeInstruction(instruction Instruction) (instructionBytes []byte, err error) {
	instructionBytes, err = json.Marshal(instruction)
	return
}

// DecodeInstruction decodes instructions sent from the C2
func DecodeInstruction(data []byte) (instruction *Instruction, err error) {
	err = json.Unmarshal(data, &instruction)
	return
}

// EncodeBeacon encodes a beacon response as a byte array
func EncodeHeartbeat(heartbeat Heartbeat) (heartbeatBytes []byte, err error) {
	heartbeatBytes, err = json.Marshal(heartbeat)
	return
}

// DecodeHeartbeat decodes instructions sent from the C2
func DecodeHeartbeat(data []byte) (heartbeat *Heartbeat, err error) {
	err = json.Unmarshal(data, &heartbeat)
	return
}