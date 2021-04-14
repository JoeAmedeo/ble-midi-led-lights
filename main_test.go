package drums

import (
	"fmt"
	"testing"
)

func TestMustPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("must did not panic when given an error")
		}
	}()
	err := fmt.Errorf("mock error")
	must("test action", err)
}

func TestMustNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("must panicked when no error was provided")
		}
	}()
	must("test action", nil)
}
