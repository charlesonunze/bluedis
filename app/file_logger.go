package main

import (
	"bufio"
	"fmt"
	"os"
)

type EventType byte

const (
	_                  = iota
	EventDel EventType = iota
	EventPut
)

type Data struct {
	Key       string
	Value     string
	EventType EventType
}

type FileTransactionLogger struct {
	file *os.File
	buf  *bufio.Writer
}

func NewFileTransactionLogger(f *os.File) *FileTransactionLogger {
	if _, err := os.Stat(f.Name()); err != nil {
		panic(err)
	}

	return &FileTransactionLogger{
		file: f,
		buf:  bufio.NewWriter(f),
	}
}

func (l *FileTransactionLogger) WritePut(key, value string) {
	data := Data{
		Key:       key,
		Value:     value,
		EventType: EventPut,
	}

	_, err := fmt.Fprintf(
		l.buf,
		"%d\t%s\t%s\n",
		data.EventType, data.Key, data.Value)

	if err != nil {
		l.buf.Flush() // write any remaining data before panicking
		panic(err)
	}
}

func (l *FileTransactionLogger) WriteDelete(key string) {
	data := Data{
		Key:       key,
		EventType: EventDel,
		Value:     "del", // temporary :)
	}

	_, err := fmt.Fprintf(
		l.buf,
		"%d\t%s\t%s\n",
		data.EventType, data.Key, data.Value)

	if err != nil {
		l.buf.Flush() // write any remaining data before panicking
		panic(err)
	}
}

func (l *FileTransactionLogger) Restore(kvStore *Store) {
	l.buf.Flush()

	f, err := os.Open(l.file.Name())
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var d Data

	for scanner.Scan() {
		line := scanner.Text()

		_, err := fmt.Sscanf(
			line,
			"%d\t%s\t%s",
			&d.EventType, &d.Key, &d.Value)

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		switch d.EventType {
		case EventPut:
			kvStore.Set(d.Key, d.Value)
		case EventDel:
			kvStore.Del(d.Key)
		}
	}
}

func (l *FileTransactionLogger) Close() {
	l.buf.Flush() // write any remaining data before closing
	l.file.Close()
}
