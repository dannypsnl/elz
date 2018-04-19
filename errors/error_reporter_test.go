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

func TestHasNoError(t *testing.T) {
	reporter := NewReporter()
	if !reporter.HasNoError() {
		t.Error(`Initial state of Reporter is wrong, should not contain any error`)
	}
	reporter.Emit("some error")
	if reporter.HasNoError() {
		t.Error("After emit some error, HasNoError should return false but true")
	}
}
