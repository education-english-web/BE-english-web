package logkeeper

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import "time"

// LogKeeper provides a method to push log to its store
type LogKeeper interface {
	Push(message string, timestamp time.Time) error
}
