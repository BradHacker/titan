package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BradHacker/titan/ent"
	"github.com/BradHacker/titan/ent/action"
	"github.com/BradHacker/titan/models"
	"github.com/denisbrodbeck/machineid"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	action, err := client.Action.Create().SetCmd("whoami").SetArgs([]string(nil)).SetActionType(action.ActionType(models.Execute)).Save(ctx)
	if err != nil {
		log.Fatalf("error creating action: %v", err)
	}
	// fmt.Printf("%d: %s %s\n", action.ID, action.Cmd, strings.Join(action.Args, " "))

	machineId, err := machineid.ID()
	if err != nil {
		log.Fatalf("error getting machine id %v", err)
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("error getting hostname %v", err)
	}
	
	agent, err := client.Agent.Create().SetUUID(machineId).SetHostname(hostname).SetIP("127.0.0.1").SetPort("1337").SetPid(os.Getpid()).Save(ctx)
	if err != nil {
		log.Fatalf("error creating agent: %v", err)
	}
	// fmt.Printf("%d: %s - %s\n", agent.ID, agent.UUID, agent.Hostname)

	instruction, err := client.Instruction.Create().SetSentAt(time.Now()).SetAgent(agent).SetAction(action).Save(ctx)
	if err != nil {
		log.Fatalf("error creating instruction: %v", err)
	}
	// fmt.Printf("%d: %s\n", instruction.ID, instruction.SentAt)

	beacon, err := client.Beacon.Create().SetSentAt(time.Now()).SetInstruction(instruction).Save(ctx)
	if err != nil {
		log.Fatalf("error creating beacon: %v", err)
	}
	// fmt.Printf("%d: %s\n", beacon.ID, beacon.SentAt)
	
	fmt.Printf("Beacon: %d\n", beacon.ID)
	fmt.Printf("\tSent At: %s\n", beacon.SentAt)
	fmt.Printf("\tReceived At: %s\n", beacon.ReceivedAt)
	fmt.Printf("\tInstruction: %d\n", instruction.ID)
	fmt.Printf("\t\tSent At: %s\n", instruction.SentAt)
	fmt.Printf("\t\tAgent: %d\n", agent.ID)
	fmt.Printf("\t\t\tUUID: %s\n", agent.UUID)
	fmt.Printf("\t\t\tHostname: %s\n", agent.Hostname)
	fmt.Printf("\t\t\tIP: %s\n", agent.IP)
	fmt.Printf("\t\t\tPort: %s\n", agent.Port)
	fmt.Printf("\t\t\tPID: %d\n", agent.Pid)
	fmt.Printf("\t\tAction: %d\n", action.ID)
	fmt.Printf("\t\t\tType: %s\n", action.ActionType)
	fmt.Printf("\t\t\tCmd: %s\n", action.Cmd)
	fmt.Printf("\t\t\tArgs: %s\n", action.Args)
	fmt.Printf("\t\t\tOutput: \n")
}