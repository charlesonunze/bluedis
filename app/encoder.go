package main

import "fmt"

const CRLF = "\r\n"

func encodeBulkString(s string) string {
	return fmt.Sprintf("$%d%s%s%s", len(s), CRLF, s, CRLF)
}
