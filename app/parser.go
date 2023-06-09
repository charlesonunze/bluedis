package main

import (
	"strconv"
	"strings"
)

func parseArray(s string) []interface{} {
	var result []interface{}

	if len(s) < 2 {
		panic("cannot parse")
	}

	lengthOfArray, _ := strconv.Atoi(string(s[1]))
	crlfLength := len(CRLF)
	arrElements := s[4:]

	index := 0

	for i := 0; i < lengthOfArray; i++ {
		currentEl := arrElements[index]

		// Bulk Strings for now
		if currentEl == '$' {
			var sb strings.Builder

			// calculate word length
			for j := index + 1; j < len(arrElements); j++ {
				char := arrElements[j]
				if char == '\r' {
					break
				}
				sb.WriteString(string(char))
			}

			str := sb.String()
			lengthOfWord, _ := strconv.Atoi(str)

			start := index + crlfLength + len(str) + 1
			end := start + lengthOfWord
			word := arrElements[start:end]

			result = append(result, word)

			index = end + crlfLength
		}
	}

	return result
}
