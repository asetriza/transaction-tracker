package domain

type ValidationError struct {
	s string
}

func NewValidationError(text string) error {
	return ValidationError{
		s: text,
	}
}

func (e ValidationError) Error() string {
	return e.s
}
