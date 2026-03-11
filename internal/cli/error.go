package cli

// ExitCoder is implemented by errors that want to control the process exit code.
type ExitCoder interface {
	error
	ExitCode() int
}

type exitError struct {
	err  error
	code int
}

func (e *exitError) Error() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

func (e *exitError) Unwrap() error {
	return e.err
}

func (e *exitError) ExitCode() int {
	if e.code == 0 {
		return 1
	}
	return e.code
}

func newExitError(code int, err error) error {
	return &exitError{err: err, code: code}
}
