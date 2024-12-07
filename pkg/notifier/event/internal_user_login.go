package event

import (
	"fmt"
	"strings"
)

// eventInternalUserLogin holds information for the notification event when an internal user logs in to Aweb
type eventInternalUserLogin struct {
	env             string
	messageTemplate string
	channels        []Channel
}

// NewEventInternalUserLogin returns an instance of Event interface
func NewEventInternalUserLogin(env, messageTemplate string, channels []Channel) Event {
	return &eventInternalUserLogin{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

// Name return the name of event
func (e *eventInternalUserLogin) Name() Name {
	return EventNameInternalUserLogin
}

// Channels returns all channels that message will be sent to
func (e *eventInternalUserLogin) Channels() []Channel {
	return e.channels
}

// Env returns the environment
func (e *eventInternalUserLogin) Env() string {
	return e.env
}

// BuildMessage builds a notification message from the payload
func (e *eventInternalUserLogin) BuildMessage(payload map[string]interface{}) string {
	return fmt.Sprintf(e.messageTemplate, strings.ToUpper(e.env), payload["email"], payload["time"])
}
