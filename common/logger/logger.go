// Package logger handles the common task of logging
package logger

import (
	"log"
	"runtime"
)

func getFrame() runtime.Frame {
	programCounters := make([]uintptr, 15)
	programCountersCount := runtime.Callers(3, programCounters)
	frames := runtime.CallersFrames(programCounters[:programCountersCount])
	frame, _ := frames.Next()

	return frame
}

func Info(message string) {
	log.Printf("[%s] - %q\n", getFrame().Function, message)
}

func Error(message string, err error) {
	log.Printf("[%s] - %q\n", getFrame().Function, message)
	log.Println(err)
}

func Fatal(message string, err error) {
	log.Printf("[%s] - %q\n", getFrame().Function, message)
	log.Fatalln(err)
}
