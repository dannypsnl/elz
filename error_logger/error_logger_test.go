package error_logger

import (
	"testing"
)

func TestEmit(t *testing.T) {
	errLogger := New()
	errLogger.Emit("mismatch keyword `let`")

	if len(errLogger.errMsgs) != 1 {
		t.Error("Didn't emit success")
	}
}
