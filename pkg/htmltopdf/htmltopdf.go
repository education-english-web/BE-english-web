package htmltopdf

//go:generate mockgen -destination=./mock/$GOFILE -source=$GOFILE -package=mock

import "io"

type Converter interface {
	Convert(inHTML io.Reader, outPDF io.Writer) error
}
