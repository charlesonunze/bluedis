package main

import (
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	log.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal("Failed to bind to port 6379", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Error accepting connection: ", err.Error())
		}

		go readLoop(conn)
	}
}

func readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		b, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error reading from client: ", err.Error())
			return
		}

		command := string(buf[:b])

		if command[0] == '*' {
			res := parseArray(command)
			cmd := strings.ToLower(res[0].(string))

			switch cmd {
			case COMMAND_PING:
				if len(res) > 1 {
					bs := encodeBulkString(mergeStrings(res))
					conn.Write([]byte(bs))
				} else {
					conn.Write([]byte("+PONG\r\n"))
				}
			case COMMAND_ECHO:
				bs := encodeBulkString(mergeStrings(res))
				conn.Write([]byte(bs))
			default:
				conn.Write([]byte("-ERR unknown command '" + cmd + "'\r\n"))
			}
		}
	}
}
