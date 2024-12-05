package event

import (
	"reflect"
	"testing"
)

func TestNewUpdateUserSetting(t *testing.T) {
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
				env:             "development",
				messageTemplate: "message",
				channels:        []Channel{},
			},
			want: &eventUpdateUserSetting{
				env:             "development",
				messageTemplate: "message",
				channels:        []Channel{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventUpdateUserSetting(
				tt.args.env,
				tt.args.messageTemplate,
				tt.args.channels,
			); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventUpdateUserSetting = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventUpdateUserSetting_Name(t *testing.T) {
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
				messageTemplate: "message",
				channels:        []Channel{},
			},
			want: EventNameUpdateUserSetting,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventUpdateUserSetting{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Name(); got != tt.want {
				t.Errorf("eventUpdateUserSetting.Name = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventUpdateUserSetting_Channels(t *testing.T) {
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
				env:             "development",
				messageTemplate: "message",
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
			e := &eventUpdateUserSetting{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventUpdateUserSetting.Channels = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventUpdateUserSetting_Env(t *testing.T) {
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
			e := &eventUpdateUserSetting{env: tt.args.env}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventUpdateUserSetting.Env \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventUpdateUserSetting_BuildMessage(t *testing.T) {
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
				env:             "development",
				messageTemplate: "Environment: %s\nEmail: %s\nAction: Change User Setting for Email notification\nTarget: %s - %s\nValue after change: %s\nTime: %s\n",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"email":                    "email@example.com",
					"target_user_id":           "123321",
					"target_user_display_name": "user display name",
					"value_after_change":       "OFF",
					"time":                     "now",
				},
			},
			want: "Environment: DEVELOPMENT\nEmail: email@example.com\nAction: Change User Setting for Email notification\nTarget: 123321 - user display name\nValue after change: OFF\nTime: now\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventUpdateUserSetting{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.BuildMessage(tt.args.payload); got != tt.want {
				t.Errorf("eventUpdateUserSetting.BuildMessage = %v, want %v", got, tt.want)
			}
		})
	}
}
