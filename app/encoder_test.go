package main

import (
	"testing"
)

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

func Test_encodeSimpleString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"", "", "+\r\n"},
		{"", "hello", "+hello\r\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeSimpleString(tt.s); got != tt.want {
				t.Errorf("encodeSimpleString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeError(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"", "", "-\r\n"},
		{"", "error", "-error\r\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeError(tt.s); got != tt.want {
				t.Errorf("encodeError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeStrings(t *testing.T) {
	tests := []struct {
		name string
		s    []any
		want string
	}{
		{"", []any{"$", "h", "i"}, "hi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeStrings(tt.s); got != tt.want {
				t.Errorf("mergeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
