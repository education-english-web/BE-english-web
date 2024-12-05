package errors

import (
	"fmt"
)

const (
	CodeBatchSendingCSVMismatchHeader             Code = "CODE_BATCH_SENDING_CSV_MISMATCH_HEADER"
	CodeBatchSendingCSVInvalidUserEmail           Code = "CODE_BATCH_SENDING_CSV_INVALID_USER_EMAIL"
	CodeBatchSendingCSVInvalidDateFormat          Code = "CODE_BATCH_SENDING_CSV_INVALID_DATE_FORMAT"
	CodeBatchSendingCSVInvalidNumberValue         Code = "CODE_BATCH_SENDING_CSV_INVALID_NUMBER_VALUE"
	CodeBatchSendingCSVInvalidBooleanValue        Code = "CODE_BATCH_SENDING_CSV_INVALID_BOOLEAN_VALUE"
	CodeBatchSendingCSVInvalidDateValue           Code = "CODE_BATCH_SENDING_CSV_INVALID_DATETIME_VALUE"
	CodeBatchSendingCSVEndDateBeforeStartDate     Code = "CODE_BATCH_SENDING_CSV_END_DATE_BEFORE_START_DATE"
	CodeBatchSendingCSVInvalidContractValueFormat Code = "CODE_BATCH_SENDING_CSV_INVALID_CONTRACT_VALUE_FORMAT"
)

// Rule error
const (
	CodeBatchSendingCSVInvalidRuleOneOf                   Code = "CODE_BATCH_SENDING_CSV_INVALID_RULE_ONE_OF"
	CodeBatchSendingCSVInvalidRuleRequired                Code = "CODE_BATCH_SENDING_CSV_INVALID_RULE_REQUIRED"
	CodeBatchSendingCSVInvalidRuleMax                     Code = "CODE_BATCH_SENDING_CSV_INVALID_RULE_MAX"
	CodeBatchSendingCSVInvalidRuleRenewDuration           Code = "CODE_BATCH_SENDING_CSV_INVALID_RULE_RENEW_DURATION"
	CodeBatchSendingCSVInvalidRuleCancelAutoRenewDuration Code = "CODE_BATCH_SENDING_CSV_INVALID_RULE_CANCEL_AUTO_RENEW_DURATION"
	CodeBatchSendingCSVInvalidRuleContractValue           Code = "CODE_BATCH_SENDING_CSV_INVALID_RULE_CONTRACT_VALUE"
	CodeBatchSendingCSVInvalidRuleEmailFormat             Code = "CODE_BATCH_SENDING_CSV_INVALID_RULE_EMAIL_FORMAT"
	CodeBatchSendingCSVInvalidRuleNumeric                 Code = "CODE_BATCH_SENDING_CSV_INVALID_RULE_NUMERIC"
)

var mRuleErrorToCode = map[string]Code{
	"required":                   CodeBatchSendingCSVInvalidRuleRequired,
	"max":                        CodeBatchSendingCSVInvalidRuleMax,
	"oneof":                      CodeBatchSendingCSVInvalidRuleOneOf,
	"renew_duration":             CodeBatchSendingCSVInvalidRuleRenewDuration,
	"cancel_auto_renew_duration": CodeBatchSendingCSVInvalidRuleCancelAutoRenewDuration,
	"contract_value":             CodeBatchSendingCSVInvalidRuleContractValue,
	"email":                      CodeBatchSendingCSVInvalidRuleEmailFormat,
	"numeric":                    CodeBatchSendingCSVInvalidRuleNumeric,
}

var mCodeToMessage = map[Code]string{
	CodeBatchSendingCSVInvalidUserEmail:                   "【%s】に「%v」が事業者に登録されているデータに該当しません。",
	CodeBatchSendingCSVInvalidDateFormat:                  "【%s】の入力形式が正しくありません。YYYY/MM/DD の形式で入力してください。",
	CodeBatchSendingCSVInvalidNumberValue:                 "【%s】の入力形式が正しくありません。数値を入力してください。",
	CodeBatchSendingCSVInvalidBooleanValue:                "【%s】に「あり」、「なし」 のいずれかを入力してください。",
	CodeBatchSendingCSVInvalidDateValue:                   "【%s】の入力形式が正しくありません。YYYY/MM/DD の形式で入力してください。",
	CodeBatchSendingCSVInvalidRuleOneOf:                   "【%s】に%v のいずれかを入力してください。",
	CodeBatchSendingCSVInvalidRuleRequired:                "必須項目の【%s】を入力してください。",
	CodeBatchSendingCSVInvalidRuleMax:                     "【%s】を%v文字以下で入力してください。",
	CodeBatchSendingCSVInvalidRuleRenewDuration:           "【%s】に%vから%vまでの数値を入力してください。",
	CodeBatchSendingCSVInvalidRuleCancelAutoRenewDuration: "【%s】に%vから%vまでの数値を入力してください。",
	CodeBatchSendingCSVInvalidRuleContractValue:           "【%s】に0から1000000000000までの数値を入力してください。",
	CodeBatchSendingCSVInvalidRuleEmailFormat:             "【%s】に「%v」の入力形式が正しくありません。入力内容を確認してください。",
	CodeBatchSendingCSVInvalidRuleNumeric:                 "【%s】の入力形式が正しくありません。数値を入力してください。",
	CodeBatchSendingCSVEndDateBeforeStartDate:             "【%s】は【%s】の後の日付を入力してください。",
	CodeBatchSendingCSVInvalidContractValueFormat:         "【%s】の入力形式が正しくありません。数値を入力してください。",
}

func GetRuleErrorCode(rule string) Code {
	return mRuleErrorToCode[rule]
}

func GetErrorMessage(code Code, fields ...interface{}) string {
	return fmt.Sprintf(mCodeToMessage[code], fields...)
}
