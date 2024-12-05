package runeutil

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Test_NewWriter(t *testing.T) {
	type args struct {
		w        io.Writer
		encoding string
	}

	buff := bytes.NewBuffer(nil)
	tests := []struct {
		name string
		args args
		want *RuneWriter
	}{
		{
			name: "success - shift-jis",
			args: args{
				w:        buff,
				encoding: "shift-jis",
			},
			want: &RuneWriter{
				w: transform.NewWriter(buff, japanese.ShiftJIS.NewEncoder()),
			},
		},
		{
			name: "success - utf-8",
			args: args{
				w:        buff,
				encoding: "utf-8",
			},
			want: &RuneWriter{
				w: transform.NewWriter(buff, unicode.UTF8.NewEncoder()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWriter(tt.args.w, tt.args.encoding); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriter: \ngot: %v - \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_Write(t *testing.T) {
	type args struct {
		b               []byte
		encodingSetting string
	}

	type result struct {
		n int
		e error
		b []byte
	}

	tests := []struct {
		name string
		args args
		want result
	}{
		{
			name: "empty bytes",
			args: args{
				b:               []byte(""),
				encodingSetting: ShiftJIS,
			},
			want: result{
				n: 0,
				e: nil,
				b: []byte(""),
			},
		},
		{
			name: "sjis - unsupported character encoding",
			args: args{
				b:               []byte("婷"),
				encodingSetting: ShiftJIS,
			},
			want: result{
				n: 3,
				e: nil,
				b: []byte("?"),
			},
		},
		{
			name: "utf8 - success",
			args: args{
				b:               []byte("締婷"),
				encodingSetting: UTF8,
			},
			want: result{
				n: 6,
				e: nil,
				b: []byte("締婷"),
			},
		},
		{
			name: "sjis - success",
			args: args{
				b:               []byte("締婷"),
				encodingSetting: ShiftJIS,
			},
			want: result{
				n: 6,
				e: nil,
				b: []byte("締?"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			rw := NewWriter(buf, tt.args.encodingSetting)
			gotN, gotErr := rw.Write(tt.args.b)
			if gotErr != nil && gotErr.Error() != tt.want.e.Error() {
				t.Errorf("RuneWriter.Write: \ngot: %v \nwant: %v\n", gotErr, tt.want.e)

				return
			}

			if gotN != tt.want.n {
				t.Errorf("RuneWriter.Write: \ngot: %v \nwant: %v\n", gotN, tt.want.n)

				return
			}

			// check result
			b, _ := io.ReadAll(buf)
			if tt.args.encodingSetting == ShiftJIS {
				bb, err := japanese.ShiftJIS.NewDecoder().Bytes(b)
				if err != nil {
					t.Errorf("shift-jis decoder got error: %v\n", err)

					return
				}

				b = bb
			}

			if !reflect.DeepEqual(b, tt.want.b) {
				t.Errorf("RuneWriter.Write: \ngot:%v \nwant: %v\n", b, tt.want.b)
			}
		})
	}
}
