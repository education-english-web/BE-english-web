package errors

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewErrorNotFound(t *testing.T) {
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
				code:    CodeNotFoundContract,
				message: "something not found",
				param:   "id",
			},
			want: &errorNotFound{
				code:    CodeNotFoundContract,
				message: "something not found",
				param:   "id",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrorNotFound(tt.args.code, tt.args.message, tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorNotFound_Type(t *testing.T) {
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
				code:    CodeNotFoundContract,
				message: "something not found",
				param:   "id",
			},
			want: TypeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorNotFound{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Type(); got != tt.want {
				t.Errorf("errorNotFound.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorNotFound_Code(t *testing.T) {
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
				code:    CodeNotFoundContract,
				message: "something not found",
				param:   "id",
			},
			want: CodeNotFoundContract,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorNotFound{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Code(); got != tt.want {
				t.Errorf("errorNotFound.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorNotFound_Message(t *testing.T) {
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
				code:    CodeNotFoundContract,
				message: "something not found",
				param:   "id",
			},
			want: "something not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorNotFound{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Message(); got != tt.want {
				t.Errorf("errorNotFound.Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorNotFound_Param(t *testing.T) {
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
				code:    CodeNotFoundContract,
				message: "something not found",
				param:   "id",
			},
			want: "id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorNotFound{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Param(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorNotFound.Param() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorNotFound_StatusCode(t *testing.T) {
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
				code:    CodeNotFoundContract,
				message: "something not found",
				param:   "id",
			},
			want: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorNotFound{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.StatusCode(); got != tt.want {
				t.Errorf("errorNotFound.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorNotFound_Error(t *testing.T) {
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
				code:    CodeNotFoundContract,
				message: "something not found",
				param:   "id",
			},
			want: fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s, \tParam: %v",
				TypeNotFound, CodeNotFoundContract, "something not found", "id"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorNotFound{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("errorNotFound.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
