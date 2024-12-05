package sendgrid

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/education-english-web/BE-english-web/pkg/mailer"
)

func Test_sendGrid_buildMessageWithTemplate(t *testing.T) {
	t.Run("test build SGMailV3", func(t *testing.T) {
		templateID := "template_id"
		subject := "this is email subject"
		templateData := map[string]interface{}{
			"Bell Labs": "1234",
			"Google":    1234,
			"subject":   subject,
		}
		expectedTemplateData := map[string]interface{}{
			"Bell Labs": "1234",
			"Google":    1234,
			"subject":   subject,
		}
		sendGrid := sendGrid{
			env: "development",
			allowDomains: []string{
				"moneyforward.vn",
				"moneyforward.co.jp",
			},
			sender: mail.NewEmail("sender", "sender@email.com"),
		}
		expectS3Mail := &mail.SGMailV3{
			From: &mail.Email{
				Name:    "sender",
				Address: "sender@email.com",
			},
			Personalizations: []*mail.Personalization{
				{
					To: []*mail.Email{
						{
							Name:    "levi",
							Address: "levi@moneyforward.vn",
						},
						{
							Name:    "tonie",
							Address: "tonie@moneyforward.vn",
						},
					},
					DynamicTemplateData: expectedTemplateData,
				},
			},
			TemplateID:  "template_id",
			Content:     []*mail.Content{},
			Attachments: []*mail.Attachment{},
		}
		recipients := []mailer.Recipient{
			{
				Name:  "levi",
				Email: "levi@moneyforward.vn",
			},
			{
				Name:  "tonie",
				Email: "tonie@moneyforward.vn",
			},
		}
		s3mail := sendGrid.buildMessageWithTemplate(templateID, recipients, templateData)

		if diff := cmp.Diff(expectS3Mail, s3mail); diff != "" {
			t.Errorf(diff)
		}
	})

	t.Run("test build SGMailV3 with nil templateData", func(t *testing.T) {
		templateID := "template_id"
		subject := "this is email subject"
		templateData := map[string]interface{}{
			"subject": subject,
		}
		expectedTemplateData := map[string]interface{}{
			"subject": subject,
		}
		sendGrid := sendGrid{
			env: "development",
			allowDomains: []string{
				"moneyforward.vn",
				"moneyforward.co.jp",
			},
			sender: mail.NewEmail("sender", "sender@email.com"),
		}
		expectS3Mail := &mail.SGMailV3{
			From: &mail.Email{
				Name:    "sender",
				Address: "sender@email.com",
			},
			Personalizations: []*mail.Personalization{
				{
					To: []*mail.Email{
						{
							Name:    "levi",
							Address: "levi@moneyforward.co.jp",
						},
						{
							Name:    "tonie",
							Address: "tonie@moneyforward.co.jp",
						},
					},
					DynamicTemplateData: expectedTemplateData,
				},
			},
			TemplateID:  "template_id",
			Content:     []*mail.Content{},
			Attachments: []*mail.Attachment{},
		}
		recipients := []mailer.Recipient{
			{
				Name:  "levi",
				Email: "levi@moneyforward.co.jp",
			},
			{
				Name:  "tonie",
				Email: "tonie@moneyforward.co.jp",
			},
		}
		s3mail := sendGrid.buildMessageWithTemplate(templateID, recipients, templateData)

		if diff := cmp.Diff(expectS3Mail, s3mail); diff != "" {
			t.Errorf(diff)
		}
	})
}

