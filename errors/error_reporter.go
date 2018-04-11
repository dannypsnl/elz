package errors

import (
	"fmt"
)

// Reporter organize error emit by it user
type Reporter struct {
	errMsgs []string
}

func NewReporter() *Reporter {
	return &Reporter{
		errMsgs: make([]string, 0),
	}
}

func (el *Reporter) Emit(msg string) {
	el.errMsgs = append(el.errMsgs, msg)
}

func (el *Reporter) Report() {
	for i, eMsg := range el.errMsgs {
		fmt.Printf("Error[%d]: %s\n", i, eMsg)
	}
}
