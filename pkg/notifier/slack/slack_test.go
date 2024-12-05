package slack

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	slackClient "github.com/slack-go/slack"

	appConf "github.com/education-english-web/BE-english-web/app/config"
	"github.com/education-english-web/BE-english-web/pkg/notifier"
	"github.com/education-english-web/BE-english-web/pkg/notifier/event"
	mockEvent "github.com/education-english-web/BE-english-web/pkg/notifier/event/mock"
)

func Test_slack_Notify(t *testing.T) {
	t.Run("failed", func(t *testing.T) {
		backupPostWebhook := postWebhook
		defer func() {
			postWebhook = backupPostWebhook
		}()

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		loginEvent := mockEvent.NewMockEvent(mockCtrl)
		msg := "message"

		channels := []event.Channel{
			{
				Name:       "general",
				WebhookURL: "hooks.slack.com",
				Envs:       appConf.ENVDevelopment,
			},
		}
		payload := map[string]interface{}{}

		loginEvent.EXPECT().Channels().Return(channels)
		loginEvent.EXPECT().Env().Return(appConf.ENVDevelopment)
		loginEvent.EXPECT().BuildMessage(payload).Return(msg)

		slack := &slack{}

		err := errors.New("unexpected error")
		postWebhook = func(string, *slackClient.WebhookMessage) error {
			return err
		}

		wantErr := fmt.Errorf("error while sending message to slack: %w", err)
		gotErr := slack.Notify(loginEvent, payload)
		if gotErr == nil || gotErr.Error() != wantErr.Error() {
			t.Errorf("Test_slack_Notify error mismatch\ngot: %v\nwant: %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		backupPostWebhook := postWebhook
		defer func() {
			postWebhook = backupPostWebhook
		}()

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		loginEvent := mockEvent.NewMockEvent(mockCtrl)
		msg := "message"

		channels := []event.Channel{
			{
				Name:       "general",
				WebhookURL: "production.hooks.slack.com",
				Envs:       appConf.ENVProduction,
			},
			{
				Name:       "general",
				WebhookURL: "hooks.slack.com",
				Envs:       appConf.ENVDevelopment,
			},
		}
		payload := map[string]interface{}{}

		loginEvent.EXPECT().Channels().Return(channels)
		loginEvent.EXPECT().Env().Return(appConf.ENVDevelopment)
		loginEvent.EXPECT().Env().Return(appConf.ENVDevelopment)
		loginEvent.EXPECT().BuildMessage(payload).Return(msg)

		slack := &slack{}

		postWebhook = func(string, *slackClient.WebhookMessage) error {
			return nil
		}

		gotErr := slack.Notify(loginEvent, payload)
		if gotErr != nil {
			t.Errorf("Test_slack_Notify error mismatch\ngot: %v\nwant: nil", gotErr)
		}
	})
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want notifier.Notifier
	}{
		{
			name: "success",
			want: &slack{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSlack() = %v, want %v", got, tt.want)
			}
		})
	}
}
