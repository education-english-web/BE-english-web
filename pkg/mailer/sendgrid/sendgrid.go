package sendgrid

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/education-english-web/BE-english-web/app/config"
	"github.com/education-english-web/BE-english-web/pkg/mailer"
)

const (
	maxRetries = 2
	delay      = time.Second
)

// sendGrid is helper for sendgrid
type sendGrid struct {
	env          string
	allowDomains []string
	client       *sendgrid.Client
	sender       *mail.Email
}

// New returns new instance of Mailer
func New(
	env string,
	allowDomains []string,
	senderName,
	senderEmail,
	apiKey string,
) mailer.Mailer {
	return &sendGrid{
		env:          env,
		allowDomains: allowDomains,
		client:       sendgrid.NewSendClient(apiKey),
		sender:       mail.NewEmail(senderName, senderEmail),
	}
}

func (sg *sendGrid) SendWithTemplate(templateID string, recipients []mailer.Recipient, templateData mailer.EmailData, attachments ...mailer.Attachment) error {
	filteredRecipients := make([]mailer.Recipient, 0, len(recipients))

	for _, to := range recipients {
		if !sg.isAllowed(to.Email) {
			continue
		}

		filteredRecipients = append(filteredRecipients, to)
	}

	if len(filteredRecipients) == 0 {
		return nil
	}

	message := sg.buildMessageWithTemplate(templateID, filteredRecipients, templateData, attachments...)

	var (
		retry    int
		err      error
		response *rest.Response
	)

	for {
		response, err = sg.client.Send(message)
		if err == nil {
			break
		}

		if retry > maxRetries || err.Error() != "Post \"https://api.sendgrid.com/v3/mail/send\": EOF" {
			return fmt.Errorf("send request: %w", err)
		}

		time.Sleep(delay * 1 << retry)

		retry++
	}

	if response.StatusCode > 299 {
		return fmt.Errorf("response: %s", response.Body)
	}

	return nil
}

// buildMessageWithTemplate build the message to be sent by sendgrid
func (sg *sendGrid) buildMessageWithTemplate(
	templateID string,
	recipients []mailer.Recipient,
	templateData mailer.EmailData,
	attachments ...mailer.Attachment,
) *mail.SGMailV3 {
	message := mail.NewV3Mail()
	message.SetTemplateID(templateID)
	message.SetFrom(sg.sender)

	for _, attachment := range attachments {
		message.AddAttachment(mail.NewAttachment().
			SetDisposition("attachment").
			SetFilename(attachment.Name).
			SetType(attachment.Type).
			SetContent(base64.StdEncoding.EncodeToString(attachment.Bytes)))
	}

	personalization := mail.Personalization{}
	tos := make([]*mail.Email, len(recipients))

	for i, to := range recipients {
		tos[i] = &mail.Email{
			Name:    to.Name,
			Address: to.Email,
		}
	}

	personalization.AddTos(tos...)

	personalization.DynamicTemplateData = templateData

	message.AddPersonalizations(&personalization)

	return message
}

// isAllowed indicates whether domain of the recipient email is allowed to receive emails,
// this is needed because we don't want to send emails to outside users in NOT PRODUCTION environments.
func (sg *sendGrid) isAllowed(email string) bool {
	if email == "" {
		return false
	}

	if strings.EqualFold(sg.env, config.ENVProduction) || len(sg.allowDomains) == 0 {
		return true
	}

	mailParts := strings.Split(email, "@")
	if len(mailParts) <= 1 {
		return false
	}

	domain := mailParts[1]

	for _, v := range sg.allowDomains {
		if v == domain {
			return true
		}
	}

	return false
}
