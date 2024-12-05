package errors

import "testing"

func TestGetErrorMessage(t *testing.T) {
	type args struct {
		code   Code
		fields []interface{}
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "invalid user email",
			args: args{
				code:   CodeBatchSendingCSVInvalidUserEmail,
				fields: []interface{}{"field1", "field2"},
			},
			want: "【field1】に「field2」が事業者に登録されているデータに該当しません。",
		},
		{
			name: "invalid date format",
			args: args{
				code:   CodeBatchSendingCSVInvalidDateFormat,
				fields: []interface{}{"field"},
			},
			want: "【field】の入力形式が正しくありません。YYYY/MM/DD の形式で入力してください。",
		},
		{
			name: "invalid number value",
			args: args{
				code:   CodeBatchSendingCSVInvalidNumberValue,
				fields: []interface{}{"field"},
			},
			want: "【field】の入力形式が正しくありません。数値を入力してください。",
		},
		{
			name: "invalid boolean value",
			args: args{
				code:   CodeBatchSendingCSVInvalidBooleanValue,
				fields: []interface{}{"field"},
			},
			want: "【field】に「あり」、「なし」 のいずれかを入力してください。",
		},
		{
			name: "invalid date value",
			args: args{
				code:   CodeBatchSendingCSVInvalidDateValue,
				fields: []interface{}{"field"},
			},
			want: "【field】の入力形式が正しくありません。YYYY/MM/DD の形式で入力してください。",
		},
		{
			name: "invalid rule oneof",
			args: args{
				code:   CodeBatchSendingCSVInvalidRuleOneOf,
				fields: []interface{}{"field", "「あり」、「なし」"},
			},
			want: "【field】に「あり」、「なし」 のいずれかを入力してください。",
		},
		{
			name: "invalid rule required",
			args: args{
				code:   CodeBatchSendingCSVInvalidRuleRequired,
				fields: []interface{}{"field"},
			},
			want: "必須項目の【field】を入力してください。",
		},
		{
			name: "invalid rule max",
			args: args{
				code:   CodeBatchSendingCSVInvalidRuleMax,
				fields: []interface{}{"field", 100},
			},
			want: "【field】を100文字以下で入力してください。",
		},
		{
			name: "invalid rule renew_duration",
			args: args{
				code:   CodeBatchSendingCSVInvalidRuleRenewDuration,
				fields: []interface{}{"field", 1, 12},
			},
			want: "【field】に1から12までの数値を入力してください。",
		},
		{
			name: "invalid rule cancel_auto_renew_duration",
			args: args{
				code:   CodeBatchSendingCSVInvalidRuleCancelAutoRenewDuration,
				fields: []interface{}{"field", 1, 99},
			},
			want: "【field】に1から99までの数値を入力してください。",
		},
		{
			name: "invalid rule contract_value",
			args: args{
				code:   CodeBatchSendingCSVInvalidRuleContractValue,
				fields: []interface{}{"field"},
			},
			want: "【field】に0から1000000000000までの数値を入力してください。",
		},
		{
			name: "invalid rule email format",
			args: args{
				code:   CodeBatchSendingCSVInvalidRuleEmailFormat,
				fields: []interface{}{"field", "format"},
			},
			want: "【field】に「format」の入力形式が正しくありません。入力内容を確認してください。",
		},
		{
			name: "invalid rule numeric",
			args: args{
				code:   CodeBatchSendingCSVInvalidRuleNumeric,
				fields: []interface{}{"field"},
			},
			want: "【field】の入力形式が正しくありません。数値を入力してください。",
		},
		{
			name: "end date before start date",
			args: args{
				code:   CodeBatchSendingCSVEndDateBeforeStartDate,
				fields: []interface{}{"end_date", "start_date"},
			},
			want: "【end_date】は【start_date】の後の日付を入力してください。",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetErrorMessage(tt.args.code, tt.args.fields...); got != tt.want {
				t.Errorf("GetErrorMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRuleErrorCode(t *testing.T) {
	type args struct {
		rule string
	}

	tests := []struct {
		name string
		args args
		want Code
	}{
		{
			name: "required",
			args: args{
				rule: "required",
			},
			want: CodeBatchSendingCSVInvalidRuleRequired,
		},
		{
			name: "max",
			args: args{
				rule: "max",
			},
			want: CodeBatchSendingCSVInvalidRuleMax,
		},
		{
			name: "oneof",
			args: args{
				rule: "oneof",
			},
			want: CodeBatchSendingCSVInvalidRuleOneOf,
		},
		{
			name: "renew_duration",
			args: args{
				rule: "renew_duration",
			},
			want: CodeBatchSendingCSVInvalidRuleRenewDuration,
		},
		{
			name: "cancel_auto_renew_duration",
			args: args{
				rule: "cancel_auto_renew_duration",
			},
			want: CodeBatchSendingCSVInvalidRuleCancelAutoRenewDuration,
		},
		{
			name: "contract_value",
			args: args{
				rule: "contract_value",
			},
			want: CodeBatchSendingCSVInvalidRuleContractValue,
		},
		{
			name: "email",
			args: args{
				rule: "email",
			},
			want: CodeBatchSendingCSVInvalidRuleEmailFormat,
		},
		{
			name: "numeric",
			args: args{
				rule: "numeric",
			},
			want: CodeBatchSendingCSVInvalidRuleNumeric,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRuleErrorCode(tt.args.rule); got != tt.want {
				t.Errorf("GetRuleErrorCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
