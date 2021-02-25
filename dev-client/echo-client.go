package dev_client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "localHost"
	connPort = "1337"
	connType = "tcp"
)

func main() {
	conn, err := net.Dial(connType, connHost+":"+connPort)

	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Message: ")

		input, _ := reader.ReadBytes('\n')

		conn.Write(input)

		message, _ := bufio.NewReader(conn).ReadString('\n')

		log.Print("Server echo: ", message)
	}
}
