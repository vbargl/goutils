package errs

import (
	"errors"
	"fmt"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

// Todo function will panic with ErrNotImplemented error.
func Todo() {
	panic(ErrNotImplemented)
}

// TodoWithMessage function will panic with 'not implemented: ' concatenate with custom message
func TodoWithMessage(message string) {
	panic(fmt.Errorf("not implemented: %s", message))
}
