package event

import (
	"reflect"
	"testing"
)

func TestNewCreateOfficeUsage(t *testing.T) {
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
			want: &eventCreateOfficeUsage{
				env:             "development",
				messageTemplate: "message",
				channels:        []Channel{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventCreateOfficeUsage(
				tt.args.env,
				tt.args.messageTemplate,
				tt.args.channels,
			); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventCreateOfficeUsage = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventCreateOfficeUsage_Name(t *testing.T) {
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
			want: EventNameCreateOfficeUsage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventCreateOfficeUsage{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Name(); got != tt.want {
				t.Errorf("eventCreateOfficeUsage.Name = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventCreateOfficeUsage_Channels(t *testing.T) {
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
			e := &eventCreateOfficeUsage{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventCreateOfficeUsage.Channels = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventCreateOfficeUsage_Env(t *testing.T) {
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
			e := &eventCreateOfficeUsage{env: tt.args.env}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventCreateOfficeUsage.Env \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventCreateOfficeUsage_BuildMessage(t *testing.T) {
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
			name: "success - middle",
			fields: fields{
				env:             "development",
				messageTemplate: "Environment: %s\nEmail: %s\nAction: Change Contract Status\nTarget: %s - %s\nPlan: %s\nExpired date: %s\nStatus: %s\nRelevant documents: %s\nSF linkage: %s\nReason: %s\nTime: %s\n",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"email":                             "email@example.com",
					"target_office_identification_code": "ident code",
					"target_office_name":                "office name",
					"plan":                              "Middle",
					"expired_date":                      "end date",
					"status":                            "Trial",
					"relevant_documents":                "no use",
					"sf_linkage":                        "no use",
					"reason":                            "for testing purpose",
					"time":                              "now",
				},
			},
			want: "Environment: DEVELOPMENT\nEmail: email@example.com\nAction: Change Contract Status\nTarget: ident code - office name\nPlan: Middle\nExpired date: end date\nStatus: Trial\nRelevant documents: no use\nSF linkage: no use\nReason: for testing purpose\nTime: now\n",
		},
		{
			name: "success - web",
			fields: fields{
				env:             "development",
				messageTemplate: "Environment: %s\nEmail: %s\nAction: Change Contract Status\nTarget: %s - %s\nPlan: %s\nReason: %s\nTime: %s\n",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"email":                             "email@example.com",
					"target_office_identification_code": "ident code",
					"target_office_name":                "office name",
					"plan":                              "Web",
					"reason":                            "for testing purpose",
					"time":                              "now",
				},
			},
			want: "Environment: DEVELOPMENT\nEmail: email@example.com\nAction: Change Contract Status\nTarget: ident code - office name\nPlan: Web\nReason: for testing purpose\nTime: now\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventCreateOfficeUsage{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.BuildMessage(tt.args.payload); got != tt.want {
				t.Errorf("eventCreateOfficeUsage.BuildMessage = %v, want %v", got, tt.want)
			}
		})
	}
}
