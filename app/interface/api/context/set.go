package context

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func setContextKey(ctx *gin.Context, k key, value interface{}) {
	ctx.Set(k.String(), value)
}

func SetClientIP(ctx *gin.Context, value string) {
	childCtx := context.WithValue(ctx.Request.Context(), keyClientIP, value)
	*ctx.Request = *ctx.Request.WithContext(childCtx)
}

func SetUserAgent(ctx *gin.Context, value string) {
	childCtx := context.WithValue(ctx.Request.Context(), keyUserAgent, value)
	*ctx.Request = *ctx.Request.WithContext(childCtx)
}

// SetMFIDUserID sets mfid_user_id into context
func SetMFIDUserID(ctx *gin.Context, value uint32) {
	setContextKey(ctx, keyMFIDUserID, value)
}

// SetNavisOfficeID set navs_office_id into context
func SetNavisOfficeID(ctx *gin.Context, value uint32) {
	setContextKey(ctx, keyNavisOfficeID, value)
}

// SetTenantUID set tenant_uid into context
func SetTenantUID(ctx *gin.Context, value uint32) {
	setContextKey(ctx, keyTenantUID, value)
}

// SetPartnerCompanyID set partner company id into context
func SetPartnerCompanyID(ctx *gin.Context, value uint32) {
	setContextKey(ctx, keyPartnerCompanyID, value)
}

// SetPartnerEmail set partner_email into context
func SetPartnerEmail(ctx *gin.Context, value string) {
	setContextKey(ctx, keyPartnerEmail, value)
}

// SetPartnerAccessKey set partner access key into context
func SetPartnerAccessKey(ctx *gin.Context, value string) {
	setContextKey(ctx, keyPartnerAccessKey, value)
}

// SetUserID set user_id into context
func SetUserID(ctx *gin.Context, value uuid.UUID) {
	setContextKey(ctx, keyUserID, value)
}

// SetUserEmail set internal_user_email into context
func SetUserEmail(ctx *gin.Context, value string) {
	setContextKey(ctx, keyUserEmail, value)
}

// SetPhoneNumber set phone_number into context
func SetPhoneNumber(ctx *gin.Context, value string) {
	setContextKey(ctx, keyPhoneNumber, value)
}

// SetSalesforceOrgID set salesforce_org_id into context
func SetSalesforceOrgID(ctx *gin.Context, value string) {
	setContextKey(ctx, keySalesforceOrgID, value)
}

// SetSalesforceRequestedEmail set salesforce_requested_email into context
func SetSalesforceRequestedEmail(ctx *gin.Context, value string) {
	setContextKey(ctx, keySalesforceRequestedEmail, value)
}

//// SetOfficePlan set office_plan into context
//func SetOfficePlan(ctx *gin.Context, value entity.OfficePlan) {
//	setContextKey(ctx, keyOfficePlan, value)
//}

//// SetOfficePlanOrigin set office_plan into context
//func SetOfficePlanOrigin(ctx *gin.Context, value entity.OfficePlanOrigin) {
//	setContextKey(ctx, keyOfficePlanOrigin, value)
//}

// SetOfficeUsageEndDate set office_usage_end_date into context
func SetOfficeUsageEndDate(ctx *gin.Context, value time.Time) {
	setContextKey(ctx, keyOfficeUsageEndDate, value.Format(time.RFC3339))
}

// SetOfficeUsageStartDate set office_usage_start_date into context
func SetOfficeUsageStartDate(ctx *gin.Context, value time.Time) {
	setContextKey(ctx, keyOfficeUsageStartDate, value.Format(time.RFC3339))
}

// SetOfficeUsageSalesforceLinkage set office_usage_salesforce_linkage into context
func SetOfficeUsageSalesforceLinkage(ctx *gin.Context, value bool) {
	setContextKey(ctx, keyOfficeUsageSalesforceLinkage, value)
}

// SetOfficeUsageRelevantDocuments set office_usage_relevant_documents into context
func SetOfficeUsageRelevantDocuments(ctx *gin.Context, value bool) {
	setContextKey(ctx, keyOfficeUsageRelevantDocuments, value)
}

//// SetUserRoleCode set user_role_code into context
//func SetUserRoleCode(ctx *gin.Context, value entity.RoleCode) {
//	setContextKey(ctx, keyUserRoleCode, value)
//}

//// SetUserLinkages set user_linkages into context
//func SetUserLinkages(ctx *gin.Context, values []entity.UserLinkage) {
//	setContextKey(ctx, keyUserLinkages, values)
//}

func SetProxyLoginEventID(ctx *gin.Context, value uint32) {
	setContextKey(ctx, keyProxyLoginEventID, value)
}

func SetContractID(ctx *gin.Context, value uint32) {
	setContextKey(ctx, keyContractID, value)
}
