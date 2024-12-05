package event

import (
	"fmt"
)

type ConclusionJobErrType string

var (
	ErrTypeSecomFailure     ConclusionJobErrType = "secom_failure"
	ErrTypeSendGridIssue    ConclusionJobErrType = "sendGrid_issue"
	ErrTypeInvalidPDFFormat ConclusionJobErrType = "invalid_pdf_format"
	ErrTypeOthers           ConclusionJobErrType = "others"
)

const (
	ErrMsgFmtEN       = "We got a failed process due to %s error, please check it.%s"
	ErrMsgSecomEN     = "SECOM"
	ErrMsgSendGridEN  = "SendGrid’s"
	ErrMsgPDFFormatEN = "PDF format"
	ErrMsgOthersEN    = "some"
	ErrMsgSecomJP     = "下記事業者にて締結処理に失敗した書類が発生しています。"
	ErrMsgSendGridJP  = "下記事業者にて締結処理に失敗した書類が発生しています。\nこのプロセスは再処理によってリトライ可能である可能性があります。"
	ErrMsgPDFFormatJP = "下記事業者にてPDFフォーマットが原因で締結処理に失敗した書類が発生しています。\nアップロードしたユーザーへの連絡を行ってください。"
	ErrMsgOthersJP    = "下記事業者にて何らかのエラーで締結処理に失敗した書類が発生しています。\nこのプロセスは再処理によって解決可能である可能性があります。再処理して、解決しない場合開発側と原因調査を行なってください"
)

// eventFailConclusionJob holds information for the event when the conclusion process failed
type eventFailConclusionJob struct {
	env             string
	messageTemplate string
	channels        []Channel
}

func NewEventFailConclusionJob(env, messageTemplate string, channels []Channel) Event {
	return &eventFailConclusionJob{
		env:             env,
		messageTemplate: messageTemplate,
		channels:        channels,
	}
}

func (e *eventFailConclusionJob) Name() Name {
	return EventNameFailConclusionJob
}

func (e *eventFailConclusionJob) Channels() []Channel {
	return e.channels
}

func (e *eventFailConclusionJob) Env() string {
	return e.env
}

func (e *eventFailConclusionJob) BuildMessage(payload map[string]interface{}) string {
	switch payload["error_type"] {
	case ErrTypeSecomFailure:
		return e.prepareMessage(
			payload,
			fmt.Sprintf(ErrMsgFmtEN, ErrMsgSecomEN, ""),
			ErrMsgSecomJP,
		)
	case ErrTypeSendGridIssue:
		return e.prepareMessage(
			payload,
			fmt.Sprintf(ErrMsgFmtEN, ErrMsgSendGridEN, ""),
			ErrMsgSendGridJP,
		)
	case ErrTypeInvalidPDFFormat:
		return e.prepareMessage(
			payload,
			fmt.Sprintf(ErrMsgFmtEN, ErrMsgPDFFormatEN, "\nThis process would need to contact the user."),
			ErrMsgPDFFormatJP,
		)
	default:
		return e.prepareMessage(
			payload,
			fmt.Sprintf(ErrMsgFmtEN, ErrMsgOthersEN, "\nThis process is likely to be retried."),
			ErrMsgOthersJP,
		)
	}
}

func (e *eventFailConclusionJob) prepareMessage(
	payload map[string]interface{},
	errMsgEN string,
	errMsgJP string,
) string {
	return fmt.Sprintf(
		e.messageTemplate,
		errMsgEN,
		payload["fqdn"],
		errMsgJP,
		payload["office_identification_code"],
		payload["contract_number"],
		payload["applicant_mfid_uid"],
		payload["confirmed_at"],
	)
}
