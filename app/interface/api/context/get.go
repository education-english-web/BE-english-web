package context

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetClientIP(ctx context.Context) string {
	if clientIP, ok := ctx.Value(keyClientIP).(string); ok {
		return clientIP
	}

	return ""
}

func GetUserAgent(ctx context.Context) string {
	if userAgent, ok := ctx.Value(keyUserAgent).(string); ok {
		return userAgent
	}

	return ""
}

// GetMFIDUserID extract mfid_user_id from Context
func GetMFIDUserID(ctx *gin.Context) uint32 {
	if sMFIDUserID, ok := ctx.Get(keyMFIDUserID.String()); ok {
		if mfidUserID, ok := sMFIDUserID.(uint32); ok {
			return mfidUserID
		}
	}

	return 0
}

// GetUserID extract user_id from Context
func GetUserID(ctx *gin.Context) uuid.UUID {
	if uID, ok := ctx.Get(keyUserID.String()); ok {
		if userID, ok := uID.(uuid.UUID); ok {
			return userID
		}
	}

	return uuid.Nil
}

// GetNavisOfficeID extract office_id from Context
func GetNavisOfficeID(ctx *gin.Context) uint32 {
	if oID, ok := ctx.Get(keyNavisOfficeID.String()); ok {
		if officeID, ok := oID.(uint32); ok {
			return officeID
		}
	}

	return 0
}

// GetTenantUID extract tenant_uid from Context
func GetTenantUID(ctx *gin.Context) uint32 {
	if tID, ok := ctx.Get(keyTenantUID.String()); ok {
		if tenantUID, ok := tID.(uint32); ok {
			return tenantUID
		}
	}

	return 0
}

// GetPartnerEmail extract partner_email from Context
func GetPartnerEmail(ctx *gin.Context) string {
	if iEmail, ok := ctx.Get(keyPartnerEmail.String()); ok {
		if email, ok := iEmail.(string); ok {
			return email
		}
	}

	return ""
}

// GetPartnerAccessKey extract partner_access_key from context
func GetPartnerAccessKey(ctx *gin.Context) string {
	if iPartnerAccessKey, ok := ctx.Get(keyPartnerAccessKey.String()); ok {
		if partnerAccessKey, ok := iPartnerAccessKey.(string); ok {
			return partnerAccessKey
		}
	}

	return ""
}

// GetPartnerCompanyID extract partner_company_id from Context
func GetPartnerCompanyID(ctx *gin.Context) uint32 {
	if iPartnerCompanyID, ok := ctx.Get(keyPartnerCompanyID.String()); ok {
		if partnerCompanyID, ok := iPartnerCompanyID.(uint32); ok {
			return partnerCompanyID
		}
	}

	return 0
}

// GetUserEmail extracts user_email from Context
func GetUserEmail(ctx *gin.Context) string {
	if iEmail, ok := ctx.Get(keyUserEmail.String()); ok {
		if email, ok := iEmail.(string); ok {
			return email
		}
	}

	return ""
}

// GetPhoneNumber extracts phone_number from context
func GetPhoneNumber(ctx *gin.Context) string {
	if iPhoneNumber, ok := ctx.Get(keyPhoneNumber.String()); ok {
		if phoneNumber, ok := iPhoneNumber.(string); ok {
			return phoneNumber
		}
	}

	return ""
}

// GetSalesforceOrgID extracts salesforce_org_id from context
func GetSalesforceOrgID(ctx *gin.Context) string {
	if iOrgID, ok := ctx.Get(keySalesforceOrgID.String()); ok {
		if orgID, ok := iOrgID.(string); ok {
			return orgID
		}
	}

	return ""
}

// GetSalesforceRequestedEmail returns salesforce_requested_email from context
func GetSalesforceRequestedEmail(ctx *gin.Context) string {
	if iRequestedEmail, ok := ctx.Get(keySalesforceRequestedEmail.String()); ok {
		if requestedEmail, ok := iRequestedEmail.(string); ok {
			return requestedEmail
		}
	}

	return ""
}

