package errors

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewErrorForbidden(t *testing.T) {
	tests := []struct {
		name string
		want SystemError
	}{
		{
			name: "success",
			want: &errorForbidden{
				code:    CodeForbidden,
				message: "forbidden",
				param:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrorForbidden(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorForbidden() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorForbidden_Type(t *testing.T) {
	type fields struct {
		code    Code
		message string
		param   interface{}
	}

	tests := []struct {
		name   string
		fields fields
		want   TypeError
	}{
		{
			name: "success",
			fields: fields{
				code:    "code",
				message: "message",
				param:   nil,
			},
			want: TypeForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorForbidden{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Type(); got != tt.want {
				t.Errorf("errorForbidden.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorForbidden_Code(t *testing.T) {
	type fields struct {
		code    Code
		message string
		param   interface{}
	}

	tests := []struct {
		name   string
		fields fields
		want   Code
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeForbidden,
				message: "forbidden",
				param:   nil,
			},
			want: CodeForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorForbidden{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Code(); got != tt.want {
				t.Errorf("errorForbidden.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorForbidden_Message(t *testing.T) {
	type fields struct {
		code    Code
		message string
		param   interface{}
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeForbidden,
				message: "forbidden",
				param:   nil,
			},
			want: "forbidden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorForbidden{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Message(); got != tt.want {
				t.Errorf("errorForbidden.Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorForbidden_Param(t *testing.T) {
	type fields struct {
		code    Code
		message string
		param   interface{}
	}

	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeForbidden,
				message: "forbidden",
				param:   1,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorForbidden{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Param(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorForbidden.Param() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorForbidden_StatusCode(t *testing.T) {
	type fields struct {
		code    Code
		message string
		param   interface{}
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeForbidden,
				message: "forbidden",
				param:   1,
			},
			want: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorForbidden{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.StatusCode(); got != tt.want {
				t.Errorf("errorForbidden.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorForbidden_Error(t *testing.T) {
	type fields struct {
		code    Code
		message string
		param   interface{}
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeUserInsufficientAuthorization,
				message: "user is not authorized for deleting other users",
				param:   nil,
			},
			want: fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s, \tParam: %v",
				TypeForbidden, CodeUserInsufficientAuthorization, "user is not authorized for deleting other users", nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorForbidden{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("errorForbidden.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCustomErrorForbidden(t *testing.T) {
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
				code:    CodeUserInsufficientAuthorization,
				message: "user is not authorized for deleting other users",
				param:   nil,
			},
			want: &errorForbidden{
				code:    CodeUserInsufficientAuthorization,
				message: "user is not authorized for deleting other users",
				param:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomErrorForbidden(tt.args.code, tt.args.message, tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomErrorForbidden() = %v, want %v", got, tt.want)
			}
		})
	}
}
