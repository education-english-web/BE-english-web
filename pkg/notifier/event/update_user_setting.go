package event

import (
	"fmt"
	"strings"
)

// eventUpdateUserSetting holds information for the notification event when an internal user update user setting
type eventUpdateUserSetting struct {
	env             string
	messageTemplate string
	channels        []Channel
}

// NewEventUpdateUserSetting return an instance of Event interface
func NewEventUpdateUserSetting(env, messageTemplate string, channels []Channel) Event {
	return &eventUpdateUserSetting{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

// Name return the name of event
func (e *eventUpdateUserSetting) Name() Name {
	return EventNameUpdateUserSetting
}

// Channels returns all channels that message will be sent to
func (e *eventUpdateUserSetting) Channels() []Channel {
	return e.channels
}

// Env returns the environment
func (e *eventUpdateUserSetting) Env() string {
	return e.env
}

// BuildMessage builds a notification message from the payload
func (e *eventUpdateUserSetting) BuildMessage(payload map[string]interface{}) string {
	splitMsg := strings.Split(e.messageTemplate, "\n")
	msgSlice := make([]string, 0, len(splitMsg))

	for i := range splitMsg {
		if !strings.Contains(splitMsg[i], "%") || strings.Contains(splitMsg[i], "Environment") {
			msgSlice = append(msgSlice, splitMsg[i])

			continue
		}

		for key := range payload {
			if strings.Contains(strings.ToLower(splitMsg[i]), strings.Split(key, "_")[0]) {
				msgSlice = append(msgSlice, splitMsg[i])

				break
			}
		}
	}

	msgTemplate := strings.Join(msgSlice, "\n")

	return fmt.Sprintf(msgTemplate,
		strings.ToUpper(e.env),
		payload["email"],
		payload["target_user_id"],
		payload["target_user_display_name"],
		payload["value_after_change"],
		payload["time"],
	)
}
