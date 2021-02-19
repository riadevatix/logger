package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// ILogger is an inteface to actual logger
type ILogger interface {
	Logln(vals ...interface{})
	Logf(fomatStr string, vals ...interface{})
}

// logger is an object that writes to an io.Writer
type logger struct {
	mu  sync.Mutex
	out io.Writer
}

// Logln logs the values passed into it and add a new line at the end
func (l *logger) Logln(vals ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	buf := []byte{}
	buf = append(buf, fmt.Sprint(vals...)...)
	buf = append(buf, '\n') // Add a new line at the end of the output
	l.out.Write(buf)
}

// Logf works like fmt.Printf
func (l *logger) Logf(s string, vals ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	buf := []byte{}
	buf = append(buf, fmt.Sprintf(s, vals...)...)
	l.out.Write(buf)
}

// NewLogger creates and returns a Logger wrapped into ILogger interface
func NewLogger(out io.Writer) ILogger {
	return &logger{out: out}
}

var std = NewLogger(os.Stdout)

// DefaultLogger returns the default loger
func DefaultLogger() ILogger {
	return std
}
