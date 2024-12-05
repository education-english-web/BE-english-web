package errors

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewErrorConflict(t *testing.T) {
	type args struct {
		code    Code
		message string
	}

	tests := []struct {
		name string
		args args
		want SystemError
	}{
		{
			name: "success",
			args: args{
				code:    CodeConflictApprovalSession,
				message: "conflict approval session",
			},
			want: &errorConflict{
				code:    CodeConflictApprovalSession,
				message: "conflict approval session",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrorConflict(tt.args.code, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorConflict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorConflict_Type(t *testing.T) {
	type fields struct {
		code    Code
		message string
	}

	tests := []struct {
		name   string
		fields fields
		want   TypeError
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeConflictApprovalSession,
				message: "conflict approval session",
			},
			want: TypeConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorConflict{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.Type(); got != tt.want {
				t.Errorf("errorConflict.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorConflict_Code(t *testing.T) {
	type fields struct {
		code    Code
		message string
	}

	tests := []struct {
		name   string
		fields fields
		want   Code
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeConflictApprovalSession,
				message: "conflict approval session",
			},
			want: CodeConflictApprovalSession,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorConflict{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.Code(); got != tt.want {
				t.Errorf("errorConflict.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorConflict_Message(t *testing.T) {
	type fields struct {
		code    Code
		message string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeConflictApprovalSession,
				message: "conflict approval session",
			},
			want: "conflict approval session",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorConflict{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.Message(); got != tt.want {
				t.Errorf("errorConflict.Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorConflict_Param(t *testing.T) {
	type fields struct {
		code    Code
		message string
	}

	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeConflictApprovalSession,
				message: "conflict approval session",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorConflict{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.Param(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorConflict.Param() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorConflict_StatusCode(t *testing.T) {
	type fields struct {
		code    Code
		message string
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeConflictApprovalSession,
				message: "conflict approval session",
			},
			want: http.StatusConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorConflict{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.StatusCode(); got != tt.want {
				t.Errorf("errorConflict.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorConflict_Error(t *testing.T) {
	type fields struct {
		code    Code
		message string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				code:    CodeConflictApprovalSession,
				message: "conflict approval session",
			},
			want: fmt.Sprintf("Type: %v, \tCode: %v, \tMessage: %s",
				TypeConflict, CodeConflictApprovalSession, "conflict approval session"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorConflict{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("errorConflict.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
