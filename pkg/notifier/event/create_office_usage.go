package event

import (
	"fmt"
	"strings"
)

// eventCreateOfficeUsage holds information for the notification event when an internal user create office usage
type eventCreateOfficeUsage struct {
	env             string
	messageTemplate string
	channels        []Channel
}

// NewEventCreateOfficeUsage return an instance of Event interface
func NewEventCreateOfficeUsage(env, messageTemplate string, channels []Channel) Event {
	return &eventCreateOfficeUsage{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

// Name return the name of event
func (e *eventCreateOfficeUsage) Name() Name {
	return EventNameCreateOfficeUsage
}

// Channels returns all channels that message will be sent to
func (e *eventCreateOfficeUsage) Channels() []Channel {
	return e.channels
}

// Env returns the environment
func (e *eventCreateOfficeUsage) Env() string {
	return e.env
}

// BuildMessage builds a notification message from the payload
func (e *eventCreateOfficeUsage) BuildMessage(payload map[string]interface{}) string {
	splitMsg := strings.Split(e.messageTemplate, "\n")
	msgSlice := make([]string, 0, len(splitMsg))

	for _, msgField := range splitMsg {
		if !strings.Contains(msgField, "%") || strings.Contains(msgField, "Environment") {
			msgSlice = append(msgSlice, msgField)

			continue
		}

		for key := range payload {
			if strings.Contains(strings.ToLower(msgField), strings.Split(key, "_")[0]) {
				msgSlice = append(msgSlice, msgField)

				break
			}
		}
	}

	msgTemplate := strings.Join(msgSlice, "\n")

	if payload["plan"] == "Middle" {
		return fmt.Sprintf(msgTemplate,
			strings.ToUpper(e.env),
			payload["email"],
			payload["target_office_identification_code"],
			payload["target_office_name"],
			payload["plan"],
			payload["expired_date"],
			payload["status"],
			payload["relevant_documents"],
			payload["sf_linkage"],
			payload["reason"],
			payload["time"],
		)
	}

	return fmt.Sprintf(msgTemplate,
		strings.ToUpper(e.env),
		payload["email"],
		payload["target_office_identification_code"],
		payload["target_office_name"],
		payload["plan"],
		payload["reason"],
		payload["time"],
	)
}
