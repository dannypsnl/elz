package error_logger

import (
	"testing"
)

func TestEmit(t *testing.T) {
	reporter := New()
	reporter.Emit("mismatch keyword `let`")

	if len(reporter.errMsgs) != 1 {
		t.Error("Didn't emit success")
	}
}
