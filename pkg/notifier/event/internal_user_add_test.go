package event

import (
	"reflect"
	"testing"
)

func TestNewEventInternalUserAdd(t *testing.T) {
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
				messageTemplate: "user %s is added",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: &eventInternalUserAdd{
				env:             "development",
				messageTemplate: "user %s is added",
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
			if got := NewEventInternalUserAdd(tt.args.env, tt.args.messageTemplate, tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventInternalUserAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserAdd_Name(t *testing.T) {
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
				env:             "development",
				messageTemplate: "user %s is added",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: EventNameInternalUserAdd,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventInternalUserAdd{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Name(); got != tt.want {
				t.Errorf("eventInternalUserAdd.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserAdd_Channels(t *testing.T) {
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
			e := &eventInternalUserAdd{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventInternalUserAdd.Channels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserAdd_Env(t *testing.T) {
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
			e := &eventInternalUserAdd{env: tt.args.env}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventInternalUserAdd.Env \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventInternalUserAdd_BuildMessage(t *testing.T) {
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
				messageTemplate: "%s: user %s add user %s with role %s to aweb at %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"email":  "email@example.com",
					"target": "other.email@example.com",
					"role":   "operator",
					"time":   "night",
				},
			},
			want: "DEVELOPMENT: user email@example.com add user other.email@example.com with role operator to aweb at night",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventInternalUserAdd{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.BuildMessage(tt.args.payload); got != tt.want {
				t.Errorf("eventInternalUserAdd.BuildMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
