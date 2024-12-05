package pdfutil

//go:generate mockgen -destination=./mock/$GOFILE -source=$GOFILE -package=mock

import (
	"io"
)

// PDFFactory provides comprehensive methods to operate with pdf file using pdfcpu package
type PDFFactory interface {
	ExtractPageNumber(r io.Reader) (total int, err error)
	Validate(r io.ReadSeeker) error
	Dimensions(inFile string) (width float64, height float64, err error)
	AddCustomField(inFile, text string, page, posX, posY int, scale float32) error
	AddCheckIcon(inFile string, page, posX, posY int) error
	AddFooter(inFile, text string, page int, scale float32) error
	AddImage(inFile, image string, page, posX, posY int, scale float32) error
	AddImageFromReader(inFile string, image io.Reader, page, posX, posY int, relativeScale float32) error
}

// PDFHelper provides methods to operate with pdf file
type PDFHelper interface {
	FontScale(width float64) float32
	FontScaleByFontSize(width, fontSize float64) float32
	MarginScale(width float64, posX, posY int, hasIconInFront bool) (x, y int)
}
