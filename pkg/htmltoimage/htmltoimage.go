package htmltoimage

//go:generate mockgen -destination=./mock/$GOFILE -source=$GOFILE -package=mock

import "io"

type Generator interface {
	GenerateByID(html, elementID string, out io.Writer) error
}
