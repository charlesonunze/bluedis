package main

import (
	"net"
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
