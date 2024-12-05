package event

import (
	"reflect"
	"testing"
)

func TestNewEventProxyLogin(t *testing.T) {
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
			want: &eventProxyLogin{
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
			if got := NewEventProxyLogin(tt.args.env, tt.args.messageTemplate, tt.args.channels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventProxyLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventProxyLogin_Name(t *testing.T) {
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
			want: EventNameProxyLogin,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventProxyLogin{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.Name(); got != tt.want {
				t.Errorf("eventProxyLogin.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventProxyLogin_Channels(t *testing.T) {
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
			e := &eventProxyLogin{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventProxyLogin.Channels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventProxyLogin_Env(t *testing.T) {
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
			e := &eventProxyLogin{env: tt.args.env}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventProxyLogin.Env \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventProxyLogin_BuildMessage(t *testing.T) {
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
				messageTemplate: "user %s do proxy login with identification code: %s, tenant user uid: %s, reason: %s",
				channels: []Channel{
					{
						Name:       "general",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"email":               "email@example.com",
					"identification_code": "navis office identification code",
					"tenant_user_uid":     "tenant user uid",
					"reason":              "reason for proxy login",
				},
			},
			want: "user email@example.com do proxy login with identification code: navis office identification code, tenant user uid: tenant user uid, reason: reason for proxy login",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventProxyLogin{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}

			if got := e.BuildMessage(tt.args.payload); got != tt.want {
				t.Errorf("eventProxyLogin.BuildMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
