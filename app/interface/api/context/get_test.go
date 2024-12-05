package context

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
)

func TestGetMFIDUserID(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         uint32
	}{
		{
			name: "no mfid_user_id in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
			},
			want: uint32(0),
		},
		{
			name: "mfid_user_id value type not uint32",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyMFIDUserID.String(), "1")
			},
			want: uint32(0),
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetMFIDUserID(ctx, uint32(1))
			},
			want: uint32(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetMFIDUserID(tt.args.ctx))
		})
	}
}

func TestGetUserID(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         uint32
	}{
		{
			name: "no user_id in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
			},
			want: uint32(0),
		},
		{
			name: "user_id value type not uint32",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyUserID.String(), "1")
			},
			want: uint32(0),
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetUserID(ctx, uint32(1))
			},
			want: uint32(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetUserID(tt.args.ctx))
		})
	}
}

func TestGetNavisOfficeID(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         uint32
	}{
		{
			name: "no office_id in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
			},
			want: uint32(0),
		},
		{
			name: "office_id value type not uint32",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyNavisOfficeID.String(), "1")
			},
			want: uint32(0),
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetNavisOfficeID(ctx, uint32(1))
			},
			want: uint32(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetNavisOfficeID(tt.args.ctx))
		})
	}
}

func TestGetTenantUID(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         uint32
	}{
		{
			name: "no tenant_uid in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
			},
			want: uint32(0),
		},
		{
			name: "tenant_uid value type not uint32",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyTenantUID.String(), "1")
			},
			want: uint32(0),
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetTenantUID(ctx, uint32(1))
			},
			want: uint32(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetTenantUID(tt.args.ctx))
		})
	}
}

func TestGetPartnerEmail(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         string
	}{
		{
			name: "no partner_email in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         "",
		},
		{
			name: "partner_email value type not string",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyPartnerEmail.String(), 1)
			},
			want: "",
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetPartnerEmail(ctx, "email@example.com")
			},
			want: "email@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetPartnerEmail(tt.args.ctx))
		})
	}
}

func TestGetPartnerAccessKey(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         string
	}{
		{
			name: "no partner access key in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         "",
		},
		{
			name: "partner_access_key value type not string",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyPartnerAccessKey.String(), 1)
			},
			want: "",
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetPartnerAccessKey(ctx, "key")
			},
			want: "key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetPartnerAccessKey(tt.args.ctx))
		})
	}
}

func TestGetPartnerCompanyID(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         uint32
	}{
		{
			name: "no partner company id in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         0,
		},
		{
			name: "partner_company_id value type not uint32",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyPartnerCompanyID.String(), "1")
			},
			want: uint32(0),
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetPartnerCompanyID(ctx, 1)
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetPartnerCompanyID(tt.args.ctx))
		})
	}
}

func TestGetInternalUserEmail(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         string
	}{
		{
			name: "no internal_user_email in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         "",
		},
		{
			name: "internal_user_email value type not string",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyUserEmail.String(), 1)
			},
			want: "",
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetUserEmail(ctx, "email@example.com")
			},
			want: "email@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetInternalUserEmail(tt.args.ctx))
		})
	}
}

func TestGetSalesforceOrgID(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         string
	}{
		{
			name: "no salesforce_org_id in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         "",
		},
		{
			name: "salesforce_org_id value type not string",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keySalesforceOrgID.String(), 1)
			},
			want: "",
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetSalesforceOrgID(ctx, "XYZ")
			},
			want: "XYZ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetSalesforceOrgID(tt.args.ctx))
		})
	}
}

func TestGetSalesforceRequestedEmail(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         string
	}{
		{
			name: "no salesforce_requested_email in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         "",
		},
		{
			name: "salesforce_requested_email value type not string",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keySalesforceRequestedEmail.String(), 1)
			},
			want: "",
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetSalesforceRequestedEmail(ctx, "email@example.com")
			},
			want: "email@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetSalesforceRequestedEmail(tt.args.ctx))
		})
	}
}

func TestGetOfficePlan(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         entity.OfficePlan
	}{
		{
			name: "no office_plan in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         entity.OfficePlanFree,
		},
		{
			name: "office_plan value type not entity.OfficePlan",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyOfficePlan.String(), 1)
			},
			want: entity.OfficePlanFree,
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetOfficePlan(ctx, entity.OfficePlanTrial)
			},
			want: entity.OfficePlanTrial,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetOfficePlan(tt.args.ctx))
		})
	}
}

func TestGetOfficePlanOrigin(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         entity.OfficePlanOrigin
	}{
		{
			name: "no office_plan in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         entity.OfficePlanOriginAweb,
		},
		{
			name: "office_plan_origin value type not entity.OfficePlanOrigin",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyOfficePlanOrigin.String(), 1)
			},
			want: entity.OfficePlanOriginAweb,
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetOfficePlanOrigin(ctx, entity.OfficePlanOriginSinglePlanContract)
			},
			want: entity.OfficePlanOriginSinglePlanContract,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetOfficePlanOrigin(tt.args.ctx))
		})
	}
}

