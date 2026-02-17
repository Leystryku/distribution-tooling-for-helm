// Package pterm provides a logger implementation using the pterm library
package pterm

import (
	"fmt"
	"io"
	"os"

	"github.com/pterm/pterm"
	"github.com/vmware-labs/distribution-tooling-for-helm/pkg/dtlog"
)

// NewLogger returns a new Logger implemented by pterm
func NewLogger() *Logger {
	return &Logger{writer: os.Stdout, level: dtlog.InfoLevel}
}

// Logger defines a logger implemented using pterm
type Logger struct {
	writer io.Writer
	level  dtlog.Level
	prefix string
}

func (l *Logger) printMessage(messageLevel dtlog.Level, printer *pterm.PrefixPrinter, format string, args ...interface{}) {
	if messageLevel > l.level {
		return
	}
	pterm.Fprintln(l.writer, l.prefix+printer.Sprint(fmt.Sprintf(format, args...)))
}

// SetWriter sets the internal writer used by the log
func (l *Logger) SetWriter(w io.Writer) {
	l.writer = w
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level dtlog.Level) {
	l.level = level
}

// Failf logs a formatted error and returns it back
func (l *Logger) Failf(format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	l.Errorf("%v", err)
	return &dtlog.LoggedError{Err: err}
}

// Printf prints a message in the log
func (l *Logger) Printf(format string, args ...interface{}) {
	l.printMessage(dtlog.AlwaysLevel, Plain, format, args...)
}

// Errorf logs an error message
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.printMessage(dtlog.ErrorLevel, Error, format, args...)
}

// Infof logs an information message
func (l *Logger) Infof(format string, args ...interface{}) {
	l.printMessage(dtlog.InfoLevel, Info, format, args...)
}

// Debugf logs a debug message
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.printMessage(dtlog.DebugLevel, Debug, format, args...)
}

// Warnf logs a warning message
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.printMessage(dtlog.WarnLevel, Warning, format, args...)
}
