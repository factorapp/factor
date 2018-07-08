// Package log implements a simple logging package.
// It indicates log level, datetime and file location before printing a message.
package log

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

var (
	info  = log.New(os.Stdout, "\033[36mINFO\033[00m  ", log.LstdFlags)
	warn  = log.New(os.Stdout, "\033[33mWARN\033[00m  ", log.LstdFlags)
	error = log.New(os.Stderr, "\033[31mERROR\033[00m ", log.LstdFlags)
)

// Info calls Output to print to the info logger.
// Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {
	v = append([]interface{}{caller()}, v...)
	info.Print(v...)
}

// Infof calls Output to print to the info logger.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	v = append([]interface{}{caller()}, v...)
	format = addCallerFormat(format)
	info.Printf(format, v...)
}

// Warn calls Output to print to the warn logger.
// Arguments are handled in the manner of fmt.Print.
func Warn(v ...interface{}) {
	v = append([]interface{}{caller()}, v...)
	warn.Print(v...)
}

// Warnf calls Output to print to the warn logger.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	v = append([]interface{}{caller()}, v...)
	format = addCallerFormat(format)
	warn.Printf(format, v...)
}

// Error calls Output to print to the error logger.
// Arguments are handled in the manner of fmt.Print.
func Error(v ...interface{}) {
	v = append([]interface{}{caller()}, v...)
	error.Print(v...)
}

// Errorf calls Output to print to the error logger.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	v = append([]interface{}{caller()}, v...)
	format = addCallerFormat(format)
	error.Printf(format, v...)
}

// Panic is equivalent to Error() followed by a call to panic().
func Panic(v ...interface{}) {
	v = append([]interface{}{caller()}, v...)
	error.Print(v...)
	panic(fmt.Sprint(v...))
}

// Panicf is equivalent to Errorf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	v = append([]interface{}{caller()}, v...)
	format = addCallerFormat(format)
	error.Printf(format, v...)
	panic(fmt.Sprintf(format, v...))
}

func caller() string {
	_, f, l, _ := runtime.Caller(2)
	f = path.Base(f)

	return fmt.Sprintf("%v:%v: ", f, l)
}

func addCallerFormat(format string) string {
	return "%v" + format
}
