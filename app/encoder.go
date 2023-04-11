package main

import (
	"fmt"
	"strings"
)

const (
	CRLF = "\r\n"
	NULL = "$-1\r\n"
)

func encodeSimpleString(s string) string {
	return fmt.Sprintf("+%s%s", s, CRLF)
}

func encodeBulkString(s string) string {
	return fmt.Sprintf("$%d%s%s%s", len(s), CRLF, s, CRLF)
}

func encodeError(s string) string {
	return fmt.Sprintf("-%s%s", s, CRLF)
}

func mergeStrings(s []any) string {
	var sb strings.Builder
	// skip the first element in the slice
	for i := 1; i < len(s); i++ {
		sb.WriteString(s[i].(string))
	}
	return sb.String()
}
