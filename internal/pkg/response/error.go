package response

//BadRequestError incorrect request from user
type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}

func BadRequest(message string) error {
	return &BadRequestError{
		Message: message,
	}
}

func ErrBadRequest(err error) error {
	return &BadRequestError{
		Message: err.Error(),
	}
}

//NotFoundError resource not found in service
type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "not found"
}

func ErrNotFound() error {
	return &NotFoundError{}
}

//InternalServerError unexpected error for service
type InternalServerError struct{}

func (e *InternalServerError) Error() string {
	return "internal server error"
}

func ErrInternal() error {
	return &InternalServerError{}
}
