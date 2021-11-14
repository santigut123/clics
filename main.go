package main

import (
	"bufio"
	"fmt"
	"net"
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
	ln, err := net.Listen("tcp", ":5000")
	defer ln.Close()
	if err != nil {
		// handle error
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println(conn)
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
	}
}
