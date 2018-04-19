package errors

import (
	"fmt"
)

// Reporter organize error emit by it user
type Reporter struct {
	ErrMsgs []string
}

func NewReporter() *Reporter {
	return &Reporter{
		ErrMsgs: make([]string, 0),
	}
}

func (el *Reporter) Emit(msg string) {
	el.ErrMsgs = append(el.ErrMsgs, msg)
}

func (el *Reporter) Report() {
	for i, eMsg := range el.ErrMsgs {
		fmt.Printf("Error[%d]: %s\n", i, eMsg)
	}
}

func (r *Reporter) HasNoError() bool {
	return len(r.ErrMsgs) == 0
}
