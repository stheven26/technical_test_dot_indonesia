package utils

type ApplicationError struct {
	ErrorCode string
	Message   string
}

func New(errorCode string, message string) error {
	return &ApplicationError{
		ErrorCode: errorCode,
		Message:   message,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func (e *ApplicationError) SetErrorCode(errorCode string) *ApplicationError {
	e.ErrorCode = errorCode
	return e
}

func (e *ApplicationError) SetMessage(message string) *ApplicationError {
	e.Message = message
	return e
}
