package models

import (
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
	UUID []byte `json:"uuid"`
	Hostname string `json:"hostname"`
	IP string `json:"ip"`
	Port string `json:"port"`
	PID int `json:"pid"`
}

// Instruction is a command sent by the C2 to the agent
type Instruction struct {
	Agent Agent `json:"agent"`
	Action Action `json:"action"`
	SentAt time.Time `json:"sentAt"`
}

// Beacon is the object sent back from the agent to the C2 when beaconing back
type Beacon struct {
	Agent Agent `json:"agent"`
	Action Action `json:"action"`
	SentAt time.Time `json:"sentAt"`
	ReceivedAt time.Time `json:"receivedAt"`
	Instruction Instruction `json:"instruction"`
}

func (b Beacon) String() string {
	return "----Beacon----" +
		"\nAgent" + 
		"\n-----" + 
		"\nHost\t" +
		b.Agent.Hostname +
		"\nIP\t" +
		b.Agent.IP +
		"\nPort\t" +
		b.Agent.Port +
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