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
	// fmt.Println("Cool we're printing some shit...\nVery nice.")
	// ln, err := net.Listen("tcp", ":9000")
	// conn, err := ln.Accept()
	// defer ln.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// io.WriteString(conn, fmt.Sprint("Hello World \n", time.Now(), "\n"))

	// buf := make([]byte, 1024)
	// conn.Read(buf)
	// conn.Write([]byte("Writing to connection..."))

	// Create thread to listen for connections
	go listen()

	i := 0

	for {
		i++
		if i%10000000000 == 0 {
			for _, e := range connections {
				fmt.Fprintln(e, i)
			}
		}
	}

}

// Listen for, accept, and append new connections
func listen() {
	ln, err := net.Listen("tcp", ":8080")
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
	bufio.NewReader(conn).ReadString('\n')
	for {
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// handle error
		}
		fmt.Print(status)
	}
}
