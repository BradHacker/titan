package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/BradHacker/titan/ent"
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
	
	heartbeat, err := WaitForBeaconHeartbeat(conn)
	if err != nil {
		log.Fatalf("error while waiting for Beacon heartbeat: %v\n", err)
	}
	
	dbHeartbeat, err := models.CreateHeartbeat(ctx, client, *heartbeat)
	if err != nil {
		log.Fatalf("couldn't create heartbeat in database: %v\n", err)
	}
	fmt.Printf("Got heartbeat (ID: %d) from %s\n", dbHeartbeat.ID, heartbeat.Agent.Hostname)
	
	testInstruction := generateTestInstruction()
	dbInstruction, err := models.CreateInstruction(ctx, client, testInstruction)
	if err != nil {
		log.Fatalf("couldn't create instruction in database: %v\n", err)
	}
	fmt.Printf("Created instruction with ID %d\n", dbInstruction.ID)
	err = SendInstruction(ctx, testInstruction, dbInstruction, conn)
	if err != nil {
		log.Fatalf("something went wrong while sending instruction: %v\n", err)
	}

	err = WaitForBeaconReturn(conn)
	if err != nil {
		log.Fatalf("error while waiting for Beacon response: %v", err)
	}
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