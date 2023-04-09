package main

import (
	"fmt"
	"strings"
)

const CRLF = "\r\n"

func encodeBulkString(s string) string {
	return fmt.Sprintf("$%d%s%s%s", len(s), CRLF, s, CRLF)
}

func mergeStrings(s []interface{}) string {
	var sb strings.Builder
	for i := 1; i < len(s); i++ {
		sb.WriteString(s[i].(string))
	}
	return sb.String()
}
