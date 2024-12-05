package mailer

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

// Recipient holds information of recipient to send
type Recipient struct {
	Name  string
	Email string
}

// EmailData for template
type EmailData map[string]interface{}

// Attachment information of attachment file that follow an email
type Attachment struct {
	Name  string
	Type  string
	Bytes []byte
}

// Mailer interface represents what an email service could do
type Mailer interface {
	// SendWithTemplate sends email using pre-defined sendgrid's template
	SendWithTemplate(templateID string, recipients []Recipient, templateData EmailData, attachments ...Attachment) error
}