//// GetOfficePlan returns office_plan from context
//func GetOfficePlan(ctx *gin.Context) entity.OfficePlan {
//	if iOP, iOPOK := ctx.Get(keyOfficePlan.String()); iOPOK {
//		if op, opOK := iOP.(entity.OfficePlan); opOK {
//			return op
//		}
//	}
//
//	return entity.OfficePlanFree
//}

//// GetOfficePlanOrigin returns office_plan_origin from context
//func GetOfficePlanOrigin(ctx *gin.Context) entity.OfficePlanOrigin {
//	if iOPOrigin, ok := ctx.Get(keyOfficePlanOrigin.String()); ok {
//		if opOrigin, ok := iOPOrigin.(entity.OfficePlanOrigin); ok {
//			return opOrigin
//		}
//	}
//
//	return entity.OfficePlanOriginAweb
//}

// GetOfficeUsageEndDate returns office_usage_end_date from context
func GetOfficeUsageEndDate(ctx *gin.Context) *time.Time {
	iUsageEndDate, ok := ctx.Get(keyOfficeUsageEndDate.String())
	if !ok {
		return nil
	}

	endDate, err := time.Parse(time.RFC3339, fmt.Sprintf("%v", iUsageEndDate))
	if err != nil {
		return nil
	}

	return &endDate
}

// GetOfficeUsageStartDate returns office_usage_start_date from context
func GetOfficeUsageStartDate(ctx *gin.Context) *time.Time {
	iUsageStartDate, ok := ctx.Get(keyOfficeUsageStartDate.String())
	if !ok {
		return nil
	}

	startDate, err := time.Parse(time.RFC3339, fmt.Sprintf("%v", iUsageStartDate))
	if err != nil {
		return nil
	}

	return &startDate
}

// GetOfficeUsageSalesforceLinkage returns office_usage_salesforce_linkage from context
func GetOfficeUsageSalesforceLinkage(ctx *gin.Context) bool {
	if iSalesforceLinkage, ok := ctx.Get(keyOfficeUsageSalesforceLinkage.String()); ok {
		if salesforceLinkage, ok := iSalesforceLinkage.(bool); ok {
			return salesforceLinkage
		}
	}

	return false
}

// GetOfficeUsageRelevantDocuments returns office_usage_relevant_documents from context
func GetOfficeUsageRelevantDocuments(ctx *gin.Context) bool {
	if iRelevantDocuments, ok := ctx.Get(keyOfficeUsageRelevantDocuments.String()); ok {
		if relevantDocuments, ok := iRelevantDocuments.(bool); ok {
			return relevantDocuments
		}
	}

	return false
}

//// GetUserRoleCode returns user_role_code from context
//func GetUserRoleCode(ctx *gin.Context) entity.RoleCode {
//	if iRoleCode, ok := ctx.Get(keyUserRoleCode.String()); ok {
//		if roleCode, ok := iRoleCode.(entity.RoleCode); ok {
//			return roleCode
//		}
//	}
//
//	return entity.RoleCodeUnknown
//}

//// GetUserLinkages returns user_linkages from context
//func GetUserLinkages(ctx *gin.Context) []entity.UserLinkage {
//	if iUserLinkages, ok := ctx.Get(keyUserLinkages.String()); ok {
//		if userLinkages, ok := iUserLinkages.([]entity.UserLinkage); ok {
//			return userLinkages
//		}
//	}
//
//	return nil
//}

func GetProxyLoginEventID(ctx *gin.Context) uint32 {
	if iProxyLoginEventID, ok := ctx.Get(keyProxyLoginEventID.String()); ok {
		if proxyLoginEventID, ok := iProxyLoginEventID.(uint32); ok {
			return proxyLoginEventID
		}
	}

	return 0
}

func GetContractID(ctx *gin.Context) uint32 {
	if iContractID, ok := ctx.Get(keyContractID.String()); ok {
		if contractID, ok := iContractID.(uint32); ok {
			return contractID
		}
	}

	return 0
}
