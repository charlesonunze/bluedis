package main

import "testing"

func Test_encodeBulkString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"", "", "$0\r\n\r\n"},
		{"", "hello", "$5\r\nhello\r\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBulkString(tt.s); got != tt.want {
				t.Errorf("encodeBulkString() = %v, want %v", got, tt.want)
			}
		})
	}
}
