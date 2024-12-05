package errors

import (
	"fmt"
	"net/http"
)

//nolint:errname
type errorUnauthorized struct {
	code    Code
	message string
	param   interface{}
}

// NewErrorUnauthorized inits an unauthorized error
func NewErrorUnauthorized() SystemError {
	return &errorUnauthorized{
		code:    CodeUnauthorized,
		message: "unauthorized",
		param:   nil,
	}
}

// Type returns error type
func (e *errorUnauthorized) Type() TypeError {
	return TypeUnauthorized
}

// Code returns error code
func (e *errorUnauthorized) Code() Code {
	return e.code
}

// Message returns error message
func (e *errorUnauthorized) Message() string {
	return e.message
}

// Param returns error param value
func (e *errorUnauthorized) Param() interface{} {
	return e.param
}

// StatusCode returns http status code
func (e *errorUnauthorized) StatusCode() int {
	return http.StatusUnauthorized
}

// Error implements error interface
func (e *errorUnauthorized) Error() string {
	return fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s",
		TypeUnauthorized, e.Code(), e.Message())
}

// NewCustomErrorUnauthorized alternatives function to return UNAUTHORIZED error with custom values
func NewCustomErrorUnauthorized(code Code, message string, param interface{}) SystemError {
	return &errorUnauthorized{
		code:    code,
		message: message,
		param:   param,
	}
}
