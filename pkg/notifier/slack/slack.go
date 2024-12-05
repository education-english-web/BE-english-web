package slack

import (
	"fmt"
	"strings"

	slackClient "github.com/slack-go/slack"

	"github.com/education-english-web/BE-english-web/pkg/notifier"
	"github.com/education-english-web/BE-english-web/pkg/notifier/event"
)

var postWebhook = slackClient.PostWebhook

// slack is an implementation of Notifier interface
type slack struct{}

// New returns a instance of Notifier interface
func New() notifier.Notifier {
	return &slack{}
}

// Notify sends messages to receivers
func (s *slack) Notify(e event.Event, payload map[string]interface{}) error {
	channels := e.Channels()
	msg := e.BuildMessage(payload)

	for _, channel := range channels {
		if !strings.Contains(channel.Envs, e.Env()) {
			continue
		}

		if err := postWebhook(channel.WebhookURL, &slackClient.WebhookMessage{
			Channel: channel.Name,
			Text:    msg,
		}); err != nil {
			return fmt.Errorf("error while sending message to slack: %w", err)
		}
	}

	return nil
}
