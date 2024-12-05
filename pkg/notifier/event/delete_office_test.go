package event

import (
	"reflect"
	"testing"
)

func TestNewEventDeleteOffice(t *testing.T) {
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
			want: &eventDeleteOffice{
				env:             "development",
				messageTemplate: "message",
				channels:        []Channel{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventDeleteOffice(tt.args.env, tt.args.messageTemplate, tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventDeleteOffice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventDeleteOffice_Name(t *testing.T) {
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
			want: EventNameDeleteOffice,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventDeleteOffice{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Name(); got != tt.want {
				t.Errorf("eventDeleteOffice.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventDeleteOffice_Channels(t *testing.T) {
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
			e := &eventDeleteOffice{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventDeleteOffice.Channels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventDeleteOffice_Env(t *testing.T) {
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
			e := &eventDeleteOffice{env: tt.args.env}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventDeleteOffice.Env \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventDeleteOffice_BuildMessage(t *testing.T) {
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
				messageTemplate: "%s: email %s delete office %s - %s on aweb at %s",
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
					"time":                       "night",
				},
			},
			want: "DEVELOPMENT: email email@example.com delete office 1111111 - Office Name on aweb at night",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventDeleteOffice{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.BuildMessage(tt.args.payload); got != tt.want {
				t.Errorf("eventDeleteOffice.BuildMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
