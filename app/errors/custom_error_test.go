package errors

import (
	"strings"
	"testing"
)

func TestSystemErrors_Error(t *testing.T) {
	error1 := NewErrorInvalidArgument(CodeInvalidOfficeID, "invalid office ID", "")
	error2 := NewErrorInvalidArgument(CodeInvalidContractID, "invalid contract ID", "")
	errs := []SystemError{error1, error2}
	errsString := make([]string, 0, len(errs))

	for _, err := range errs {
		errsString = append(errsString, err.Error())
	}

	expectedError := strings.Join(errsString, "\n")

	tests := []struct {
		name string
		errs SystemErrors
		want string
	}{
		{
			name: "Test SystemError.Error()",
			errs: []SystemError{
				error1,
				error2,
			},
			want: expectedError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.errs.Error(); got != tt.want {
				t.Errorf("SystemErrors.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
