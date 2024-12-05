package pdfhelper

import "testing"

func Test_pdfCPU_MarginScale(t *testing.T) {
	type args struct {
		width          float64
		posX           int
		posY           int
		hasIconInFront bool
	}

	tests := []struct {
		name     string
		args     args
		wantPosX int
		wantPosY int
	}{
		{
			name: "success - no icon in front",
			args: args{
				width:          1000,
				posX:           100,
				posY:           200,
				hasIconInFront: false,
			},
			wantPosX: 107,
			wantPosY: 208,
		},
		{
			name: "success - no icon in front",
			args: args{
				width:          1200,
				posX:           100,
				posY:           200,
				hasIconInFront: false,
			},
			wantPosX: 114,
			wantPosY: 216,
		},
		{
			name: "success - has icon in front",
			args: args{
				width:          1215,
				posX:           100,
				posY:           200,
				hasIconInFront: true,
			},
			wantPosX: 150,
			wantPosY: 216,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pdfHelper{}
			got, got1 := p.MarginScale(tt.args.width, tt.args.posX, tt.args.posY, tt.args.hasIconInFront)
			if got != tt.wantPosX {
				t.Errorf("marginScale() got = %v, want %v", got, tt.wantPosX)
			}
			if got1 != tt.wantPosY {
				t.Errorf("marginScale() got1 = %v, want %v", got1, tt.wantPosY)
			}
		})
	}
}

func Test_pdfCPU_FontScale(t *testing.T) {
	type args struct {
		width float64
	}

	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "success",
			args: args{
				width: float64(100),
			},
			want: float32(0.058847737),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pdfHelper{}
			if got := p.FontScale(tt.args.width); got != tt.want {
				t.Errorf("pdfCPU.FontScale: \ngot: %v \nwant: %f", got, tt.want)

				return
			}
		})
	}
}

func Test_pdfCPU_FontScaleByFontSize(t *testing.T) {
	type args struct {
		width    float64
		fontSize float64
	}

	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "success",
			args: args{
				width:    float64(100),
				fontSize: float64(30),
			},
			want: float32(0.18106996),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pdfHelper{}
			if got := p.FontScaleByFontSize(tt.args.width, tt.args.fontSize); got != tt.want {
				t.Errorf("pdfCPU.FontScaleByFontSize: \ngot: %v \nwant: %f", got, tt.want)

				return
			}
		})
	}
}
