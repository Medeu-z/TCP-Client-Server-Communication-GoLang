package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	connection, err := net.Dial("tcp", "localhost:8000")
	logFatal(err)

	defer connection.Close()

	fmt.Println("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')

	logFatal(err)

	username = strings.Trim(username, " \r\n")

	helloMsg := fmt.Sprintf("Hello, %s.\n", username)

	connection.Write([]byte(helloMsg))

	server_reader := bufio.NewReader(connection)
	message, err := server_reader.ReadString('\n')

	fmt.Println(message)

}
