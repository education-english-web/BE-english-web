package spreadsheet

import (
	"context"
	"errors"
)

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

var ErrBadRequest = errors.New("bad request")

// Client to work with a spreadsheet
type Client interface {
	CreateSheet(ctx context.Context, sheetName string) error
	AppendRows(ctx context.Context, sheetName, rowRange string, records [][]interface{}) error
	ClearSheet(ctx context.Context, sheetName string) error
}
