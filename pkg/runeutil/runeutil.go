package runeutil

import (
	"fmt"
	"io"
	"unicode/utf8"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type RuneWriter struct {
	w io.WriteCloser
}

const (
	ShiftJIS = "shift-jis"
	UTF8     = "utf-8"
)

func NewWriter(w io.Writer, encodingSetting string) *RuneWriter {
	var enc *encoding.Encoder

	switch encodingSetting {
	case ShiftJIS:
		enc = japanese.ShiftJIS.NewEncoder()
	default:
		enc = unicode.UTF8.NewEncoder()
	}

	return &RuneWriter{
		w: transform.NewWriter(w, enc),
	}
}

func (rw *RuneWriter) Write(b []byte) (int, error) {
	var l int

	defer func() {
		_ = rw.w.Close()
	}()

	for len(b) > 0 {
		r, n := utf8.DecodeRune(b)
		if n == 0 {
			break
		}

		if _, err := rw.w.Write(b[:n]); err != nil {
			// replace unsupported character encoding with "?"
			if _, err := rw.w.Write([]byte("?")); err != nil {
				return l, fmt.Errorf("replace [?] for unsupported rune [%c]: %w", r, err)
			}
		}

		l += n
		b = b[n:]
	}

	return l, nil
}
