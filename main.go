package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type ChatServer struct {
}

func main() {
	fmt.Println("Cool we're printing some shit...\nVery nice.")
	ln, err := net.Listen("tcp", ":9000")
	conn, err := ln.Accept()
	defer ln.Close()
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, fmt.Sprint("Hello World \n", time.Now(), "\n"))

	buf := make([]byte, 1024)
	conn.Read(buf)
	conn.Write([]byte("Writing to connection..."))

}
