package utils

import (
	"fmt"
	"log"
	"os"
)

// Logger ...
type Logger struct {
	k *log.Logger
}

// NewLogger creates a new LOGGER
func NewLogger() *Logger {
	return &Logger{log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)}
}

// Info ...
func (l *Logger) Info(v ...any) {
	_ = l.k.Output(2, "INFO: "+fmt.Sprintln(v...))
}

// Warn ...
func (l *Logger) Warn(v ...any) {
	_ = l.k.Output(2, "WARN: "+fmt.Sprintln(v...))
}

// Error ...
func (l *Logger) Error(v ...any) {
	_ = l.k.Output(2, "ERROR: "+fmt.Sprintln(v...))
}

// Debug ...
func (l *Logger) Debug(v ...any) {
	_ = l.k.Output(2, "DEBUG: "+fmt.Sprintln(v...))
}
