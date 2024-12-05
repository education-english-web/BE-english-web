package wkhtmltopdf

import (
	"fmt"
	"io"

	wrapper "github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"github.com/education-english-web/BE-english-web/pkg/htmltopdf"
)

type wkHTMLToPDF struct{}

func New() htmltopdf.Converter {
	return &wkHTMLToPDF{}
}

func (c *wkHTMLToPDF) Convert(inHTML io.Reader, outPDF io.Writer) error {
	pdg, err := wrapper.NewPDFGenerator()
	if err != nil {
		return fmt.Errorf("find wkhtmltopdf in system & create wrapper: %w", err)
	}

	page := wrapper.NewPageReader(inHTML)

	pdg.AddPage(page)

	if err := pdg.Create(); err != nil {
		return fmt.Errorf("generate html to pdf: %w", err)
	}

	if _, err := outPDF.Write(pdg.Buffer().Bytes()); err != nil {
		return fmt.Errorf("write data to out : %w", err)
	}

	return nil
}
