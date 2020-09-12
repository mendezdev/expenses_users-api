package api_errors

import (
	"errors"
	"fmt"
	"net/http"
)

// RestErr is the interface to get access of restErr
type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

// Error get the structure of the error
func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

// Message get the message description of the error
func (e restErr) Message() string {
	return e.ErrMessage
}

// Status is the StatusCode for the api error
func (e restErr) Status() int {
	return e.ErrStatus
}

// Causes is the list of causes
func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

//NewError creates a generic error
func NewError(message string) error {
	return errors.New(message)
}

//NewRestError creates a RestError given this parameters
func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	if causes == nil {
		causes = make([]interface{}, 0)
	}
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

//NewBadRequestError creates a BadRequest - 400 RestErr interface
func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

//NewNotFoundError creates a NotFound - 404 RestErr interface
func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

//NewUnauthorizedError creates a UnauthorizeError - 401 RestErr interface
func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

//NewInternalServerError creates a InternalServerError - 500 RestErr interface
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}
