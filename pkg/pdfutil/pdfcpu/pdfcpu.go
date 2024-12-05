package pdfcpu

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pkg/errors"

	"github.com/education-english-web/BE-english-web/pkg/pdfutil"
)

type pdfCPU struct {
	config *model.Configuration
}

var (
	instance pdfCPU
	once     sync.Once
	fontName string
)

// Setup pdfCPU
func Setup() {
	once.Do(func() {
		config := model.NewDefaultConfiguration()
		config.ValidationMode = model.ValidationRelaxed
		_ = api.InstallFonts([]string{"assets/BIZUDPGothic-Regular.ttf"})
		fontName = "BIZUDPGothic-Regular"
		instance = pdfCPU{config}
	})
}

// New initiates pdf cpu tool
func New() pdfutil.PDFFactory {
	return &instance
}

//nolint:nonamedreturns
func (p *pdfCPU) ExtractPageNumber(r io.Reader) (total int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	// create a temporary file for the reading
	f, err := os.CreateTemp("", "temp_pdf")
	if err != nil {
		return 0, errors.Wrap(err, "error while create temp file")
	}

	// destroy the file once done
	defer func() {
		n := f.Name()
		_ = f.Close()
		_ = os.Remove(n)
	}()

	// transfer the bytes to the file
	if _, err := io.Copy(f, r); err != nil {
		return 0, errors.Wrap(err, "error while copy bytes to temp file")
	}

	pdfcpuCtx, err := pdfcpu.Read(f, p.config)
	if err != nil {
		return 0, errors.Wrap(err, "error while reading pdf file")
	}

	if err := api.OptimizeContext(pdfcpuCtx); err != nil {
		return 0, errors.Wrap(err, "error while optimizing a pdf context")
	}

	// PageCount
	if pdfcpuCtx.PageCount == 0 {
		return 0, errors.Wrap(err, "error while getting pdf page count")
	}

	return pdfcpuCtx.PageCount, nil
}

// Validate validates a PDF file with pdfcpu tool
func (p *pdfCPU) Validate(r io.ReadSeeker) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	buf := bytes.NewBuffer(nil)

	wm := model.DefaultWatermarkConfig()
	wm.TextString = " "

	if err := api.AddWatermarks(
		r,
		buf,
		[]string{"1"},
		wm,
		model.NewDefaultConfiguration(),
	); err != nil {
		return fmt.Errorf("add sample text: %w", err)
	}

	return nil
}

//nolint:nonamedreturns
func (p *pdfCPU) Dimensions(inFile string) (width, height float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	dims, err := api.PageDimsFile(inFile)
	if err != nil || len(dims) == 0 {
		return 0, 0, errors.Wrap(err, "error while getting pdf file dimensions")
	}

	return dims[0].Width, dims[0].Height, nil
}

func (p *pdfCPU) AddCustomField(inFile, text string, page, posX, posY int, scale float32) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	return api.AddTextWatermarksFile(
		inFile,
		inFile,
		[]string{strconv.Itoa(page)},
		true,
		text,
		fmt.Sprintf("font:%s,rot:0,scale:%.1f abs,offset:%d %d,al:l,pos:tl,fillcolor:#000000,ma:1", fontName, scale, posX, posY*-1),
		nil,
	)
}

// AddCheckIcon adds a check icon into a PDF file
func (p *pdfCPU) AddCheckIcon(inFile string, page, posX, posY int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	return api.AddTextWatermarksFile(
		inFile,
		inFile,
		[]string{strconv.Itoa(page)},
		true,
		"\063",
		fmt.Sprintf("font:ZapfDingbats,rot:0,scale:0.6 abs,offset:%d %d,al:l,pos:tl,fillcolor:#000000", posX, posY*-1),
		nil,
	)
}

// AddImage adds an image from filepath into a PDF file with absolute scale factor
func (p *pdfCPU) AddImage(inFile, image string, page, posX, posY int, scale float32) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	return api.AddImageWatermarksFile(
		inFile,
		inFile,
		[]string{strconv.Itoa(page)},
		true,
		image,
		fmt.Sprintf("rot:0,scale:%.3f abs,offset:%d %d,pos:tl", scale, posX, posY*-1),
		nil,
	)
}

// AddImageFromReader add image from reader into a PDF file with relative scale factor
func (p *pdfCPU) AddImageFromReader(inFile string, image io.Reader, page, posX, posY int, relScale float32) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	return api.AddImageWatermarksForReaderFile(
		inFile,
		inFile,
		[]string{strconv.Itoa(page)},
		true,
		image,
		fmt.Sprintf("rot:0,scale:%.3f rel,offset:%d %d,pos:tl", relScale, posX, posY*-1),
		nil,
	)
}

func (p *pdfCPU) AddFooter(inFile, text string, page int, scale float32) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	return api.AddTextWatermarksFile(
		inFile,
		inFile,
		[]string{strconv.Itoa(page)},
		true,
		text,
		fmt.Sprintf("font:%s,rot:0,scale:%.1f abs,al:l,pos:bl,fillcolor:#000000,ma:10", fontName, scale),
		nil,
	)
}
