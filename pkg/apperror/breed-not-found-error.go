package apperror

type BreedNotFoundError struct {
	message string
}

func NewBreedNotFoundError(msg string) *BreedNotFoundError {
	return &BreedNotFoundError{
		message: msg,
	}
}

func (e BreedNotFoundError) Error() string {
	return e.message
}