func Test_sendGrid_isAllowed(t *testing.T) {
	type fields struct {
		env          string
		allowDomains []string
		client       *sendgrid.Client
		sender       *mail.Email
	}

	type args struct {
		email string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "empty email",
			fields: fields{
				env:          "development",
				allowDomains: nil,
				client:       nil,
				sender:       nil,
			},
			args: args{
				email: "",
			},
			want: false,
		},
		{
			name: "environment production",
			fields: fields{
				env:          "production",
				allowDomains: nil,
				client:       nil,
				sender:       nil,
			},
			args: args{
				email: "hello@example.com",
			},
			want: true,
		},
		{
			name: "allowDomains is empty list",
			fields: fields{
				env:          "development",
				allowDomains: nil,
				client:       nil,
				sender:       nil,
			},
			args: args{
				email: "hello@example.com",
			},
			want: true,
		},
		{
			name: "email string has no @",
			fields: fields{
				env: "development",
				allowDomains: []string{
					"moneyforward.vn",
					"moneyforward.co.jp",
				},
				client: nil,
				sender: nil,
			},
			args: args{
				email: "helloexample.com",
			},
			want: false,
		},
		{
			name: "email is in white list",
			fields: fields{
				env: "development",
				allowDomains: []string{
					"moneyforward.vn",
					"moneyforward.co.jp",
				},
				client: nil,
				sender: nil,
			},
			args: args{
				email: "nhan@moneyforward.vn",
			},
			want: true,
		},
		{
			name: "email is not in white list on development",
			fields: fields{
				env: "development",
				allowDomains: []string{
					"moneyforward.vn",
					"moneyforward.co.jp",
				},
				client: nil,
				sender: nil,
			},
			args: args{
				email: "nhan@moneyforward.com.vn",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sg := &sendGrid{
				env:          tt.fields.env,
				allowDomains: tt.fields.allowDomains,
				client:       tt.fields.client,
				sender:       tt.fields.sender,
			}

			if got := sg.isAllowed(tt.args.email); got != tt.want {
				t.Errorf("sendGrid.isAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		env          string
		allowDomains []string
		senderName   string
		senderEmail  string
		apiKey       string
	}

	tests := []struct {
		name string
		args args
		want mailer.Mailer
	}{
		{
			name: "success",
			args: args{
				env: "development",
				allowDomains: []string{
					"moneyforward.vn",
					"moneyforward.co.jp",
				},
				senderName:  "Sender name",
				senderEmail: "sender@email.com",
				apiKey:      "api_key",
			},
			want: &sendGrid{
				env: "development",
				allowDomains: []string{
					"moneyforward.vn",
					"moneyforward.co.jp",
				},
				client: sendgrid.NewSendClient("api_key"),
				sender: mail.NewEmail("Sender name", "sender@email.com"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.env, tt.args.allowDomains, tt.args.senderName, tt.args.senderEmail, tt.args.apiKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sendGrid_SendWithTemplate(t *testing.T) {
	type fields struct {
		env          string
		allowDomains []string
		client       *sendgrid.Client
		sender       *mail.Email
	}

	type args struct {
		templateID   string
		recipients   []mailer.Recipient
		templateData mailer.EmailData
		attachments  []mailer.Attachment
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				env: "development",
				allowDomains: []string{
					"google.com",
				},
				client: &sendgrid.Client{},
				sender: &mail.Email{
					Name:    "",
					Address: "",
				},
			},
			args: args{
				recipients: []mailer.Recipient{
					{
						Name:  "frank1",
						Email: "frakn1@google.com",
					},
					{
						Name:  "frank1",
						Email: "frakn1@google1.com",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "all filtered out",
			fields: fields{
				env: "development",
				allowDomains: []string{
					"google.com",
				},
				client: &sendgrid.Client{},
				sender: &mail.Email{
					Name:    "",
					Address: "",
				},
			},
			args: args{
				recipients: []mailer.Recipient{
					{
						Name:  "frank1",
						Email: "frakn1@google0.com",
					},
					{
						Name:  "frank1",
						Email: "frakn1@google1.com",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sg := &sendGrid{
				env:          tt.fields.env,
				allowDomains: tt.fields.allowDomains,
				client:       tt.fields.client,
				sender:       tt.fields.sender,
			}

			if err := sg.SendWithTemplate(tt.args.templateID, tt.args.recipients, tt.args.templateData, tt.args.attachments...); (err != nil) != tt.wantErr {
				t.Errorf("sendGrid.SendWithTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
