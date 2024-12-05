package googlesheet

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/sheets/v4"

	"github.com/education-english-web/BE-english-web/pkg/spreadsheet"
)

const (
	rangeTemplate        = "%s!%s"
	maxRetry             = 3
	waitTimeBetweenRetry = 1 * time.Minute
)

type client struct {
	service       *sheets.Service
	spreadSheetID string
}

func New(service *sheets.Service, spreadSheetID string) spreadsheet.Client {
	return &client{
		service:       service,
		spreadSheetID: spreadSheetID,
	}
}

func (c *client) CreateSheet(ctx context.Context, sheetName string) error {
	req := sheets.Request{
		AddSheet: &sheets.AddSheetRequest{
			Properties: &sheets.SheetProperties{
				Title: sheetName,
			},
		},
	}
	rbb := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{&req},
	}

	retries := 0

PROCESS:

	if _, err := c.service.Spreadsheets.
		BatchUpdate(c.spreadSheetID, rbb).
		Context(ctx).
		Do(); err != nil {
		if googleapiError, ok := err.(*googleapi.Error); ok {
			switch googleapiError.Code {
			case 400:
				return spreadsheet.ErrBadRequest
			case 503:
				retries++
				if retries < maxRetry {
					time.Sleep(waitTimeBetweenRetry)

					goto PROCESS
				}
			}
		}

		return err
	}

	return nil
}

func (c *client) AppendRows(ctx context.Context, sheetName, rowRange string, records [][]interface{}) error {
	valueRange := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         records,
	}

	retries := 0
	rangeFmt := fmt.Sprintf(rangeTemplate, sheetName, rowRange)

PROCESS:

	if _, err := c.service.Spreadsheets.Values.
		Append(c.spreadSheetID, rangeFmt, valueRange).
		ValueInputOption("USER_ENTERED").
		Context(ctx).
		Do(); err != nil {
		if googleapiError, ok := err.(*googleapi.Error); ok && googleapiError.Code == 503 {
			retries++
			if retries < maxRetry {
				time.Sleep(waitTimeBetweenRetry)

				goto PROCESS
			}
		}

		return err
	}

	return nil
}

func (c *client) ClearSheet(ctx context.Context, sheetName string) error {
	retries := 0

PROCESS:

	if _, err := c.service.Spreadsheets.Values.
		Clear(c.spreadSheetID, sheetName, &sheets.ClearValuesRequest{}).
		Context(ctx).
		Do(); err != nil {
		if googleapiError, ok := err.(*googleapi.Error); ok && googleapiError.Code == 503 {
			retries++
			if retries < maxRetry {
				time.Sleep(waitTimeBetweenRetry)

				goto PROCESS
			}
		}

		return err
	}

	return nil
}
