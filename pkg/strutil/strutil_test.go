package strutil

import "testing"

func TestLeftPad(t *testing.T) {
	type args struct {
		str  string
		n    int
		char rune
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				str:  "1",
				n:    8,
				char: '0',
			},
			want: "00000001",
		},
		{
			name: "Case 2",
			args: args{
				str:  "0000001",
				n:    8,
				char: '0',
			},
			want: "00000001",
		},
		{
			name: "Case 3",
			args: args{
				str:  "00000001",
				n:    8,
				char: '0',
			},
			want: "00000001",
		},
		{
			name: "Case 4",
			args: args{
				str:  "00000001",
				n:    0,
				char: '0',
			},
			want: "00000001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeftPad(tt.args.str, tt.args.n, tt.args.char); got != tt.want {
				t.Errorf("LeftPad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberFormat(t *testing.T) {
	type args struct {
		n         float64
		precision int
		isEmpty   bool
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case empty",
			args: args{
				n:         0,
				precision: 0,
				isEmpty:   true,
			},
			want: "",
		},
		{
			name: "case 0",
			args: args{
				n:         0,
				precision: 0,
				isEmpty:   false,
			},
			want: "0",
		},
		{
			name: "case 1",
			args: args{
				n:         0,
				precision: 1,
				isEmpty:   false,
			},
			want: "0.0",
		},
		{
			name: "case 2",
			args: args{
				n:         700,
				precision: 2,
				isEmpty:   false,
			},
			want: "700.00",
		},
		{
			name: "case 3",
			args: args{
				n:         7000.2345,
				precision: 3,
				isEmpty:   false,
			},
			want: "7,000.235",
		},
		{
			name: "case 4",
			args: args{
				n:         7000.235,
				precision: 4,
				isEmpty:   false,
			},
			want: "7,000.2350",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberFormat(tt.args.n, tt.args.precision, tt.args.isEmpty); got != tt.want {
				t.Errorf("NumberFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
