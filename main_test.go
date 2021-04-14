package drums

import (
	"fmt"
	"testing"
)

type MustTest struct {
	action      string
	err         error
	shouldPanic bool
}

func TestMust(t *testing.T) {

	mustTests := []MustTest{
		{
			action:      "test action",
			err:         nil,
			shouldPanic: false,
		},
		{
			action:      "test action",
			err:         fmt.Errorf("mock error"),
			shouldPanic: true,
		},
	}

	for _, column := range mustTests {
		defer func(shouldPanic bool) {
			r := recover()
			if r == nil && shouldPanic {
				t.Errorf("must did not panic when given an error")
			}
			if r != nil && !shouldPanic {
				t.Errorf("must panicked when given nil")
			}
		}(column.shouldPanic)
		must(column.action, column.err)
	}
}
