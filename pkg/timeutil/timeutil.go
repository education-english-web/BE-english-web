package timeutil

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import "time"

//nolint:gomnd
var (
	VST = time.FixedZone("Asia/Ho_Chi_Minh", 7*60*60)
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)

const (
	ISOLayout       = "2006-01-02"
	YearMonthLayout = "2006-01"
	SheetDataLayout = "2006/01/02"
)

// TimeFactory represents timekeeper
type TimeFactory interface {
	Now() time.Time
}

type timeFactory struct{}

// NewTimeFactory initiates a time factory
func NewTimeFactory() TimeFactory {
	return &timeFactory{}
}

// Now returns current time
func (f *timeFactory) Now() time.Time {
	return time.Now()
}
