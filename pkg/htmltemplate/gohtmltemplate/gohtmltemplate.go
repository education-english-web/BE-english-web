package gohtmltemplate

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"text/template"

	"github.com/education-english-web/BE-english-web/pkg/htmltemplate"
)

type goHTMLTemplate struct{}

func New() htmltemplate.Parser {
	return &goHTMLTemplate{}
}

func (t *goHTMLTemplate) Parse(
	out io.Writer,
	filePath string,
	data map[string]interface{},
	templateFuncs ...htmltemplate.TemplateFunc,
) error {
	parsedHTML := bytes.Buffer{}
	templateFuncMap := make(template.FuncMap)
	_, filename := filepath.Split(filePath)

	for i := range templateFuncs {
		templateFuncMap[templateFuncs[i].Name] = templateFuncs[i].Fn
	}

	temp, err := template.New(filename).Funcs(templateFuncMap).ParseFiles(filePath)
	if err != nil {
		return fmt.Errorf("html template parse file: %w", err)
	}

	if err := temp.Execute(&parsedHTML, data); err != nil {
		return fmt.Errorf("html template map data: %w", err)
	}

	if _, err := out.Write(parsedHTML.Bytes()); err != nil {
		return fmt.Errorf("write data to out : %w", err)
	}

	return nil
}
