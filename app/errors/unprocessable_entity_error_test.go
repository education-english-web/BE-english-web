package errors

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewUnprocessableEntityError(t *testing.T) {
	tests := []struct {
		name string
		want SystemError
	}{
		{
			name: "success",
			want: &unprocessableEntityError{
				code:    CodeUnprocessableEntity,
				message: "unprocessable entity",
				param:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUnprocessableEntityError(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnprocessableEntityError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unprocessableEntityError_Type(t *testing.T) {
	tests := []struct {
		name string
		e    *unprocessableEntityError
		want TypeError
	}{
		{
			name: "success",
			e:    &unprocessableEntityError{},
			want: TypeUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &unprocessableEntityError{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.Type(); got != tt.want {
				t.Errorf("unprocessableEntityError.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unprocessableEntityError_Code(t *testing.T) {
	tests := []struct {
		name string
		e    *unprocessableEntityError
		want Code
	}{
		{
			name: "success",
			e: &unprocessableEntityError{
				code:    CodeUnprocessableEntity,
				message: "unprocessable entity",
				param:   nil,
			},
			want: CodeUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &unprocessableEntityError{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.Code(); got != tt.want {
				t.Errorf("unprocessableEntityError.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unprocessableEntityError_Message(t *testing.T) {
	tests := []struct {
		name string
		e    *unprocessableEntityError
		want string
	}{
		{
			name: "success",
			e: &unprocessableEntityError{
				code:    CodeUnprocessableEntity,
				message: "unprocessable entity",
				param:   nil,
			},
			want: "unprocessable entity",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &unprocessableEntityError{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.Message(); got != tt.want {
				t.Errorf("unprocessableEntityError.Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unprocessableEntityError_Param(t *testing.T) {
	tests := []struct {
		name string
		e    *unprocessableEntityError
		want interface{}
	}{
		{
			name: "success",
			e: &unprocessableEntityError{
				code:    CodeUnprocessableEntity,
				message: "unprocessable entity",
				param:   nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &unprocessableEntityError{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.Param(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unprocessableEntityError.Param() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unprocessableEntityError_StatusCode(t *testing.T) {
	tests := []struct {
		name string
		e    *unprocessableEntityError
		want int
	}{
		{
			name: "success",
			e: &unprocessableEntityError{
				code:    CodeUnprocessableEntity,
				message: "unprocessable entity",
				param:   nil,
			},
			want: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &unprocessableEntityError{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.StatusCode(); got != tt.want {
				t.Errorf("unprocessableEntityError.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unprocessableEntityError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *unprocessableEntityError
		want string
	}{
		{
			name: "success",
			e: &unprocessableEntityError{
				code:    CodeUnprocessableEntity,
				message: "unprocessable entity",
				param:   nil,
			},
			want: fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s",
				TypeUnprocessableEntity, CodeUnprocessableEntity, "unprocessable entity"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &unprocessableEntityError{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("unprocessableEntityError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
