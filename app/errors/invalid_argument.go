package errors

import (
	"fmt"
	"net/http"
)

//nolint:errname
type errorInvalidArgument struct {
	code    Code
	message string
	param   interface{}
}

// NewErrorInvalidArgument inits a system invalid argument error
func NewErrorInvalidArgument(code Code, message string, param interface{}) SystemError {
	return &errorInvalidArgument{code, message, param}
}

// Type returns error type
func (e *errorInvalidArgument) Type() TypeError {
	return TypeInvalidArgument
}

// Code returns error code
func (e *errorInvalidArgument) Code() Code {
	return e.code
}

// Message returns error message
func (e *errorInvalidArgument) Message() string {
	return e.message
}

// Param returns error param value
func (e *errorInvalidArgument) Param() interface{} {
	return e.param
}

// StatusCode returns http status code
func (e *errorInvalidArgument) StatusCode() int {
	return http.StatusBadRequest
}

// Error implements error interface
func (e *errorInvalidArgument) Error() string {
	return fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s, \tParam: %v",
		TypeInvalidArgument, e.Code(), e.Message(), e.Param())
}
