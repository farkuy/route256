package flogger

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type LogStruct struct {
	text  string
	level LogLevel
	time  string
}

type Model struct {
	logCh chan LogStruct
}

func new() *Model {

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	model := &Model{logCh: make(chan LogStruct, 100)}
	defer close(model.logCh)

	go func() {

		for {
			select {
			case logData, ok := <-model.logCh:
				if !ok {
					fmt.Println("закртыие в logData case")
					return
				}

				color := getColorLog(logData.level)
				fmt.Printf("%v %v%v%v: %v \n", logData.time, color, logData.level, Reset, logData.text)
			case <-stopChan:
				fmt.Println("закртыие в stopChan case")
				return
			}
		}
	}()

	return model
}

var logger *Model = new()

func timeNow() string {
	return time.Now().Format("02-01-2006 15:04:05")
}

func Info(text string) {
	logger.logCh <- LogStruct{text, InfoLvl, timeNow()}
}

func Warn(text string) {
	logger.logCh <- LogStruct{text, WarnLvl, timeNow()}
}

func Error(text string) {
	logger.logCh <- LogStruct{text, ErrorLvl, timeNow()}
}

func getColorLog(lvl LogLevel) LoggerColor {
	switch lvl {
	case InfoLvl:
		return Blue
	case WarnLvl:
		return Yellow
	case ErrorLvl:
		return Red
	default:
		return Gray
	}
}
