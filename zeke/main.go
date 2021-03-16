package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/BradHacker/titan/ent"
	"github.com/BradHacker/titan/ent/agent"
	"github.com/BradHacker/titan/models"
	_ "github.com/mattn/go-sqlite3"
)

const (
	listenHost = "localhost"
	listenPort = "1337"
)

func main() {
	client, err := ent.Open("sqlite3", "file:test.sqlite?_loc=auto&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	fmt.Println("ENT Databasse Initialized...")
	
	defer client.Close()
	ctx := context.Background()

	// Auto migrate the database
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	listener, err := CreateTCPListener(listenHost, listenPort)
	if err != nil {
		log.Fatal(fmt.Printf("Error listening on %s:%s", listenHost, listenPort))
	}
	fmt.Printf("Listening on %s:%s\n", listenHost, listenPort)

	defer CloseListener(listener)

	for {
		c, err := AcceptConnection(listener)
		if err != nil {
			fmt.Println("Error connecting: ", err.Error())
			return
		}
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
		go handleConnection(ctx, client, c)
	}
}

func handleConnection(ctx context.Context, client *ent.Client, conn net.Conn) {
	defer CloseConnection(conn)

	// agent, err := client.Agent.Query()

	heartbeat, err := WaitForBeaconHeartbeat(conn)
	if err != nil {
		log.Fatalf("error while waiting for Beacon heartbeat: %v", err)
	}
	agent, err := client.Agent.Query().Where(agent.And(agent.UUIDEQ((*heartbeat).Agent.UUID), agent.HostnameEQ((*heartbeat).Agent.Hostname))).Only(ctx)
	if err != nil {
		fmt.Printf("No agent found for %s (%s), creating one...", (*heartbeat).Agent.Hostname, (*heartbeat).Agent.UUID);
		agent, err = client.Agent.Create().
			SetUUID((*heartbeat).Agent.UUID).
			SetHostname((*heartbeat).Agent.Hostname).
			SetIP((*heartbeat).Agent.IP).
			SetPort((*heartbeat).Agent.Port).
			SetPid((*heartbeat).Agent.PID).
			Save(ctx)
		if err != nil {
			log.Fatalf("couldn't create agent in database: %v\n", err)
		}
	}
	fmt.Printf("Got heartbeat from %s\n", agent.Hostname)

	// testInstruction := generateTestInstruction()
	// err = SendInstruction(testInstruction, conn)
	// if err != nil {
	// 	log.Fatalf("something went wrong while sending instruction: %v\n", err)
	// }

	// err = WaitForBeaconReturn(conn)
	// if err != nil {
	// 	log.Fatalf("error while waiting for Beacon response: %v", err)
	// }
}

func generateTestInstruction() (instruction models.Instruction) {
	var testAction models.Action

	testAction.ActionType = "EXEC"
	input := ""
	fmt.Println("Enter command to run:")
	fmt.Scanln(&input)
	parts := strings.Split(input, " ")
	fmt.Println(parts)
	testAction.Cmd = parts[0]
	testAction.Args = parts[1:]

	instruction.Action = testAction
	return
}