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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter ip:port : ")
	scanner.Scan()

	conn, err := net.Dial("tcp", scanner.Text())
	if err != nil {
		// handle error
		panic(err)
	}
	fmt.Println("Connected to " + scanner.Text())
	connection = conn
	go listen()

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
