package errors

import (
	"fmt"
	"net/http"
)

//nolint:errname
type errorForbidden struct {
	code    Code
	message string
	param   interface{}
}

// NewErrorForbidden inits an forbidden error
func NewErrorForbidden() SystemError {
	return &errorForbidden{
		code:    CodeForbidden,
		message: "forbidden",
		param:   nil,
	}
}

// Type returns error type
func (e *errorForbidden) Type() TypeError {
	return TypeForbidden
}

// Code returns error code
func (e *errorForbidden) Code() Code {
	return e.code
}

// Message returns error message
func (e *errorForbidden) Message() string {
	return e.message
}

// Param returns error param value
func (e *errorForbidden) Param() interface{} {
	return e.param
}

// StatusCode returns http status code
func (e *errorForbidden) StatusCode() int {
	return http.StatusForbidden
}

// Error implements error interface
func (e *errorForbidden) Error() string {
	return fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s, \tParam: %v",
		TypeForbidden, e.Code(), e.Message(), e.Param())
}

// NewCustomErrorForbidden alternative function to return FORBIDDEN error with custom values
func NewCustomErrorForbidden(code Code, message string, param interface{}) SystemError {
	return &errorForbidden{
		code:    code,
		message: message,
		param:   param,
	}
}
