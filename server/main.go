package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	openConnections = make(map[net.Conn]bool)
	newConnection   = make(chan net.Conn)
	deadConnection  = make(chan net.Conn)
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	logFatal(err)

	defer ln.Close()

	go func() {
		for {
			conn, err := ln.Accept()
			logFatal(err)

			openConnections[conn] = true
			newConnection <- conn
		}

	}()

	connection := <-newConnection

	reader := bufio.NewReader(connection)

	message, err := reader.ReadString('\n')

	logFatal(err)

	fmt.Println(message)

	msg := "Message from server: " + message
	for item := range openConnections {
		if item == connection {
			item.Write([]byte(msg))
		}
	}
}
