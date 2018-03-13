package error_logger

import (
	"fmt"
)

// ErrorLogger organize error emit by it user
type ErrorLogger struct {
	errMsgs []string
}

func New() *ErrorLogger {
	return &ErrorLogger{
		errMsgs: make([]string, 0),
	}
}

func (el *ErrorLogger) Emit(msg string) {
	el.errMsgs = append(el.errMsgs, msg)
}

func (el *ErrorLogger) Report() {
	for i, eMsg := range el.errMsgs {
		fmt.Printf("Error[%d]: %s\n", i, eMsg)
	}
}
