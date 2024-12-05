package eventfactory

//go:generate mockgen -destination=./mock/$GOFILE -source=$GOFILE -package=mock

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/education-english-web/BE-english-web/pkg/notifier/event"
)

var globalConfig EventMapping

// EventFactory provides a method to get notification event by its name
type EventFactory interface {
	GetEventByName(name event.Name) event.Event
}

// eventFactory implements EventFactory interface
type eventFactory struct {
	mEventByName map[event.Name]event.Event
}

// EventMapping holds mapping information for events that send message to slack
type EventMapping map[string]EventConfig

//nolint:tagliatelle
type EventConfig struct {
	MessageTemplate string          `yaml:"message_template"`
	Channels        []event.Channel `yaml:"channels"`
}

// GetEventByName returns the notification event by its name
func (e *eventFactory) GetEventByName(name event.Name) event.Event {
	return e.mEventByName[name]
}

// LoadEventConfig reads slack config information from yaml file
func LoadEventConfig(filePath string) error {
	// #nosec
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error while reading config file: %w", err)
	}

	if err := yaml.Unmarshal(bytes, &globalConfig); err != nil {
		return fmt.Errorf("error while unmarshaling data: %w", err)
	}

	return nil
}

// GetEventConfig returns slack config information for sending a notification
func GetEventConfig() EventMapping {
	return globalConfig
}

// NewEventFactory returns an instance of EventFactory
func NewEventFactory(events ...event.Event) EventFactory {
	mEventByName := make(map[event.Name]event.Event)

	for _, e := range events {
		mEventByName[e.Name()] = e
	}

	return &eventFactory{
		mEventByName: mEventByName,
	}
}
