package errs

type (
	// ErrorWrapper is type alias used in Wrap(...) function
	// See Wrap(...) function
	ErrorWrapper = func(err error) error

	// ErrorHandler is type alias used in Run(...) function
	// See Run(...) function
	ErrorHandler = func(err error)
)
