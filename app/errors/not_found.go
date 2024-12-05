package errors

import (
	"fmt"
	"net/http"
)

//nolint:errname
type errorNotFound struct {
	code    Code
	message string
	param   interface{}
}

// NewErrorNotFound inits a system not found error
func NewErrorNotFound(code Code, message string, param interface{}) SystemError {
	return &errorNotFound{code, message, param}
}

// Type returns error type
func (e *errorNotFound) Type() TypeError {
	return TypeNotFound
}

// Code returns error code
func (e *errorNotFound) Code() Code {
	return e.code
}

// Message returns error message
func (e *errorNotFound) Message() string {
	return e.message
}

// Param returns error param value
func (e *errorNotFound) Param() interface{} {
	return e.param
}

// StatusCode returns http status code
func (e *errorNotFound) StatusCode() int {
	return http.StatusNotFound
}

// Error implements error interface
func (e *errorNotFound) Error() string {
	return fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s, \tParam: %v",
		TypeNotFound, e.Code(), e.Message(), e.Param())
}
