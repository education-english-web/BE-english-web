package event

import (
	"fmt"
	"strings"
)

// eventTurnOffLegalCheck holds information for the notification event when legal check of an office is turned off
type eventTurnOffLegalCheck struct {
	env             string
	messageTemplate string
	channels        []Channel
}

// NewEventTurnOffLegalCheck returns an instance of Event interface
func NewEventTurnOffLegalCheck(env, messageTemplate string, channels []Channel) Event {
	return &eventTurnOffLegalCheck{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

// Name return the name of event
func (e *eventTurnOffLegalCheck) Name() Name {
	return EventNameTurnOffLegalCheck
}

// Channels returns all channels that message will be sent to
func (e *eventTurnOffLegalCheck) Channels() []Channel {
	return e.channels
}

// Env returns the environment
func (e *eventTurnOffLegalCheck) Env() string {
	return e.env
}

// BuildMessage builds a notification message from the payload
func (e *eventTurnOffLegalCheck) BuildMessage(payload map[string]interface{}) string {
	return fmt.Sprintf(
		e.messageTemplate,
		strings.ToUpper(e.env),
		payload["office_identification_code"],
		payload["office_name"],
		payload["time"],
	)
}
