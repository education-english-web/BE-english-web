package event

//go:generate mockgen -destination=./mock/$GOFILE -source=$GOFILE -package=mock

const (
	// event name iota value
	EventNameInternalUserLogin Name = iota + 1
	EventNameInternalUserAdd
	EventNameInternalUserUpdateRole
	EventNameDeleteOffice
	EventNameForcefullyAddUser
	EventNameCreateOfficeUsage
	EventNameUpdateUserSetting
	EventNameFailConclusionJob
	EventNameFailMFCBoxJob
	EventNameProxyLogin
	EventNameTurnOffLegalCheck
)

// #nosec G101
const (
	// event name string value
	eventNameInternalUserLogin      = "aweb_login"
	eventNameInternalUserAdd        = "aweb_add_user"
	eventNameInternalUserUpdateRole = "aweb_update_user_role"
	eventNameDeleteOffice           = "aweb_delete_office"
	eventNameForcefullyAddUser      = "aweb_forcefully_add_user"
	eventNameCreateOfficeUsage      = "aweb_create_office_usage"
	eventNameUpdateUserSetting      = "aweb_update_user_setting"
	eventNameFailConclusionJob      = "aweb_fail_conclusion_job"
	eventNameFailMFCBoxJob          = "aweb_fail_mfc_box_job"
	eventNameProxyLogin             = "proxy_login"
	eventNameTurnOffLegalCheck      = "turn_off_legal_check"
)

// Event provides methods to handle a notification event
type Event interface {
	Name() Name
	Channels() []Channel
	BuildMessage(payload map[string]interface{}) string
	Env() string
}

// Name is the enum of event types
type Name int

//nolint:tagliatelle
type Channel struct {
	WebhookURL string `yaml:"webhook_url"`
	Name       string `yaml:"name"`
	Envs       string `yaml:"envs"`
}

// String return a string value of event name
func (e Name) String() string {
	switch e {
	case EventNameInternalUserLogin:
		return eventNameInternalUserLogin
	case EventNameInternalUserAdd:
		return eventNameInternalUserAdd
	case EventNameInternalUserUpdateRole:
		return eventNameInternalUserUpdateRole
	case EventNameDeleteOffice:
		return eventNameDeleteOffice
	case EventNameForcefullyAddUser:
		return eventNameForcefullyAddUser
	case EventNameCreateOfficeUsage:
		return eventNameCreateOfficeUsage
	case EventNameUpdateUserSetting:
		return eventNameUpdateUserSetting
	case EventNameFailConclusionJob:
		return eventNameFailConclusionJob
	case EventNameFailMFCBoxJob:
		return eventNameFailMFCBoxJob
	case EventNameProxyLogin:
		return eventNameProxyLogin
	case EventNameTurnOffLegalCheck:
		return eventNameTurnOffLegalCheck
	default:
		return ""
	}
}
