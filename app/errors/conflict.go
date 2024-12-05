package errors

import (
	"fmt"
	"net/http"
)

//nolint:errname
type errorConflict struct {
	code    Code ``
	message string
}

// NewErrorConflict inits a system invalid argument error
func NewErrorConflict(code Code, message string) SystemError {
	return &errorConflict{code, message}
}

// Type returns error type
func (e *errorConflict) Type() TypeError {
	return TypeConflict
}

// Code returns error code
func (e *errorConflict) Code() Code {
	return e.code
}

// Message returns error message
func (e *errorConflict) Message() string {
	return e.message
}

// Param returns error param value
func (e *errorConflict) Param() interface{} {
	return nil
}

// StatusCode returns http status code
func (e *errorConflict) StatusCode() int {
	return http.StatusConflict
}

// Error implements error interface
func (e *errorConflict) Error() string {
	return fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s",
		TypeConflict, e.Code(), e.Message())
}
