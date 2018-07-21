package listener

import (
	"testing"
)

func TestReturnNotInFunction(t *testing.T) {
	source := `
	return 10
	`

	listener := listener(source)

	if listener.context.Reporter.ErrMsgs[0] != "return statement must in function" {
		t.Error(listener.context.Reporter.ErrMsgs)
	}
}
