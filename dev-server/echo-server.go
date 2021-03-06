package dev_server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	listenHost = "localhost"
	listenPort = "1337"
	listenType = "tcp"
)

func main() {
	fmt.Printf("Starting %s server on %s:%s\n", listenType, listenHost, listenPort)
	l, err := net.Listen(listenType, listenHost+":"+listenPort)

	if err != nil {
		fmt.Printf("Error listening on %s:%s", listenHost, listenPort)
		os.Exit(1)
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting: ", err.Error())
			return
		}
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client dropped the connection")
		conn.Close()
		return
	}

	log.Println(string(buffer[:len(buffer)-1]))

	conn.Write(buffer)

	handleConnection(conn)
}
