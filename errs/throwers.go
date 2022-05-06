package errs

import "log"

// Check function checks if given error is nil,
// otherwise panic
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// Wrap function checks if given error is nil.
// If error is non-nil it tries to wrap it via WrapperFunc
// and panic. If WrapperFunc returns nil,
// error is considered as handled and does not panic
func Wrap(err error, wrapper ErrorWrapper) {
	if err != nil {
		if w := wrapper(err); w != nil {
			panic(w)
		}
	}
}

// Mask function checks if given error is nil.
// if error is non-nil it panic with other error.
func Mask(err error, other error) {
	if err != nil {
		panic(other)
	}
}

// Log function check if given error is nil,
// otherwise it log it with l.Print(...) function.
func Log(err error, l *log.Logger) {
	if err != nil {
		l.Print(err)
	}
}

// Log function check if given error is nil,
// otherwise it log it with log.Print(...) function.
func StdLog(err error) {
	if err != nil {
		log.Print(err)
	}
}

// Fatal function check if given error is nil,
// otherwise it log it with l.Fatal(...) function.
func Fatal(err error, l *log.Logger) {
	if err != nil {
		log.Fatal(err)
	}
}

// StdFatal function check if given error is nil,
// otherwise it log it with log.Fatal(...) function.
func StdFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
