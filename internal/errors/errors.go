package errors

type Error struct {
	message string
	code    Code
}

func New(code Code) *Error {
	return &Error{
		code: code,
	}
}

func (e *Error) WithMessage(message string) *Error {
	e.message = message
	return e
}

func (e *Error) Status() string {
	return e.code.Status()
}

func (e *Error) Code() int {
	return int(e.code)
}

func (e *Error) Error() string {
	if e.message != "" {
		return e.message
	}
	return e.Status()
}

func (e *Error) ShouldRetry() bool {
	return e.code.ShouldRetry()
}
