package event

import (
	"fmt"
	"strings"
)

// eventForcefullyAddUser holds information for the event when an office is deleted by an operator
type eventForcefullyAddUser struct {
	env             string
	messageTemplate string
	channels        []Channel
}

// NewEventForcefullyAddUser returns an instance of Event interface
func NewEventForcefullyAddUser(env, messageTemplate string, channels []Channel) Event {
	return &eventForcefullyAddUser{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

// Name return the name of event
func (e *eventForcefullyAddUser) Name() Name {
	return EventNameForcefullyAddUser
}

// Channels returns all channels that message will be sent to
func (e *eventForcefullyAddUser) Channels() []Channel {
	return e.channels
}

// Env returns the environment
func (e *eventForcefullyAddUser) Env() string {
	return e.env
}

// BuildMessage builds a notification message from the payload
func (e *eventForcefullyAddUser) BuildMessage(payload map[string]interface{}) string {
	return fmt.Sprintf(e.messageTemplate,
		strings.ToUpper(e.env), payload["email"], payload["office_identification_code"], payload["office_name"], payload["user_email"], payload["time"])
}
