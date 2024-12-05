package event

import (
	"fmt"
	"strings"
)

// eventInternalUserAdd holds information for the event when an internal user is added to aweb
type eventInternalUserAdd struct {
	env             string
	messageTemplate string
	channels        []Channel
}

// NewEventInternalUserAdd returns an instance of Event interface
func NewEventInternalUserAdd(env, messageTemplate string, channels []Channel) Event {
	return &eventInternalUserAdd{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

// Name return the name of event
func (e *eventInternalUserAdd) Name() Name {
	return EventNameInternalUserAdd
}

// Channels returns all channels that message will be sent to
func (e *eventInternalUserAdd) Channels() []Channel {
	return e.channels
}

// Env returns the environment
func (e *eventInternalUserAdd) Env() string {
	return e.env
}

// BuildMessage builds a notification message from the payload
func (e *eventInternalUserAdd) BuildMessage(payload map[string]interface{}) string {
	return fmt.Sprintf(e.messageTemplate,
		strings.ToUpper(e.env), payload["email"], payload["target"], payload["role"], payload["time"])
}
