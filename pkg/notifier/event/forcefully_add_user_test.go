package event

import (
	"reflect"
	"testing"
)

func TestNewEventForcefullyAddUser(t *testing.T) {
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
			want: &eventForcefullyAddUser{
				env:             "development",
				messageTemplate: "message",
				channels:        []Channel{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventForcefullyAddUser(tt.args.env, tt.args.messageTemplate, tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventForcefullyAddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventForcefullyAddUser_Name(t *testing.T) {
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
			want: EventNameForcefullyAddUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventForcefullyAddUser{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.Name(); got != tt.want {
				t.Errorf("eventForcefullyAddUser.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventForcefullyAddUser_Channels(t *testing.T) {
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
				messageTemplate: "user %s is added",
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
			e := &eventForcefullyAddUser{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventForcefullyAddUser.Channels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventForcefullyAddUser_Env(t *testing.T) {
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
			e := &eventForcefullyAddUser{env: tt.args.env}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventForcefullyAddUser.Env \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventForcefullyAddUser_BuildMessage(t *testing.T) {
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
				messageTemplate: "%s: email %s forcefully added into office %s - %s user %s on aweb at %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"email":                      "email@example.com",
					"office_identification_code": "1111111",
					"office_name":                "Office Name",
					"user_email":                 "user@email.com",
					"time":                       "night",
				},
			},
			want: "DEVELOPMENT: email email@example.com forcefully added into office 1111111 - Office Name user user@email.com on aweb at night",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventForcefullyAddUser{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.BuildMessage(tt.args.payload); got != tt.want {
				t.Errorf("eventForcefullyAddUser.BuildMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
