// package client

// type Client struct {
// 	port int
// }

// func MakeClient() *Client {
// 	newClient := Client{port: 1}
// 	return &newClient
// }
// func (c *Client) SetPort(port int) {
// 	c.port = port
// }

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var connection net.Conn

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
	}
	connection = conn
	go listen()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		fmt.Fprintln(conn, scanner.Text())
	}
}

func listen() {

	fmt.Fprintf(connection, "GET / HTTP/1.0\r\n\r\n")

	for {
		status, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			// handle error
		}
		fmt.Print(status)
	}
}
