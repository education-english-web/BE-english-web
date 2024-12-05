package event

import "testing"

func TestEventName_String(t *testing.T) {
	tests := []struct {
		name string
		e    Name
		want string
	}{
		{
			name: "unknown",
			e:    0,
			want: "",
		},
		{
			name: "aweb_login",
			e:    EventNameInternalUserLogin,
			want: "aweb_login",
		},
		{
			name: "aweb_add_user",
			e:    EventNameInternalUserAdd,
			want: "aweb_add_user",
		},
		{
			name: "aweb_update_user_role",
			e:    EventNameInternalUserUpdateRole,
			want: "aweb_update_user_role",
		},
		{
			name: "aweb_delete_office",
			e:    EventNameDeleteOffice,
			want: "aweb_delete_office",
		},
		{
			name: "aweb_forcefully_add_user",
			e:    EventNameForcefullyAddUser,
			want: "aweb_forcefully_add_user",
		},
		{
			name: "aweb_create_office_usage",
			e:    EventNameCreateOfficeUsage,
			want: "aweb_create_office_usage",
		},
		{
			name: "aweb_update_user_setting",
			e:    EventNameUpdateUserSetting,
			want: "aweb_update_user_setting",
		},
		{
			name: "aweb_fail_conclusion_job",
			e:    EventNameFailConclusionJob,
			want: "aweb_fail_conclusion_job",
		},
		{
			name: "aweb_fail_mfc_box_job",
			e:    EventNameFailMFCBoxJob,
			want: "aweb_fail_mfc_box_job",
		},
		{
			name: "proxy_login",
			e:    EventNameProxyLogin,
			want: "proxy_login",
		},
		{
			name: "turn_off_legal_check",
			e:    EventNameTurnOffLegalCheck,
			want: "turn_off_legal_check",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("EventName.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
