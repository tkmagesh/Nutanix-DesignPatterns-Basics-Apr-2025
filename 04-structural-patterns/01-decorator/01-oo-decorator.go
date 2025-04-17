package main

import (
	"fmt"
	"time"
)

// Component
type Logger interface {
	Log(message string)
}

// Concrete Component
type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

// Base Decorator
type LoggerDecorator struct {
	logger Logger
}

func (d LoggerDecorator) Log(message string) {
	d.logger.Log(message)
}

// Timestamp Decorator
type TimestampLogger struct {
	LoggerDecorator
}

func (t TimestampLogger) Log(message string) {
	msg := fmt.Sprintf("%s - %s", time.Now().Format(time.RFC3339), message)
	t.logger.Log(msg)
}

// Level Decorator
type LevelLogger struct {
	LoggerDecorator
	level string
}

func (l LevelLogger) Log(message string) {
	msg := fmt.Sprintf("[%s] %s", l.level, message)
	l.logger.Log(msg)
}

func main() {
	var logger Logger = ConsoleLogger{}
	logger = TimestampLogger{LoggerDecorator{logger}}
	logger = LevelLogger{LoggerDecorator{logger}, "INFO"}

	logger.Log("App started")
}
