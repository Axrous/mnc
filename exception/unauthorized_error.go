package exception

type UnauthorizeError struct {
	Error string
}

func NewUnauthorizeError(err string) UnauthorizeError {
	return UnauthorizeError{
		Error: err,
	}
}