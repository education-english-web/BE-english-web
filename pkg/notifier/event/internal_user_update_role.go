package event

import (
	"fmt"
	"strings"
)

// eventInternalUserUpdateRole holds information for notification event when an admin changes internal user role
type eventInternalUserUpdateRole struct {
	env             string
	messageTemplate string
	channels        []Channel
}

// NewEventInternalUserUpdateRole returns an instance of Event interface
func NewEventInternalUserUpdateRole(env, messageTemplate string, channels []Channel) Event {
	return &eventInternalUserUpdateRole{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

// Name return the name of event
func (e *eventInternalUserUpdateRole) Name() Name {
	return EventNameInternalUserUpdateRole
}

// Channels returns all channels that message will be sent to
func (e *eventInternalUserUpdateRole) Channels() []Channel {
	return e.channels
}

// Env returns the environment
func (e *eventInternalUserUpdateRole) Env() string {
	return e.env
}

// BuildMessage builds a notification message from the payload
func (e *eventInternalUserUpdateRole) BuildMessage(payload map[string]interface{}) string {
	return fmt.Sprintf(e.messageTemplate,
		strings.ToUpper(e.env), payload["email"], payload["target"], payload["old_role"], payload["new_role"], payload["time"])
}
