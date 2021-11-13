package main

import (
	"fmt"
	"net"
	"io"
	"time"
)
type ChatServer struct{

}

func main() {
	fmt.Println("Cool we're printing some shit...\nVery nice.")
	ln,err := net.Listen("tcp",":9000")
	conn,err := ln.Accept()
	defer ln.Close()
	if err != nil{
		panic(err)
	}
	io.WriteString(conn,fmt.Sprint("Hello World \n",time.Now(),"\n"))
}
