package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var connection net.Conn
var name string

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
	}
	connection = conn
	go listen()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()

	name = scanner.Text() + ": "

	for {
		fmt.Print(name)
		scanner.Scan()
		fmt.Fprintln(conn, name+scanner.Text())
	}
}

func listen() {
	for {
		status, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			// handle error
		}
		fmt.Print(status)
	}
}
