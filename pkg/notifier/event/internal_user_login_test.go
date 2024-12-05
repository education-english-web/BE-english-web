package event

import (
	"reflect"
	"testing"
)

func TestNewEventInternalUserLogin(t *testing.T) {
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
				messageTemplate: "user %s logged in",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: &eventInternalUserLogin{
				env:             "dev",
				messageTemplate: "user %s logged in",
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
			if got := NewEventInternalUserLogin(tt.args.env, tt.args.messageTemplate, tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventInternalUserLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserLogin_Name(t *testing.T) {
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
				messageTemplate: "user %s logged in",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: EventNameInternalUserLogin,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventInternalUserLogin{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Name(); got != tt.want {
				t.Errorf("eventInternalUserLogin.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserLogin_Channels(t *testing.T) {
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
				messageTemplate: "user %s logged in",
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
			e := &eventInternalUserLogin{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventInternalUserLogin.Channels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserLogin_Env(t *testing.T) {
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
			e := &eventInternalUserLogin{env: tt.args.env}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventInternalUserLogin.Env \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserLogin_BuildMessage(t *testing.T) {
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
				messageTemplate: "%s: user %s logged in %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"email": "email@example.com",
					"time":  "this morning",
				},
			},
			want: "DEV: user email@example.com logged in this morning",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventInternalUserLogin{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.BuildMessage(tt.args.payload); got != tt.want {
				t.Errorf("eventInternalUserLogin.BuildMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
