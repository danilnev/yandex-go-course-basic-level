package main

import "fmt"

type Logger interface {
	Log(message string)
}

type LogLevel string

const (
	Error LogLevel = "Error"
	Info  LogLevel = "Info"
)

type Log struct {
	Level LogLevel
}

func (l Log) Log(message string) {
	switch l.Level {
	case Error:
		fmt.Println("ERROR:", message)
	case Info:
		fmt.Println("INFO:", message)
	}
}

func main() {
	var logger Logger = Log{Level: Error}
	logger.Log("This is an error message.")

	logger = Log{Level: Info}
	logger.Log("This is an info message.")
}
