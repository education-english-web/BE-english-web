package event

import (
	"fmt"
)

// eventFailMFCBoxJob holds information for the event when the mfc box job process failed
type eventFailMFCBoxJob struct {
	env             string
	messageTemplate string
	channels        []Channel
}

func NewEventFailMFCBoxJob(env, messageTemplate string, channels []Channel) Event {
	return &eventFailMFCBoxJob{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

func (e *eventFailMFCBoxJob) Name() Name {
	return EventNameFailMFCBoxJob
}

func (e *eventFailMFCBoxJob) Channels() []Channel {
	return e.channels
}

func (e *eventFailMFCBoxJob) Env() string {
	return e.env
}

func (e *eventFailMFCBoxJob) BuildMessage(payload map[string]interface{}) string {
	return fmt.Sprintf(
		e.messageTemplate,
		payload["fqdn"],
		payload["office_identification_code"],
		payload["concluded_contract_id"],
		payload["applicant_mfid_uid"],
	)
}
