package exceptions

type ValidationError struct {
	Message string
	Details map[string]string
}

func (e ValidationError) Error() string {
	return e.Message
}

func (e ValidationError) Add(key, message string) {
	e.Details[key] = message
}
