package event

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventTurnOffLegalCheck(t *testing.T) {
	type args struct {
		env             string
		messageTemplate string
		channels        []Channel
	}

	tests := []struct {
		name string
		args args
		want Event
	}{
		{
			name: "success",
			args: args{
				env:             "dev",
				messageTemplate: "turned off legal check for office %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: &eventTurnOffLegalCheck{
				env:             "dev",
				messageTemplate: "turned off legal check for office %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEventTurnOffLegalCheck(tt.args.env, tt.args.messageTemplate, tt.args.channels)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_eventTurnOffLegalCheck_Name(t *testing.T) {
	type fields struct {
		env             string
		messageTemplate string
		channels        []Channel
	}

	tests := []struct {
		name   string
		fields fields
		want   Name
	}{
		{
			name: "success",
			fields: fields{
				env:             "dev",
				messageTemplate: "turned off legal check for office %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: EventNameTurnOffLegalCheck,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventTurnOffLegalCheck{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Name(); got != tt.want {
				t.Errorf("eventTurnOffLegalCheck.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventTurnOffLegalCheck_Channels(t *testing.T) {
	type fields struct {
		env             string
		messageTemplate string
		channels        []Channel
	}

	tests := []struct {
		name   string
		fields fields
		want   []Channel
	}{
		{
			name: "success",
			fields: fields{
				env:             "dev",
				messageTemplate: "turned off legal check for office %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: []Channel{
				{
					Name:       "general",
					WebhookURL: "hooks.slack.com",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventTurnOffLegalCheck{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventTurnOffLegalCheck.Channels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventTurnOffLegalCheck_Env(t *testing.T) {
	type args struct {
		env string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{env: "development"},
			want: "development",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventTurnOffLegalCheck{env: tt.args.env}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventTurnOffLegalCheck.Env \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventTurnOffLegalCheck_BuildMessage(t *testing.T) {
	type fields struct {
		env             string
		messageTemplate string
		channels        []Channel
	}

	type args struct {
		payload map[string]interface{}
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "success",
			fields: fields{
				env:             "dev",
				messageTemplate: "On %s: Turned off legal check for office %s - %s at %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"office_identification_code": "123456",
					"office_name":                "Office Name",
					"time":                       "this morning",
				},
			},
			want: "On DEV: Turned off legal check for office 123456 - Office Name at this morning",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventTurnOffLegalCheck{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			got := e.BuildMessage(tt.args.payload)

			assert.Equal(t, tt.want, got)
		})
	}
}
