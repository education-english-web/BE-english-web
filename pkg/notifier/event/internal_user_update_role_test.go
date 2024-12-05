package event

import (
	"reflect"
	"testing"
)

func TestNewEventInternalUserUpdateRole(t *testing.T) {
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
				messageTemplate: "admin change the role of user %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: &eventInternalUserUpdateRole{
				env:             "dev",
				messageTemplate: "admin change the role of user %s",
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
			if got := NewEventInternalUserUpdateRole(tt.args.env, tt.args.messageTemplate, tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventInternalUserUpdateRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserUpdateRole_Name(t *testing.T) {
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
				messageTemplate: "admin change the role of user %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: EventNameInternalUserUpdateRole,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventInternalUserUpdateRole{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Name(); got != tt.want {
				t.Errorf("eventInternalUserUpdateRole.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserUpdateRole_Channels(t *testing.T) {
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
				messageTemplate: "admin change the role of user %s",
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
			e := &eventInternalUserUpdateRole{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventInternalUserUpdateRole.Channels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserUpdateRole_Env(t *testing.T) {
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
			e := &eventInternalUserUpdateRole{env: tt.args.env}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventInternalUserUpdateRole.Env \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserUpdateRole_BuildMessage(t *testing.T) {
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
				messageTemplate: "%s: user %s changes the role of user %s from %s to %s in %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"email":    "mail@example.com",
					"target":   "other.mail@example.com",
					"old_role": "normal",
					"new_role": "operator",
					"time":     "yesterday",
				},
			},
			want: "DEV: user mail@example.com changes the role of user other.mail@example.com from normal to operator in yesterday",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventInternalUserUpdateRole{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.BuildMessage(tt.args.payload); got != tt.want {
				t.Errorf("eventInternalUserUpdateRole.BuildMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
