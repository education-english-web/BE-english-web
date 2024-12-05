package errors

import (
	"fmt"
	"net/http"
)

type serviceUnavailableError struct {
	code    Code
	message string
	params  interface{}
}

// NewServiceUnavailableError inits a system service unavailable error
func NewServiceUnavailableError(
	code Code,
	message string,
	params interface{},
) SystemError {
	return &serviceUnavailableError{
		code:    code,
		message: message,
		params:  params,
	}
}

// Type returns error type
func (e *serviceUnavailableError) Type() TypeError {
	return TypeServiceUnavailable
}

// Code returns error code
func (e *serviceUnavailableError) Code() Code {
	return e.code
}

// Message returns error message
func (e *serviceUnavailableError) Message() string {
	return e.message
}

// Param returns error param value
func (e *serviceUnavailableError) Param() interface{} {
	return e.params
}

// StatusCode returns http status code
func (e *serviceUnavailableError) StatusCode() int {
	return http.StatusServiceUnavailable
}

// Error implements error interface
func (e *serviceUnavailableError) Error() string {
	return fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s",
		e.Type(), e.Code(), e.Message())
}
