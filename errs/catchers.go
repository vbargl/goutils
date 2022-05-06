package errs

import (
	"errors"
	"log"
)

var (
	// ErrUnknownType error tells that recovered value was not error
	ErrUnknownType = errors.New("panic was not an error")
)

// Identity function returns given error.
func Identity(err error) error {
	return err
}

// Logger function will recover error if any and pass it into l.Print() function.
func Logger(l *log.Logger) {
	if e := recover(); e != nil {
		l.Print(e)
	}
}

// StdLogger function will recover error if any and pass it into log.Print() function.
func StdLogger() {
	if e := recover(); e != nil {
		log.Print(e)
	}
}

// Assign function will assign recovered error, if any, into given error address.
// It tries to cast recovered value into error, if it does not succeed it will
// assign ErrUnknownType.
func Assign(err *error) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			*err = e
		} else {
			*err = ErrUnknownType
		}
	}
}

// Transform function will recover error if any occurs.
// Recovered error will be passed into wrapper and result will be assigned into given error pointer.
// It tries to cast recovered value into error, if it does not succeed it will
// assign ErrUnknownType.
//
// DO NOT return nil from wrapper and rather use your own recover mechanism.
func Transform(err *error, wrapper ErrorWrapper) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			*err = wrapper(e)
		} else {
			*err = ErrUnknownType
		}
	}
}

// Run function will recover error if any and pass it into given handler.
// It tries to cast recovered value into error, if it does not succeed it will
// pass a ErrUnknownType.
//
// You SHOULD rather use your own recover mechanism.
func Run(handler ErrorHandler) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			handler(e)
		} else {
			handler(ErrUnknownType)
		}
	}
}
