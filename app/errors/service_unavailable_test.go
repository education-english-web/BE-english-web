package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServiceUnavailableError(t *testing.T) {
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
				code:    CodeInternal,
				message: "internal error",
				param:   nil,
			},
			want: &serviceUnavailableError{
				code:    CodeInternal,
				message: "internal error",
				params:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewServiceUnavailableError(
				tt.args.code,
				tt.args.message,
				tt.args.param,
			)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_serviceUnavailableError(t *testing.T) {
	type want = struct {
		statusCode int
		typeError  TypeError
		code       Code
		message    string
		params     interface{}
		errorMsg   string
	}

	tests := []struct {
		name string
		obj  *serviceUnavailableError
		want want
	}{
		{
			name: "success",
			obj: &serviceUnavailableError{
				code:    CodeInternal,
				message: "internal error",
				params:  nil,
			},
			want: want{
				statusCode: http.StatusServiceUnavailable,
				typeError:  TypeServiceUnavailable,
				code:       CodeInternal,
				message:    "internal error",
				params:     nil,
				errorMsg:   "Type: TYPE_SERVICE_UNAVAILABLE, \tCode: CODE_INTERNAL, \tMessage: internal error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want.statusCode, tt.obj.StatusCode())
			assert.Equal(t, tt.want.code, tt.obj.Code())
			assert.Equal(t, tt.want.typeError, tt.obj.Type())
			assert.Equal(t, tt.want.message, tt.obj.Message())
			assert.Equal(t, tt.want.statusCode, tt.obj.StatusCode())
			assert.Equal(t, tt.want.params, tt.obj.Param())
			assert.Equal(t, tt.want.errorMsg, tt.obj.Error())
		})
	}
}
