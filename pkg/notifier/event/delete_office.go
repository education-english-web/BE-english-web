package event

import (
	"fmt"
	"strings"
)

// eventDeleteOffice holds information for the event when an office is deleted by an operator
type eventDeleteOffice struct {
	env             string
	messageTemplate string
	channels        []Channel
}

// NewEventDeleteOffice returns an instance of Event interface
func NewEventDeleteOffice(env, messageTemplate string, channels []Channel) Event {
	return &eventDeleteOffice{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

// Name return the name of event
func (e *eventDeleteOffice) Name() Name {
	return EventNameDeleteOffice
}

// Channels returns all channels that message will be sent to
func (e *eventDeleteOffice) Channels() []Channel {
	return e.channels
}

// Env returns the environment
func (e *eventDeleteOffice) Env() string {
	return e.env
}

// BuildMessage builds a notification message from the payload
func (e *eventDeleteOffice) BuildMessage(payload map[string]interface{}) string {
	return fmt.Sprintf(e.messageTemplate,
		strings.ToUpper(e.env), payload["email"], payload["office_identification_code"], payload["office_name"], payload["time"])
}
