package context

type key string

func (k key) String() string {
	return "middleware.context.key." + string(k)
}

// context keys
const (
	keyClientIP                     key = "client_ip"
	keyUserAgent                    key = "user_agent"
	keyUserID                       key = "user_id"
	keyNavisOfficeID                key = "navis_office_id"
	keyTenantUID                    key = "tenant_uid"
	keyPartnerCompanyID             key = "partner_company_id"
	keyPartnerEmail                 key = "partner_email"
	keyPartnerAccessKey             key = "partner_access_key"
	keyUserEmail                    key = "user_email"
	keyPhoneNumber                  key = "phone_number"
	keySalesforceOrgID              key = "salesforce_org_id"
	keySalesforceRequestedEmail     key = "salesforce_requested_email"
	keyMFIDUserID                   key = "mfid_user_id"
	keyOfficePlan                   key = "office_plan"
	keyOfficePlanOrigin             key = "office_plan_origin"
	keyOfficeUsageEndDate           key = "office_usage_end_date"
	keyOfficeUsageStartDate         key = "office_usage_start_date"
	keyOfficeUsageSalesforceLinkage key = "office_usage_salesforce_linkage"
	keyOfficeUsageRelevantDocuments key = "office_usage_relevant_documents"
	keyUserRoleCode                 key = "user_role_code"
	keyUserLinkages                 key = "user_linkages"
	keyProxyLoginEventID            key = "proxy_login_event_id"
	keyContractID                   key = "contract_id"
)
