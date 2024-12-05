package comparison

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualPointers(t *testing.T) {
	val := 1
	tests := []struct {
		name       string
		a          *int
		b          *int
		wantResult bool
	}{
		{
			name:       "both nil",
			a:          nil,
			b:          nil,
			wantResult: true,
		},
		{
			name:       "a nil, b not",
			a:          nil,
			b:          &val,
			wantResult: false,
		},
		{
			name:       "a not, b nil",
			a:          &val,
			b:          nil,
			wantResult: false,
		},
		{
			name:       "both are not nil",
			a:          &val,
			b:          &val,
			wantResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := EqualPointers(tt.a, tt.b)
			assert.Equal(t, tt.wantResult, gotResult)
		})
	}
}
