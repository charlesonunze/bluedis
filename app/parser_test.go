package main

import (
	"reflect"
	"testing"
)

func Test_parseArray(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []interface{}
	}{
		{"", "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n", []interface{}{"hello", "world"}},
		{"", "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n", []interface{}{"ECHO", "hey"}},
		{"", "*2\r\n$4\r\nECHO\r\n$11\r\nhello world\r\n", []interface{}{"ECHO", "hello world"}},
		{"", "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n", []interface{}{"SET", "key", "value"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseArray(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
