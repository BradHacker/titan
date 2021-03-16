package models

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/BradHacker/titan/ent"
	"github.com/BradHacker/titan/ent/action"
	"github.com/BradHacker/titan/ent/agent"
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
	ReceivedAt time.Time `json:"receivedAt"`
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

func CreateAgent(ctx context.Context, client *ent.Client, agent Agent) (dbAgent *ent.Agent, err error) {
	dbAgent, err = client.Agent.Create().
		SetUUID(agent.UUID).
		SetHostname(agent.Hostname).
		SetIP(agent.IP).
		SetPort(agent.Port).
		SetPid(agent.PID).
		Save(ctx)
	return
}

func FindOrCreateAgent(ctx context.Context, client *ent.Client, a Agent) (dbAgent *ent.Agent, err error) {
	dbAgent, err = client.Agent.Query().Where(agent.And(agent.UUIDEQ(a.UUID), agent.HostnameEQ(a.Hostname))).Only(ctx)
	if err != nil {
		fmt.Printf("No agent found for %s (%s), creating one...\n", a.Hostname, a.UUID);
		dbAgent, err = CreateAgent(ctx, client, a)
	}
	if err != nil {
		fmt.Println(fmt.Errorf("couldn't create agent in database: %v", err))
	}
	return
}


func CreateHeartbeat(ctx context.Context, client *ent.Client, heartbeat Heartbeat) (dbHeartbeat *ent.Heartbeat, err error) {
	dbAgent, err := FindOrCreateAgent(ctx, client, heartbeat.Agent)
	if err != nil {
		fmt.Println(fmt.Errorf("error finding/creating agent: %v", err))
		return
	}
	dbHeartbeat, err = client.Heartbeat.Create().
		SetSentAt(heartbeat.SentAt).
		SetReceivedAt(heartbeat.ReceivedAt).
		SetAgent(dbAgent).
		Save(ctx)
	return
}

func CreateAction(ctx context.Context, client *ent.Client, a Action) (dbAction *ent.Action, err error) {
	dbAction, err = client.Action.Create().
		SetActionType(action.ActionType(a.ActionType)).
		SetCmd(a.Cmd).
		SetArgs(a.Args).
		SetOutput(a.Output).
		Save(ctx)
	return
}

func CreateInstruction(ctx context.Context, client *ent.Client, instruction Instruction) (dbInstruction *ent.Instruction, err error) {
	dbAction, err := CreateAction(ctx, client, instruction.Action)
	if err != nil {
		fmt.Println(fmt.Errorf("couldn't create action in db: %v", err))
		return
	}
	dbAgent, err := FindOrCreateAgent(ctx, client, instruction.Agent)
	if err != nil {
		fmt.Println(fmt.Errorf("error finding/creating agent: %v", err))
		return
	}
	dbInstruction, err = client.Instruction.Create().
		SetAction(dbAction).
		SetAgent(dbAgent).
		SetSentAt(time.Now()).
		Save(ctx)
	return
}