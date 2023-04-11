package main

import (
	"io"
	"log"
	"net"
	"strings"
)

var kvStore *Store

func init() {
	kvStore = NewStore()
}

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

		respString := string(buf[:b])

		if respString[0] == '*' {
			res := parseArray(respString)
			command := strings.ToLower(res[0].(string))

			switch command {
			case COMMAND_PING:
				handlePing(conn, res)
			case COMMAND_ECHO:
				handleEcho(conn, res)
			case COMMAND_SET:
				handleSet(conn, res)
			case COMMAND_GET:
				handleGet(conn, res)
			default:
				conn.Write([]byte(encodeError(ErrUnknownCommand + command)))
			}
		}
	}
}
