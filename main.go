package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type ChatServer struct {
}

var connections []net.Conn

func main() {
	// Create thread to listen for connections
	go listen()

	for {

	}

}

// Listen for, accept, and append new connections
func listen() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter port to listen on: ")
	scanner.Scan()
	ln, err := net.Listen("tcp", ":"+scanner.Text())
	defer ln.Close()
	if err != nil {
		// handle error
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		connections = append(connections, conn)
		go client(conn)
	}
}

func client(conn net.Conn) {
	for {
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// handle error
		}
		for _, e := range connections {
			if e != conn {
				fmt.Fprintln(e, status)
			}

		}
		fmt.Print(status)
	}
}
