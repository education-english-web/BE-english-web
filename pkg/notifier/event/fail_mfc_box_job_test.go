package event

import (
	"reflect"
	"testing"
)

func TestNewEventFailMFCBoxJob(t *testing.T) {
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
				messageTemplate: "msg",
				channels:        []Channel{},
			},
			want: &eventFailMFCBoxJob{
				env:             "dev",
				messageTemplate: "msg",
				channels:        []Channel{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventFailMFCBoxJob(tt.args.env, tt.args.messageTemplate, tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventFailMFCBoxJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventFailMFCBoxJob_Name(t *testing.T) {
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
			name: "event_Fail_MFC_Box_Job",
			want: EventNameFailMFCBoxJob,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventFailMFCBoxJob{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Name(); got != tt.want {
				t.Errorf("eventFailMFCBoxJob.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventFailMFCBoxJob_Channels(t *testing.T) {
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
				messageTemplate: "msg",
				channels: []Channel{
					{
						Envs:       "dev",
						Name:       "channel",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: []Channel{
				{
					Envs:       "dev",
					Name:       "channel",
					WebhookURL: "hooks.slack.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventFailMFCBoxJob{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventFailMFCBoxJob.Channels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventFailMFCBoxJob_Env(t *testing.T) {
	type fields struct {
		env             string
		messageTemplate string
		channels        []Channel
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "success",
			fields: fields{env: "dev"},
			want:   "dev",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventFailMFCBoxJob{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventFailMFCBoxJob.Env() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventFailMFCBoxJob_BuildMessage(t *testing.T) {
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
				messageTemplate: "fqdn: %s\noffice_identification_code: %s\nconcluded_contract_id: %s\napplicant_mfid_uid: %s",
				channels: []Channel{
					{
						Envs:       "dev",
						Name:       "channel",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"fqdn":                       "localhost",
					"office_identification_code": "1",
					"concluded_contract_id":      "001",
					"applicant_mfid_uid":         "111",
				},
			},
			want: "fqdn: localhost\noffice_identification_code: 1\nconcluded_contract_id: 001\napplicant_mfid_uid: 111",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventFailMFCBoxJob{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.BuildMessage(tt.args.payload); got != tt.want {
				t.Errorf("eventFailMFCBoxJob.BuildMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
