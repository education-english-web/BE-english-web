package htmltemplate

import "io"

//go:generate mockgen -destination=./mock/$GOFILE -source=$GOFILE -package=mock

type TemplateFunc struct {
	Name string
	Fn   interface{}
}

type Parser interface {
	Parse(out io.Writer, filePath string, data map[string]interface{}, templateFunc ...TemplateFunc) error
}
