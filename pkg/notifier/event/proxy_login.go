package event

import (
	"fmt"
)

// eventProxyLogin holds information for the notification event when an internal user do proxy log in
type eventProxyLogin struct {
	env             string
	messageTemplate string
	channels        []Channel
}

// NewEventProxyLogin returns an instance of Event interface
func NewEventProxyLogin(env, messageTemplate string, channels []Channel) Event {
	return &eventProxyLogin{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

// Name return the name of event
func (e *eventProxyLogin) Name() Name {
	return EventNameProxyLogin
}

// Channels returns all channels that message will be sent to
func (e *eventProxyLogin) Channels() []Channel {
	return e.channels
}

// Env returns the environment
func (e *eventProxyLogin) Env() string {
	return e.env
}

// BuildMessage builds a notification message from the payload
func (e *eventProxyLogin) BuildMessage(payload map[string]interface{}) string {
	return fmt.Sprintf(
		e.messageTemplate,
		payload["email"],
		payload["identification_code"],
		payload["tenant_user_uid"],
		payload["reason"],
	)
}
