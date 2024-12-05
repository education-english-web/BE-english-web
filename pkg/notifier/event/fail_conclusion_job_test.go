package event

import (
	"reflect"
	"testing"
	"time"

	"github.com/education-english-web/BE-english-web/pkg/timeutil"
)

func TestNewEventFailConclusionJob(t *testing.T) {
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
				messageTemplate: "msg",
				channels:        []Channel{},
			},
			want: &eventFailConclusionJob{
				env:             "development",
				messageTemplate: "msg",
				channels:        []Channel{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventFailConclusionJob(
				tt.args.env,
				tt.args.messageTemplate,
				tt.args.channels,
			); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventFailConclusionJob \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventFailConclusionJob_Name(t *testing.T) {
	tests := []struct {
		name string
		want Name
	}{
		{
			name: "event_fail_conclusion_job",
			want: EventNameFailConclusionJob,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventFailConclusionJob{}
			if got := e.Name(); got != tt.want {
				t.Errorf("NewEventFailConclusionJob.Name \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventFailConclusionJob_Channels(t *testing.T) {
	type args struct {
		env             string
		messageTemplate string
		channels        []Channel
	}

	tests := []struct {
		name string
		args args
		want []Channel
	}{
		{
			name: "success",
			args: args{
				env:             "development",
				messageTemplate: "msg",
				channels: []Channel{
					{
						Envs:       "development",
						Name:       "channel",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			want: []Channel{
				{
					Envs:       "development",
					Name:       "channel",
					WebhookURL: "hooks.slack.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventFailConclusionJob{
				env:             tt.args.env,
				messageTemplate: tt.args.messageTemplate,
				channels:        tt.args.channels,
			}
			if got := e.Channels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventFailConclusionJob.Channels \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventFailConclusionJob_Env(t *testing.T) {
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
			e := &eventFailConclusionJob{env: tt.args.env}
			if got := e.Env(); got != tt.want {
				t.Errorf("eventFailConclusionJob.Env \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_eventFailConclusionJob_BuildMessage(t *testing.T) {
	type fields struct {
		env             string
		messageTemplate string
		channels        []Channel
	}

	confirmedAt := time.Date(2022, time.November, 23, 5, 6, 7, 8, time.UTC)

	type args struct {
		payload map[string]interface{}
	}

	tests := []struct {
		name   string
		args   args
		fields fields
		want   string
	}{
		{
			name: "success - Secom failure",
			fields: fields{
				env:             "development",
				messageTemplate: "<@U01EPR63S3Y> <@ULDPNF4UC> <@U04F7AVQU04> <@U0311AW9APL>\n%s\n%s/stml/failed-processes/\n\n%s\n事業者番号: %s\n書類番号: %s\n申請ユーザID: %s\n承認完了日時: %s\n",
				channels: []Channel{
					{
						Envs:       "development",
						Name:       "channel",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"fqdn":                       "localhost",
					"office_identification_code": "1",
					"contract_number":            "001",
					"applicant_mfid_uid":         "111",
					"error_type":                 ErrTypeSecomFailure,
					"confirmed_at":               confirmedAt.In(timeutil.JST),
				},
			},
			want: "<@U01EPR63S3Y> <@ULDPNF4UC> <@U04F7AVQU04> <@U0311AW9APL>\nWe got a failed process due to SECOM error, please check it.\nlocalhost/stml/failed-processes/\n\n下記事業者にて締結処理に失敗した書類が発生しています。\n事業者番号: 1\n書類番号: 001\n申請ユーザID: 111\n承認完了日時: 2022-11-23 14:06:07.000000008 +0900 Asia/Tokyo\n",
		},
		{
			name: "success - Sendgrid issue",
			fields: fields{
				env:             "development",
				messageTemplate: "<@U01EPR63S3Y> <@ULDPNF4UC> <@U04F7AVQU04> <@U0311AW9APL>\n%s\n%s/stml/failed-processes/\n\n%s\n事業者番号: %s\n書類番号: %s\n申請ユーザID: %s\n承認完了日時: %s\n",
				channels: []Channel{
					{
						Envs:       "development",
						Name:       "channel",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"fqdn":                       "localhost",
					"office_identification_code": "1",
					"contract_number":            "001",
					"applicant_mfid_uid":         "111",
					"error_type":                 ErrTypeSendGridIssue,
					"confirmed_at":               confirmedAt.In(timeutil.JST),
				},
			},
			want: "<@U01EPR63S3Y> <@ULDPNF4UC> <@U04F7AVQU04> <@U0311AW9APL>\nWe got a failed process due to SendGrid’s error, please check it.\nlocalhost/stml/failed-processes/\n\n下記事業者にて締結処理に失敗した書類が発生しています。\nこのプロセスは再処理によってリトライ可能である可能性があります。\n事業者番号: 1\n書類番号: 001\n申請ユーザID: 111\n承認完了日時: 2022-11-23 14:06:07.000000008 +0900 Asia/Tokyo\n",
		},
		{
			name: "success - Invalid PDF",
			fields: fields{
				env:             "development",
				messageTemplate: "<@U01EPR63S3Y> <@ULDPNF4UC> <@U04F7AVQU04> <@U0311AW9APL>\n%s\n%s/stml/failed-processes/\n\n%s\n事業者番号: %s\n書類番号: %s\n申請ユーザID: %s\n承認完了日時: %s\n",
				channels: []Channel{
					{
						Envs:       "development",
						Name:       "channel",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"fqdn":                       "localhost",
					"office_identification_code": "1",
					"contract_number":            "001",
					"applicant_mfid_uid":         "111",
					"error_type":                 ErrTypeInvalidPDFFormat,
					"confirmed_at":               confirmedAt.In(timeutil.JST),
				},
			},
			want: "<@U01EPR63S3Y> <@ULDPNF4UC> <@U04F7AVQU04> <@U0311AW9APL>\nWe got a failed process due to PDF format error, please check it.\nThis process would need to contact the user.\nlocalhost/stml/failed-processes/\n\n下記事業者にてPDFフォーマットが原因で締結処理に失敗した書類が発生しています。\nアップロードしたユーザーへの連絡を行ってください。\n事業者番号: 1\n書類番号: 001\n申請ユーザID: 111\n承認完了日時: 2022-11-23 14:06:07.000000008 +0900 Asia/Tokyo\n",
		},
		{
			name: "success - Others",
			fields: fields{
				env:             "development",
				messageTemplate: "<@U01EPR63S3Y> <@ULDPNF4UC> <@U04F7AVQU04> <@U0311AW9APL>\n%s\n%s/stml/failed-processes/\n\n%s\n事業者番号: %s\n書類番号: %s\n申請ユーザID: %s\n承認完了日時: %s\n",
				channels: []Channel{
					{
						Envs:       "development",
						Name:       "channel",
						WebhookURL: "hooks.slack.com",
					},
				},
			},
			args: args{
				payload: map[string]interface{}{
					"fqdn":                       "localhost",
					"office_identification_code": "1",
					"contract_number":            "001",
					"applicant_mfid_uid":         "111",
					"error_type":                 ErrTypeOthers,
					"confirmed_at":               confirmedAt.In(timeutil.JST),
				},
			},
			want: "<@U01EPR63S3Y> <@ULDPNF4UC> <@U04F7AVQU04> <@U0311AW9APL>\nWe got a failed process due to some error, please check it.\nThis process is likely to be retried.\nlocalhost/stml/failed-processes/\n\n下記事業者にて何らかのエラーで締結処理に失敗した書類が発生しています。\nこのプロセスは再処理によって解決可能である可能性があります。再処理して、解決しない場合開発側と原因調査を行なってください\n事業者番号: 1\n書類番号: 001\n申請ユーザID: 111\n承認完了日時: 2022-11-23 14:06:07.000000008 +0900 Asia/Tokyo\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventFailConclusionJob{
				env:             tt.fields.env,
				messageTemplate: tt.fields.messageTemplate,
				channels:        tt.fields.channels,
			}
			if got := e.BuildMessage(tt.args.payload); got != tt.want {
				t.Errorf("NewEventFailConclusionJob.BuildMessage \ngot: %v \nwant: %v", got, tt.want)
			}
		})
	}
}
