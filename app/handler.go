package main

import (
	"net"
	"strconv"
	"strings"
	"time"
)

func handlePing(conn net.Conn, res []any) {
	var result string

	if len(res) > 1 {
		result = encodeBulkString(mergeStrings(res))
	} else {
		result = encodeSimpleString("PONG")
	}

	conn.Write([]byte(result))
}

func handleEcho(conn net.Conn, res []any) {
	result := encodeBulkString(mergeStrings(res))
	conn.Write([]byte(result))
}

func handleSet(conn net.Conn, res []any) {
	k, v := res[1].(string), res[2].(string)
	kvStore.Set(k, v)

	handleSetOptions(k, v, res[3:])

	result := encodeSimpleString("OK")
	conn.Write([]byte(result))
}

func handleGet(conn net.Conn, res []any) {
	k := res[1].(string)
	v := kvStore.Get(k)

	if len(v) == 0 {
		conn.Write([]byte(NULL))
		return
	}

	result := encodeBulkString(v)
	conn.Write([]byte(result))
}

func handleSetOptions(key, val string, options []any) {
	for i, option := range options {
		opt := strings.ToLower(option.(string))
		switch opt {
		case "ex":
			expiryTime := options[i+1].(string)
			ex, _ := strconv.Atoi(expiryTime)

			go func() {
				select {
				case <-time.After(time.Duration(ex) * time.Second):
					kvStore.Del(key)
					return
				}
			}()

		case "px":
			expiryTime := options[i+1].(string)
			px, _ := strconv.Atoi(expiryTime)

			go func() {
				select {
				case <-time.After(time.Duration(px) * time.Millisecond):
					kvStore.Del(key)
					return
				}
			}()

		case "exat":
		case "pxat":
		case "nx":
		case "xx":
		case "keepttl":
		case "get":
		}
	}
}
