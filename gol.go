// Package gol provides simple logging levels with tracing and formating.
// Simply call gol.LEVEL(fmt format) from your code. Supported levels are
// exported as functions. With Set() you can set the level to one of FATAL,
// ERROR, WARN, INFO, DEBUG, TRACE. Messages of equal or higher level will
// be printed.
//
// Sample output:
//
// INFO: Message (/Users/user/go/src/prog/main.go,:52 main.main)
//
package gol

import (
	"fmt"
	"runtime"
	"strings"
)

type GOL int32

type gol struct {
	n     GOL
	str   string
	trace GOL
}

// Supported log levels and trace status
const (
	FATAL GOL = iota
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
	ON
	OFF
)

var errorStrings = [...]string{
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
	"ON",
	"OFF",
}

func (c GOL) String() string {
	if (c >= 0) && (c < GOL(len(errorStrings))) {
		return errorStrings[c]
	}
	return "UNKNOWN"
}

var lv *gol

func init() {
	lv = new(gol)
	lv.n = INFO
	lv.str = INFO.String()
	lv.trace = ON
}

func SetTrace(trace GOL) bool {
	if trace >= ON && trace <= OFF {
		lv.trace = trace
		return true
	}
	return false
}

func SetLevel(lev GOL) bool {
	if lev >= FATAL && lev <= TRACE {
		lv.n = lev
		lv.str = errorStrings[lev]
		return true
	}
	return false
}

func SetLevelString(lev string) bool {
	for g, s := range errorStrings {
		if s == lev {
			return SetLevel(GOL(g))
		}
	}
	return false
}

func Fatal(v interface{}, a ...interface{}) (n int, err error) {
	if lv.n >= FATAL {
		return printf(FATAL.String(), v, a...)
	}
	return 0, nil
}

func Error(v interface{}, a ...interface{}) (n int, err error) {
	if lv.n >= ERROR {
		return printf(ERROR.String(), v, a...)
	}
	return 0, nil
}

func Warn(v interface{}, a ...interface{}) (n int, err error) {
	if lv.n >= WARN {
		return printf(WARN.String(), v, a...)
	}
	return 0, nil
}

func Info(v interface{}, a ...interface{}) (n int, err error) {
	if lv.n >= INFO {
		return printf(INFO.String(), v, a...)
	}
	return 0, nil
}

func Debug(v interface{}, a ...interface{}) (n int, err error) {
	if lv.n >= DEBUG {
		return printf(DEBUG.String(), v, a...)
	}
	return 0, nil
}

func Trace(v interface{}, a ...interface{}) (n int, err error) {
	if lv.n >= TRACE {
		return printf(TRACE.String(), v, a...)
	}
	return 0, nil
}

func printf(s string, v interface{}, a ...interface{}) (n int, err error) {
	var f string

	switch v.(type) {
	case string:
		f = v.(string)
	case error:
		f = v.(error).Error()
	default:
		f = ""
	}

	f = strings.TrimSuffix(f, "\n")
	if lv.trace == ON {
		f = s + ": " + f + ftrace()
	} else {
		f = s + ": " + f + "\n"
	}

	return fmt.Printf(f, a...)
}

// ftrace will print function/file/line INFOrmation
// runtime.Callers skips itself, the caller and printf & Error func below (thus 4)
func ftrace() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(4, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return fmt.Sprintf(" (%s,:%d %s)\n", frame.File, frame.Line, frame.Function)
}
