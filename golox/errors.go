package golox

type RuntimeError struct {
	message string
	token   Token
}

func (e RuntimeError) Error() string {
	return e.message
}

func NewRuntimeError(token Token, message string) error {
	return &RuntimeError{message, token}
}
