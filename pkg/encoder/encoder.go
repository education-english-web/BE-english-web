package encoder

//go:generate mockgen -destination=./mock/$GOFILE -source=$GOFILE -package=mock

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/gogs/chardet"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var ErrCharsetUndetected = errors.New("undetected charset")

const (
	charsetUTF8     = "utf-8"
	charsetShiftJIS = "shift_jis"

	confidentAcceptedNumber         = 80
	confidentShiftJISAcceptedNumber = 50
)

type Encoder interface {
	TransformToUTF8(r io.Reader) (io.Reader, error)
}

type encoder struct{}

func NewEncoder() Encoder {
	return encoder{}
}

// TransformToUTF8 transform
func (e encoder) TransformToUTF8(r io.Reader) (io.Reader, error) {
	var buff bytes.Buffer

	tee := io.TeeReader(r, &buff)

	bs, err := io.ReadAll(tee)
	if err != nil {
		return nil, fmt.Errorf("error while reading: %w", err)
	}

	rs, err := chardet.NewTextDetector().DetectBest(bs)
	if err != nil {
		return nil, fmt.Errorf("error while detecting charset: %w", err)
	}

	newReader := bytes.NewReader(buff.Bytes())

	// UTF8
	if strings.EqualFold(rs.Charset, charsetUTF8) && rs.Confidence >= confidentAcceptedNumber {
		return newReader, nil
	}

	// Shift JIS
	if strings.EqualFold(rs.Charset, charsetShiftJIS) && rs.Confidence >= confidentShiftJISAcceptedNumber {
		return transform.NewReader(newReader, japanese.ShiftJIS.NewDecoder()), nil
	}

	return nil, ErrCharsetUndetected
}
