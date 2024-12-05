package notifier

//go:generate mockgen -destination=./mock/$GOFILE -source=$GOFILE -package=mock

import "github.com/education-english-web/BE-english-web/pkg/notifier/event"

// Notifier provides a method for sending a notification to channels when an event occurs
type Notifier interface {
	Notify(event event.Event, payload map[string]interface{}) error
}
