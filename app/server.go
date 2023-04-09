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
		b, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}

			fmt.Println("error reading from client: ", err.Error())
			return
		}

		command := string(buf[:b])

		if command[0] == '*' {
			res := parseArray(command)
			if res[0].(string) == "ECHO" {
				bs := encodeBulkString(res[1].(string))
				conn.Write([]byte(bs))
			}
		}

		// conn.Write([]byte("+PONG\r\n"))
	}
}
