package logger_custome

import (
	"fmt"
	"time"
)

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
	Fatal LogLevel = "fatal"
)

type Model struct {
	Level LogLevel
	logCh chan string
}

func New(level LogLevel) *Model {
	return &Model{Level: level, logCh: make(chan string)}
}

func (m *Model) TimeNow() time.Time {
	return time.Now()
}

func (m *Model) Info(text string) {
	fmt.Printf("%vINFO%v: %v", Blue, Reset, text)
}

func (m *Model) Warn(text string) {
	fmt.Printf("%vWARN%v: %v", Yellow, Reset, text)
}

func (m *Model) Error(text string) {
	fmt.Printf("%vERROR%v: %v", Red, Reset, text)

}
