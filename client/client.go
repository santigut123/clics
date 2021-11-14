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
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	for {
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// handle error
		}
		fmt.Println(status)
	}
}
