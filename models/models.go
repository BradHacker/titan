package models

import "time"

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