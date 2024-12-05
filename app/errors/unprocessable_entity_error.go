package errors

import (
	"fmt"
	"net/http"
)

type unprocessableEntityError struct {
	code    Code
	message string
	param   interface{}
}

// NewUnprocessableEntityError inits a system unprocessable entity error
func NewUnprocessableEntityError() SystemError {
	return &unprocessableEntityError{
		code:    CodeUnprocessableEntity,
		message: "unprocessable entity",
		param:   nil,
	}
}

// Type returns error type
func (e *unprocessableEntityError) Type() TypeError {
	return TypeUnprocessableEntity
}

// Code returns error code
func (e *unprocessableEntityError) Code() Code {
	return e.code
}

// Message returns error message
func (e *unprocessableEntityError) Message() string {
	return e.message
}

// Param returns error param value
func (e *unprocessableEntityError) Param() interface{} {
	return e.param
}

// StatusCode returns http status code
func (e *unprocessableEntityError) StatusCode() int {
	return http.StatusUnprocessableEntity
}

// Error implements error interface
func (e *unprocessableEntityError) Error() string {
	return fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s",
		e.Type(), e.Code(), e.Message())
}
