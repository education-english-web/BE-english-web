package errors

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewErrorUnauthorized(t *testing.T) {
	tests := []struct {
		name string
		want SystemError
	}{
		{
			name: "success",
			want: &errorUnauthorized{
				code:    CodeUnauthorized,
				message: "unauthorized",
				param:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrorUnauthorized(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorUnauthorized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorUnauthorized_Type(t *testing.T) {
	tests := []struct {
		name string
		e    *errorUnauthorized
		want TypeError
	}{
		{
			name: "success",
			e:    &errorUnauthorized{},
			want: TypeUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorUnauthorized{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.Type(); got != tt.want {
				t.Errorf("errorUnauthorized.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorUnauthorized_Code(t *testing.T) {
	tests := []struct {
		name string
		e    *errorUnauthorized
		want Code
	}{
		{
			name: "success",
			e: &errorUnauthorized{
				code:    CodeUnauthorized,
				message: "unauthorized",
				param:   nil,
			},
			want: CodeUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorUnauthorized{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.Code(); got != tt.want {
				t.Errorf("errorUnauthorized.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorUnauthorized_Message(t *testing.T) {
	tests := []struct {
		name string
		e    *errorUnauthorized
		want string
	}{
		{
			name: "success",
			e: &errorUnauthorized{
				code:    CodeUnauthorized,
				message: "unauthorized",
				param:   nil,
			},
			want: "unauthorized",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorUnauthorized{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.Message(); got != tt.want {
				t.Errorf("errorUnauthorized.Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorUnauthorized_Param(t *testing.T) {
	tests := []struct {
		name string
		e    *errorUnauthorized
		want interface{}
	}{
		{
			name: "success",
			e: &errorUnauthorized{
				code:    CodeUnauthorized,
				message: "unauthorized",
				param:   nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorUnauthorized{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.Param(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorUnauthorized.Param() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorUnauthorized_StatusCode(t *testing.T) {
	tests := []struct {
		name string
		e    *errorUnauthorized
		want int
	}{
		{
			name: "success",
			e: &errorUnauthorized{
				code:    CodeUnauthorized,
				message: "unauthorized",
				param:   nil,
			},
			want: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorUnauthorized{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.StatusCode(); got != tt.want {
				t.Errorf("errorUnauthorized.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorUnauthorized_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *errorUnauthorized
		want string
	}{
		{
			name: "success",
			e: &errorUnauthorized{
				code:    CodeUnauthorized,
				message: "unauthorized",
				param:   nil,
			},
			want: fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s",
				TypeUnauthorized, CodeUnauthorized, "unauthorized"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorUnauthorized{
				code:    tt.e.code,
				message: tt.e.message,
				param:   tt.e.param,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("errorUnauthorized.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCustomErrorUnauthorized(t *testing.T) {
	type args struct {
		code    Code
		message string
		param   interface{}
	}

	tests := []struct {
		name string
		args args
		want SystemError
	}{
		{
			name: "success",
			args: args{
				code:    CodeUnauthorized,
				message: "unauthorized",
				param:   nil,
			},
			want: &errorUnauthorized{
				code:    CodeUnauthorized,
				message: "unauthorized",
				param:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomErrorUnauthorized(tt.args.code, tt.args.message, tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomErrorUnauthorized() = %v, want %v", got, tt.want)
			}
		})
	}
}
