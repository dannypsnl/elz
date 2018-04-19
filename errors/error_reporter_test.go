package errors

import (
	"testing"
)

func TestEmit(t *testing.T) {
	reporter := NewReporter()
	reporter.Emit("mismatch keyword `let`")

	if len(reporter.ErrMsgs) != 1 {
		t.Error("Didn't emit success")
	}
}