func TestGetOfficeUsageEndDate(t *testing.T) {
	type args struct {
		ctx *gin.Context
		d   time.Time
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context, d time.Time)
		want         *time.Time
	}{
		{
			name: "no office_usage_end_date in context",
			args: args{
				ctx: &gin.Context{},
				d:   time.Now(),
			},
			setupContext: func(ctx *gin.Context, d time.Time) {},
			want:         nil,
		},
		{
			name: "office_usage_end_date in context in not correct format time.RFC3339",
			args: args{
				ctx: &gin.Context{},
				d:   time.Now(),
			},
			setupContext: func(ctx *gin.Context, d time.Time) {
				setContextKey(ctx, keyOfficeUsageEndDate, d.Format(time.RFC1123))
			},
			want: nil,
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
				d:   time.Date(2021, time.December, 25, 1, 2, 3, 0, time.UTC),
			},
			setupContext: func(ctx *gin.Context, d time.Time) {
				SetOfficeUsageEndDate(ctx, d)
			},
			want: func() *time.Time {
				d := time.Date(2021, time.December, 25, 1, 2, 3, 0, time.UTC)

				return &d
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx, tt.args.d)
			got := GetOfficeUsageEndDate(tt.args.ctx)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetOfficeUsageStartDate(t *testing.T) {
	type args struct {
		ctx *gin.Context
		d   time.Time
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context, d time.Time)
		want         *time.Time
	}{
		{
			name: "no office_usage_start_date in context",
			args: args{
				ctx: &gin.Context{},
				d:   time.Now(),
			},
			setupContext: func(ctx *gin.Context, d time.Time) {},
			want:         nil,
		},
		{
			name: "office_usage_start_date in context in not correct format time.RFC3339",
			args: args{
				ctx: &gin.Context{},
				d:   time.Now(),
			},
			setupContext: func(ctx *gin.Context, d time.Time) {
				setContextKey(ctx, keyOfficeUsageStartDate, d.Format(time.RFC1123))
			},
			want: nil,
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
				d:   time.Date(2021, time.December, 25, 1, 2, 3, 0, time.UTC),
			},
			setupContext: func(ctx *gin.Context, d time.Time) {
				SetOfficeUsageStartDate(ctx, d)
			},
			want: func() *time.Time {
				d := time.Date(2021, time.December, 25, 1, 2, 3, 0, time.UTC)

				return &d
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx, tt.args.d)
			got := GetOfficeUsageStartDate(tt.args.ctx)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetOfficeUsageSalesforceLinkage(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         bool
	}{
		{
			name: "no office_usage_salesforce_linkage in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         false,
		},
		{
			name: "office_usage_salesforce_linkage not bool",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyOfficeUsageSalesforceLinkage.String(), 1)
			},
			want: false,
		},
		{
			name: "success - false",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetOfficeUsageSalesforceLinkage(ctx, false)
			},
			want: false,
		},
		{
			name: "success - true",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetOfficeUsageSalesforceLinkage(ctx, true)
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetOfficeUsageSalesforceLinkage(tt.args.ctx))
		})
	}
}

func TestGetOfficeUsageRelevantDocuments(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         bool
	}{
		{
			name: "no office_usage_relevant_documents in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         false,
		},
		{
			name: "office_usage_relevant_documents not bool",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyOfficeUsageRelevantDocuments.String(), 1)
			},
			want: false,
		},
		{
			name: "success - false",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetOfficeUsageRelevantDocuments(ctx, false)
			},
			want: false,
		},
		{
			name: "success - true",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetOfficeUsageRelevantDocuments(ctx, true)
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetOfficeUsageRelevantDocuments(tt.args.ctx))
		})
	}
}

func TestGetUserRoleCode(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         entity.RoleCode
	}{
		{
			name: "no user_role_code in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         entity.RoleCodeUnknown,
		},
		{
			name: "user_role_code not entity.RoleCode",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyUserRoleCode.String(), 1)
			},
			want: entity.RoleCodeUnknown,
		},
		{
			name: "success - admin",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetUserRoleCode(ctx, entity.RoleCodeAdmin)
			},
			want: entity.RoleCodeAdmin,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetUserRoleCode(tt.args.ctx))
		})
	}
}

func TestGetUserLinkages(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         []entity.UserLinkage
	}{
		{
			name: "no user_linkages in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {},
			want:         nil,
		},
		{
			name: "user_linkages not []entity.UserLinkage",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyUserLinkages.String(), 1)
			},
			want: nil,
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetUserLinkages(ctx, []entity.UserLinkage{
					{
						LinkageCode: entity.LinkageCodeFlair,
					},
					{
						LinkageCode: entity.LinkageCodeSalesforce,
					},
				})
			},
			want: []entity.UserLinkage{
				{
					LinkageCode: entity.LinkageCodeFlair,
				},
				{
					LinkageCode: entity.LinkageCodeSalesforce,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)
			got := GetUserLinkages(tt.args.ctx)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetProxyLoginEventID(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         uint32
	}{
		{
			name: "no proxy_login_event_id in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
			},
			want: uint32(0),
		},
		{
			name: "proxy_login_event_id value type not uint32",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyProxyLoginEventID.String(), "1")
			},
			want: uint32(0),
		},

		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetProxyLoginEventID(ctx, 1)
			},
			want: uint32(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetProxyLoginEventID(tt.args.ctx))
		})
	}
}

func TestGetContractID(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name         string
		args         args
		setupContext func(ctx *gin.Context)
		want         uint32
	}{
		{
			name: "no contract_id in context",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
			},
			want: uint32(0),
		},
		{
			name: "contract_id value type not uint32",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				ctx.Set(keyContractID.String(), "1")
			},
			want: uint32(0),
		},
		{
			name: "success",
			args: args{
				ctx: &gin.Context{},
			},
			setupContext: func(ctx *gin.Context) {
				SetContractID(ctx, 1)
			},
			want: uint32(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupContext(tt.args.ctx)

			assert.Equal(t, tt.want, GetContractID(tt.args.ctx))
		})
	}
}
