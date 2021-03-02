package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	listenHost = "localhost"
	listenPort = "1337"
)

func main() {
	listener, err := CreateTCPListener(listenHost, listenPort)
	if err != nil {
		log.Fatal(fmt.Printf("Error listening on %s:%s", listenHost, listenPort))
	}

	defer CloseListener(listener)

	for {
		c, err := AcceptConnection(listener)
		if err != nil {
			fmt.Println("Error connecting: ", err.Error())
			return
		}
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer, err := bufio.NewReader(conn).ReadBytes(0xff)

	if err != nil {
		fmt.Println("Client dropped the connection")
		return
	}

	if len(buffer) > 0 {
    buffer = buffer[:len(buffer)-1]
}

	receivedBeacon, err := DecodeBeacon(buffer)
	if err != nil {
		fmt.Printf("Couldn't decode beacon: %v", err)
	}
	fmt.Printf(receivedBeacon.String())
}