package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379", err)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go readLoop(conn)
	}
}

func readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		if _, err := conn.Read(buf); err != nil {
			if err == io.EOF {
				return
			}

			fmt.Println("error reading from client: ", err.Error())
			return
		}

		conn.Write([]byte("+PONG\r\n"))
	}
}
