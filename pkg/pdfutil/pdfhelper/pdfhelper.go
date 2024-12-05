package pdfhelper

import "github.com/education-english-web/BE-english-web/pkg/pdfutil"

type pdfHelper struct{}

func New() pdfutil.PDFHelper {
	return &pdfHelper{}
}

func (p *pdfHelper) FontScale(width float64) float32 {
	// these standard numbers are collected from testing experiment on some PDF files,
	// we found them sufficient to print texts onto PDF
	var (
		standardWidth     float64 = 1215
		standardTextScale         = 0.55

		scalePercentage = 1.3
	)

	widthDiff := width / standardWidth
	scale := widthDiff * standardTextScale * scalePercentage

	return float32(scale)
}

func (p *pdfHelper) FontScaleByFontSize(width, fontSize float64) float32 {
	// these standard numbers are collected from testing experiment on some PDF files,
	// we found them sufficient to print texts onto PDF
	var (
		standardWidth     float64 = 1215
		standardTextScale         = 0.55

		scalePercentage = 1.3 / 9.75 // 1pt ~ 1.3(3)px, 9.75pt ~ 13px
	)

	widthDiff := width / standardWidth
	scale := widthDiff * standardTextScale * scalePercentage * fontSize

	return float32(scale)
}

func (p *pdfHelper) MarginScale(width float64, posX, posY int, hasIconInFront bool) (int, int) {
	// these standard numbers are collected from testing experiment on some PDF files,
	// we found them sufficient to print texts onto PDF
	var (
		standardWith    = 595
		standardMarginX = 7
		standardMarginY = 8
	)

	widthDiff := int(width) / standardWith

	posX, posY = posX+widthDiff*standardMarginX, posY+widthDiff*standardMarginY

	// For some custom field types (such as DateTime), from UI there will be an icon in front of the string value (date string)
	// this X adjustment is for pushing the text to the right, to make it consistent with the UI
	if hasIconInFront {
		posX += widthDiff * 18
	}

	return posX, posY
}
