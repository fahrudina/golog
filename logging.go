package golog

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

type LogLevel int

const (
	ERROR LogLevel = iota
	WARN
	INFO
	TRACE
)

var MaxLogLevel LogLevel = TRACE

var line int
var file string

// SetLogLevel sets MaxLogLevel based on the provided string
func SetLogLevel(level string) (ok bool) {

	log.SetFlags(log.Lmicroseconds) //ngasih tau file dan baris yang dieksekusi, gitu
	//log.SetFlags( log.Lmicroseconds | log.Llongfile ) // microsecond resolution: 01:23:23.123123.  assumes Ltime.

	switch strings.ToUpper(level) {
	case "ERROR":
		MaxLogLevel = ERROR
	case "WARN":
		MaxLogLevel = WARN
	case "INFO":
		MaxLogLevel = INFO
	case "TRACE":
		MaxLogLevel = TRACE
	default:
		LogError("Unknown log level requested: %v", level)
		return false
	}
	return true
}

// Error logs a message to the 'standard' Logger (always)
func LogError(msg string, args ...interface{}) {
	msg = "[ERROR] " + msg
	log.Printf(msg, args...)

	// Thanks TO https://golang.org/src/log/log.go

	_, file, line, _ = runtime.Caller(1)
	result := strings.Split(file, "/")
	fmt.Println(result[len(result)-2]+"/"+result[len(result)-1], line, "\n")
}

// Warn logs a message to the 'standard' Logger if MaxLogLevel is >= WARN
func LogWarn(msg string, args ...interface{}) {
	if MaxLogLevel >= WARN {
		msg = "[WARN] " + msg
		log.Printf(msg, args...)

		_, file, line, _ = runtime.Caller(1)
		result := strings.Split(file, "/")
		fmt.Println(result[len(result)-2]+"/"+result[len(result)-1], line, "\n")
	}
}

// Info logs a message to the 'standard' Logger if MaxLogLevel is >= INFO
func LogInfo(msg string, args ...interface{}) {
	if MaxLogLevel >= INFO {

		msg = "[INFO] " + msg
		log.Printf(msg, args...)

		_, file, line, _ = runtime.Caller(1)
		result := strings.Split(file, "/")
		fmt.Println(result[len(result)-2]+"/"+result[len(result)-1], line, "\n")
	}
}

// Trace logs a message to the 'standard' Logger if MaxLogLevel is >= TRACE
func LogTrace(msg string, args ...interface{}) {
	if MaxLogLevel >= TRACE {
		msg = "[TRACE] " + msg
		log.Printf(msg, args...)

		_, file, line, _ = runtime.Caller(1)
		result := strings.Split(file, "/")
		fmt.Println(result[len(result)-2]+"/"+result[len(result)-1], line, "\n")
	}
}
