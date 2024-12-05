package errors

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewErrorInvalidArgument(t *testing.T) {
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
				code:    CodeInvalidContractID,
				message: "invalid contract id",
				param:   111,
			},
			want: &errorInvalidArgument{
				code:    CodeInvalidContractID,
				message: "invalid contract id",
				param:   111,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrorInvalidArgument(tt.args.code, tt.args.message, tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorInvalidArgument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorInvalidArgument_Type(t *testing.T) {
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
				code:    CodeInvalidContractID,
				message: "invalid contract id",
				param:   111,
			},
			want: TypeInvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorInvalidArgument{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Type(); got != tt.want {
				t.Errorf("errorInvalidArgument.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorInvalidArgument_Code(t *testing.T) {
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
				code:    CodeInvalidContractID,
				message: "invalid contract id",
				param:   111,
			},
			want: CodeInvalidContractID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorInvalidArgument{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Code(); got != tt.want {
				t.Errorf("errorInvalidArgument.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorInvalidArgument_Message(t *testing.T) {
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
				code:    CodeInvalidContractID,
				message: "invalid contract id",
				param:   111,
			},
			want: "invalid contract id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorInvalidArgument{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Message(); got != tt.want {
				t.Errorf("errorInvalidArgument.Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorInvalidArgument_Param(t *testing.T) {
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
				code:    CodeInvalidContractID,
				message: "invalid contract id",
				param:   111,
			},
			want: 111,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorInvalidArgument{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Param(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorInvalidArgument.Param() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorInvalidArgument_StatusCode(t *testing.T) {
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
				code:    CodeInvalidContractID,
				message: "invalid contract id",
				param:   111,
			},
			want: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorInvalidArgument{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.StatusCode(); got != tt.want {
				t.Errorf("errorInvalidArgument.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorInvalidArgument_Error(t *testing.T) {
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
				code:    CodeInvalidContractID,
				message: "invalid contract id",
				param:   111,
			},
			want: fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s, \tParam: %v",
				TypeInvalidArgument, CodeInvalidContractID, "invalid contract id", 111),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorInvalidArgument{
				code:    tt.fields.code,
				message: tt.fields.message,
				param:   tt.fields.param,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("errorInvalidArgument.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
