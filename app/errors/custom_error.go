package errors

import (
	"strings"
)

// TypeError define system error type
type TypeError string

// error types
var (
	TypeInternal            TypeError = "TYPE_INTERNAL"
	TypeServiceUnavailable  TypeError = "TYPE_SERVICE_UNAVAILABLE"
	TypeUnprocessableEntity TypeError = "TYPE_UNPROCESSABLE_ENTITY"
	TypeUnauthorized        TypeError = "TYPE_UNAUTHORIZED"
	TypeForbidden           TypeError = "TYPE_FORBIDDEN"
	TypeInvalidArgument     TypeError = "TYPE_INVALID_ARGUMENT"
	TypeNotFound            TypeError = "TYPE_NOT_FOUND"
	TypeConflict            TypeError = "TYPE_CONFLICT"
)

// Code defines system error code
type Code string

// error codes
// #nosec
//
//nolint:gosec
var (
	CodeInternal                      Code = "CODE_INTERNAL"
	CodeUnauthorized                  Code = "CODE_UNAUTHORIZED"
	CodeUnprocessableEntity           Code = "CODE_UNPROCESSABLE_ENTITY"
	CodeForbidden                     Code = "CODE_FORBIDDEN"
	CodeInvalidPayload                Code = "CODE_INVALID_PAYLOAD"
	CodeNoServiceActivationPermission Code = "CODE_NO_SERVICE_ACTIVATION_PERMISSION"

	// user error code
	CodeBadRequest          Code = "CODE_BAD_REQUEST"
	CodeUserUnauthorized    Code = "CODE_USER_UNAUTHORIZED"
	CodeNotFound            Code = "CODE_NOT_FOUND"
	CodeUserForbidden       Code = "CODE_USER_FORBIDDEN"
	CodeInternalUserExisted Code = "CODE_INTERNAL_USER_EXISTED"
	CodeInvalidPassword     Code = "CODE_INVALID_PASSWORD"

	// partner user error code
	CodePartnerUserUnauthorized Code = "CODE_PARTNER_USER_UNAUTHORIZED"
	CodePartnerUserInvalidEmail Code = "CODE_PARTNER_USER_INVALID_EMAIL"

	// error codes for resource not found
	CodeNotFoundContract                        Code = "CODE_NOT_FOUND_CONTRACT"
	CodeNotFoundWorkflow                        Code = "CODE_NOT_FOUND_WORKFLOW"
	CodeNotFoundWorkflowTemplate                Code = "CODE_NOT_FOUND_WORKFLOW_TEMPLATE"
	CodeNotFoundDocument                        Code = "CODE_NOT_FOUND_DOCUMENT"
	CodeNotFoundJob                             Code = "CODE_NOT_FOUND_JOB"
	CodeNotFoundNavisOffice                     Code = "CODE_NOT_FOUND_NAVIS_OFFICE"
	CodeNotFoundSfTemplate                      Code = "CODE_NOT_FOUND_SF_TEMPLATE"
	CodeNotFoundImportedCSV                     Code = "CODE_NOT_FOUND_IMPORTED_CSV"
	CodeNotFoundDraftContract                   Code = "CODE_NOT_FOUND_DRAFT_CONTRACT"
	CodeNotFoundAttachment                      Code = "CODE_NOT_FOUND_ATTACHMENT"
	CodeNotFoundHandoverSession                 Code = "CODE_NOT_FOUND_HANDOVER_SESSION"
	CodeNotFoundContractTemplate                Code = "CODE_NOT_FOUND_CONTRACT_TEMPLATE"
	CodeNotFoundContractTemplateAttachment      Code = "CODE_NOT_FOUND_CONTRACT_TEMPLATE_ATTACHMENT"
	CodeNotFoundBatchContractTemplate           Code = "CODE_NOT_FOUND_BATCH_CONTRACT_TEMPLATE"
	CodeNotFoundBatchContractTemplateAttachment Code = "CODE_NOT_FOUND_BATCH_CONTRACT_TEMPLATE_ATTACHMENT"
	CodeNotFoundBatchContractTemplateDocument   Code = "CODE_NOT_FOUND_BATCH_CONTRACT_TEMPLATE_DOCUMENT"
	CodeNotFoundEmailAggregation                Code = "CODE_NOT_FOUND_EMAIL_AGGREGATION"
	CodeNotFoundReceivedContract                Code = "CODE_NOT_FOUND_RECEIVED_CONTRACT"
	CodeNotFoundReceivedDocument                Code = "CODE_NOT_FOUND_RECEIVED_DOCUMENT"
	CodeNotFoundAnnouncement                    Code = "CODE_NOT_FOUND_ANNOUNCEMENT"
	CodeNotFoundUserSetting                     Code = "CODE_NOT_FOUND_USER_SETTING"
	CodeNotFoundContractFieldTemplate           Code = "CODE_NOT_FOUND_CONTRACT_FIELD_TEMPLATE"
	CodeNotFoundUserGroup                       Code = "CODE_NOT_FOUND_USER_GROUP"
	CodeNotFoundBatchSendingCsv                 Code = "CODE_NOT_FOUND_BATCH_SENDING_CSV"
	CodeNotFoundPartnerCompany                  Code = "CODE_NOT_FOUND_PARTNER_COMPANY"
	CodeNotFoundProposal                        Code = "CODE_NOT_FOUND_PROPOSAL"
	CodeNotFoundProposalPIC                     Code = "CODE_NOT_FOUND_PROPOSAL_PIC"
	CodeNotFoundProposalReviewer                Code = "CODE_NOT_FOUND_PROPOSAL_REVIEWER"
	CodeNotFoundProposalDocument                Code = "CODE_NOT_FOUND_PROPOSAL_DOCUMENT"
	CodeNotFoundProposalAppliedHistory          Code = "CODE_NOT_FOUND_PROPOSAL_APPLIED_HISTORY"
	CodeNotFoundProposalDocumentConverted       Code = "CODE_NOT_FOUND_PROPOSAL_DOCUMENT_CONVERTED"
	CodeNotFoundProposalMessage                 Code = "CODE_NOT_FOUND_PROPOSAL_MESSAGE"
	CodeNotFoundProposalPrivateMessage          Code = "CODE_NOT_FOUND_PROPOSAL_PRIVATE_MESSAGE"
	CodeNotFoundWebhook                         Code = "CODE_NOT_FOUND_WEBHOOK"
	CodeNotFoundAddressBook                     Code = "CODE_NOT_FOUND_ADDRESS_BOOK"
	CodeNotFoundMFIDUser                        Code = "CODE_NOT_FOUND_MFID_USER"

	CodeNotFoundMultipleContract           Code = "CODE_NOT_FOUND_MULTIPLE_CONTRACT"
	CodeNotFoundMultipleContractAttachment Code = "CODE_NOT_FOUND_MULTIPLE_CONTRACT_ATTACHMENT"

	// error codes for invalid ids
	CodeInvalidNavisOfficeIdentificationCode Code = "CODE_INVALID_NAVIS_OFFICE_IDENTIFICATION_CODE"
	CodeInvalidUserID                        Code = "CODE_INVALID_USER_ID"
	CodeInvalidMFIDUserID                    Code = "CODE_INVALID_MFID_USER_ID"
	CodeInvalidMFIDUID                       Code = "CODE_INVALID_MFID_UID"
	CodeInvalidOfficeID                      Code = "CODE_INVALID_OFFICE_ID"
	CodeInvalidDocumentID                    Code = "CODE_INVALID_DOCUMENT_ID"
	CodeInvalidContractID                    Code = "CODE_INVALID_CONTRACT_ID"

	CodeContractInProgress                              Code = "CODE_CONTRACT_IN_PROGRESS"
	CodeCanNotSetStampLocations                         Code = "CODE_CONTRACT_CAN_NOT_SET_STAMP_LOCATIONS"
	CodeCanNotSetCustomFields                           Code = "CODE_CONTRACT_CAN_NOT_SET_CUSTOM_FIELDS"
	CodeWorkflowStepHasNoAssignee                       Code = "CODE_WORKFLOW_STEP_HAS_NO_ASSIGNEE"
	CodeWorkflowMultipleImprintersInDifferenceStep      Code = "CODE_WORKFLOW_MULTIPLE_IMPRINTERS_IN_DIFFERENCE_STEP"
	CodeWorkflowExistNormalAssigneesInAuthorizeStep     Code = "CODE_WORKFLOW_EXIST_NORMAL_ASSIGNEES_IN_AUTHORIZE_STEP"
	CodeWorkflowExistNormalAssigneesInImprintStep       Code = "CODE_WORKFLOW_EXIST_NORMAL_ASSIGNEES_IN_IMPRINT_STEP"
	CodeWorkflowInternalStepHasNoUserID                 Code = "CODE_WORKFLOW_INTERNAL_STEP_HAS_NO_USER_ID"
	CodeWorkflowInternalStepAssigneeNotInSameOffice     Code = "CODE_WORKFLOW_INTERNAL_STEP_ASSIGNEE_NOT_IN_SAME_OFFICE"
	CodeWorkflowInternalStepImprinterNotMatchToContract Code = "CODE_WORKFLOW_INTERNAL_STEP_IMPRINTER_NOT_MATCH_TO_CONTRACT"
	CodeWorkflowExternalStepInvalidAccessKey            Code = "CODE_WORKFLOW_EXTERNAL_INVALID_ACCESS_KEY"
	CodeWorkflowExternalStepInvalidPartnerCompany       Code = "CODE_WORKFLOW_EXTERNAL_INVALID_PARTNER_COMPANY"
	CodeWorkflowExternalStepInvalidPartnerRequestNumber Code = "CODE_WORKFLOW_EXTERNAL_INVALID_PARTNER_REQUEST_NUMBER"
	CodeWorkflowExternalStepExceedAssigneeLimit         Code = "CODE_WORKFLOW_EXTERNAL_EXCEED_ASSIGNEE_LIMIT"
	CodeWorkflowInvalidStepsOrder                       Code = "CODE_WORKFLOW_INVALID_STEPS_ORDER"
	CodeWorkflowStepDuplicateAssignee                   Code = "CODE_WORKFLOW_STEP_DUPLICATE_ASSIGNEE"
	CodeWorkflowSameTemplateAlreadyExist                Code = "CODE_WORKFLOW_SAME_TEMPLATE_ALREADY_EXIST"

	// add viewers
	CodeViewerInvalidUserID           Code = "CODE_VIEWER_INVALID_USER_ID"
	CodeViewerInvalidUserGroupID      Code = "CODE_VIEWER_INVALID_USER_GROUP_ID"
	CodeViewerTemplateNotInSameOffice Code = "CODE_VIEWER_TEMPLATE_NOT_IN_SAME_OFFICE"

	// workflow template related invalid error codes
	CodeWorkflowTemplateOrderQuantityNotMatched Code = "CODE_WORKFLOW_TEMPLATE_ORDER_QUANTITY_NOT_MATCHED"
	CodeWorkflowTemplateStepHasNoAssignee       Code = "CODE_WORKFLOW_TEMPLATE_STEP_HAS_NO_ASSIGNEE"
	CodeWorkflowTemplateMissingUserID           Code = "CODE_WORKFLOW_TEMPLATE_MISSING_USER_ID"
	CodeWorkflowTemplateDeleteTheLastTemplate   Code = "CODE_WORKFLOW_TEMPLATE_DELETE_THE_LAST_TEMPLATE"

	// admin settings related error codes
	CodeUserInsufficientAuthorization            Code = "CODE_USER_INSUFFICIENT_AUTHORIZATION"
	CodeUserCanNotDeleteTheLastAdmin             Code = "CODE_USER_CAN_NOT_DELETE_THE_LAST_ADMIN"
	CodeUserNotBeInvitedToOffice                 Code = "CODE_USER_NOT_BE_INVITED_TO_OFFICE"
	CodeUserUpdateUserHasBeenDeleted             Code = "CODE_USER_UPDATE_USER_HAS_BEEN_DELETED"
	CodeUserUpdateCanNotUpdateRoleOfTheLastAdmin Code = "CODE_USER_UPDATE_CAN_NOT_UPDATE_ROLE_OF_THE_LAST_ADMIN"

	// error codes for internal APIs
	CodeTenantUserNotFound                    Code = "CODE_TENANT_USER_NOT_FOUND"
	CodeInternalFindContractsInvalidTenantUID Code = "CODE_INTERNAL_FIND_CONTRACTS_INVALID_TENANT_UID"
	CodeInternalFindInvalidAfterID            Code = "CODE_INTERNAL_FIND_INVALID_AFTER_ID"
	CodeInternalFindInvalidBeforeID           Code = "CODE_INTERNAL_FIND_INVALID_BEFORE_ID"

	// kitalive errors
	CodeKitAliveUnauthorized                          Code = "CODE_KIT_ALIVE_UNAUTHORIZED"
	CodeKitAliveInvalidObjectRecordAssigneeEmail      Code = "CODE_KIT_ALIVE_INVALID_OBJECT_RECORD_ASSIGNEE_EMAIL"
	CodeKitAliveInvalidObjectRecordAssigneeName       Code = "CODE_KIT_ALIVE_INVALID_OBJECT_RECORD_ASSIGNEE_NAME"
	CodeKitAliveInvalidObjectRecordContractName       Code = "CODE_KIT_ALIVE_INVALID_OBJECT_RECORD_CONTRACT_NAME"
	CodeKitAliveInvalidObjectRecordPartnerCompanyName Code = "CODE_KIT_ALIVE_INVALID_OBJECT_RECORD_PARTNER_COMPANY_NAME"

	// slack errors
	CodeSlackUnauthorized Code = "CODE_SLACK_UNAUTHORIZED"

	// internal partner errors
	CodeInternalPartnerUnauthorized         Code = "CODE_INTERNAL_PARTNER_UNAUTHORIZED"
	CodeContractForPartnerInvalidNumber     Code = "CODE_CONTRACT_FOR_PARTNER_INVALID_NUMBER"
	CodeContractForPartnerInvalidContractID Code = "CODE_CONTRACT_FOR_PARTNER_INVALID_CONTRACT_ID"
	CodeContractForPartnerContractNotFound  Code = "CODE_CONTRACT_FOR_PARTNER_CONTRACT_NOT_FOUND"
	CodeContractCanNotDeclineByAll          Code = "CODE_APPROVE_CONTRACT_CAN_NOT_DECLINE_BY_ALL"
	CodeInternalNoUsersInOffice             Code = "CODE_INTERNAL_NO_USERS_IN_OFFICE"
	CodeInternalInvalidUserEmail            Code = "CODE_INTERNAL_INVALID_USER_EMAIL"
	CodeInternalInsufficientUserRole        Code = "CODE_INTERNAL_INSUFFICIENT_USER_ROLE"

	// user errors
	CodeUserBeingInvolvedInWorkflowTemplates Code = "CODE_USER_BEING_INVOLVED_IN_WORKFLOW_TEMPLATES"
	CodeUserBeingInvolvedInPendingContracts  Code = "CODE_USER_BEING_INVOLVED_IN_PENDING_CONTRACTS"

	// user group errors
	CodeUserGroupAlreadyUsedInBatchContractTemplate Code = "CODE_USER_GROUP_ALREADY_USED_IN_BATCH_CONTRACT_TEMPLATE"
	CodeUserGroupAlreadyUsedInWorkflowTemplate      Code = "CODE_USER_GROUP_ALREADY_USED_IN_WORKFLOW_TEMPLATE"
	CodeUserGroupAlreadyUsedInPendingContracts      Code = "CODE_USER_GROUP_ALREADY_USED_IN_PENDING_CONTRACTS"
	CodeUserGroupAlreadyUsedInConcludedContracts    Code = "CODE_USER_GROUP_ALREADY_USED_IN_CONCLUDED_CONTRACTS"
	CodeUserGroupExceededMaximumUserLimit           Code = "CODE_USER_GROUP_EXCEEDED_MAXIMUM_USER_LIMIT"

	// job errors
	CodeJobInvalidStatusForRetry   Code = "CODE_JOB_INVALID_STATUS_FOR_RETRY"
	CodeJobInvalidStatusForResolve Code = "CODE_JOB_INVALID_STATUS_FOR_RESOLVE"
	// term of use errors
	CodeMFIDUserNeedToAcceptTermsOfUse Code = "CODE_MFID_USER_NEED_TO_ACCEPT_TERMS_OF_USE"

	CodeContractTypeInvalidContractTypeInformation Code = "CODE_CONTRACT_TYPE_INVALID_CONTRACT_TYPE_INFORMATION"
	CodeContractTypeDuplicateValue                 Code = "CODE_CONTRACT_TYPE_DUPLICATE_VALUE"
	CodeContractTypeExceedsQuantityLimit           Code = "CODE_CONTRACT_TYPE_EXCEEDS_QUANTITY_LIMIT"
	CodeContractTypeCannotDeleteTypesInUse         Code = "CODE_CONTRACT_TYPE_CAN_NOT_DELETE_TYPES_IN_USE"
	CodeContractTypeCannotDeleteDefaultType        Code = "CODE_CONTRACT_TYPE_CAN_NOT_DELETE_DEFAULT_TYPE"

	CodeStampTypeInvalidStampTypeInformation Code = "CODE_STAMP_TYPE_INVALID_STAMP_TYPE_INFORMATION"
	CodeStampTypeDuplicateValue              Code = "CODE_STAMP_TYPE_DUPLICATE_VALUE"
	CodeStampTypeExceedsQuantityLimit        Code = "CODE_STAMP_TYPE_EXCEEDS_QUANTITY_LIMIT"
	CodeStampTypeCannotDeleteDefaultType     Code = "CODE_STAMP_TYPE_CAN_NOT_DELETE_DEFAULT_TYPE"

	// salesforce errors
	CodeSalesforceOrgAlreadyLinked             Code = "CODE_SALESFORCE_ORG_ALREADY_LINKED"
	CodeSalesforceAppNotInstalledForUser       Code = "CODE_SALESFORCE_APP_NOT_INSTALLED_FOR_USER"
	CodeSalesforceOfficeAlreadyLinked          Code = "CODE_SALESFORCE_OFFICE_ALREADY_LINKED"
	CodeSalesforceOfficeNotYetLinked           Code = "CODE_SALESFORCE_OFFICE_NOT_YET_LINKED"
	CodeSalesforceInvalidSession               Code = "CODE_SALESFORCE_INVALID_SESSION"
	CodeSalesforceInvalidObjectName            Code = "CODE_SALESFORCE_INVALID_OBJECT_NAME"
	CodeSalesforceInvalidRecordID              Code = "CODE_SALESFORCE_INVALID_RECORD_ID"
	CodeSalesforceInvalidMappingField          Code = "CODE_SALESFORCE_INVALID_MAPPING_FIELD"
	CodeSalesforceInvalidRequestedEmail        Code = "CODE_SALESFORCE_INVALID_REQUESTED_EMAIL"
	CodeSalesforceInvalidObjectField           Code = "CODE_SALESFORCE_INVALID_OBJECT_FIELD"
	CodeSalesforceInvalidObjectFieldDataType   Code = "CODE_SALESFORCE_INVALID_OBJECT_FIELD_DATA_TYPE"
	CodeSalesforceInvalidObjectFieldNameFormat Code = "CODE_SALESFORCE_INVALID_OBJECT_FIELD_NAME_FORMAT"

	// sf template errors
	CodeSfTemplateDuplicateApprover                 Code = "CODE_SF_TEMPLATE_DUPLICATE_APPROVER"
	CodeSfTemplateImprinterMissing                  Code = "CODE_SF_TEMPLATE_IMPRINTER_MISSING"
	CodeSfTemplateMoreThanOneImprinter              Code = "CODE_SF_TEMPLATE_MORE_THAN_ONE_IMPRINTER"
	CodeSfTemplateSameTemplateAlreadyExisted        Code = "CODE_SF_TEMPLATE_SAME_TEMPLATE_ALREADY_EXISTED"
	CodeSfTemplateSetupNotCompleted                 Code = "CODE_SF_TEMPLATE_SETUP_NOT_COMPLETED"
	CodeSfTemplateInvalidMappingStatusDataType      Code = "CODE_SF_TEMPLATE_INVALID_MAPPING_STATUS_DATA_TYPE"
	CodeSfTemplateInvalidMappingContractURLDataType Code = "CODE_SF_TEMPLATE_INVALID_MAPPING_CONTRACT_URL_DATA_TYPE"

	// cron errors
	CodeCronDeleteUsersEndDateBeforeStartDate   Code = "CODE_CRON_DELETE_USERS_END_DATE_BEFORE_START_DATE"
	CodeCronDeleteTenantsEndDateBeforeStartDate Code = "CODE_CRON_DELETE_TENANTS_END_DATE_BEFORE_START_DATE"

	// import errors
	CodeCsvImportInvalidFormat          Code = "CODE_CSV_IMPORT_INVALID_FORMAT"
	CodeCsvImportInvalidFile            Code = "CODE_CSV_IMPORT_INVALID_FILE"
	CodeCsvImportRecordsLimitExceeded   Code = "CODE_CSV_IMPORT_RECORDS_LIMIT_EXCEEDED"
	CodeCsvImportEmptyCSV               Code = "CODE_CSV_IMPORT_EMPTY_CSV"
	CodeImportedCSVDeleteKindNewOnly    Code = "CODE_IMPORTED_CSV_DELETE_KIND_NEW_ONLY"
	CodeImportedCSVDeleteStatusDoneOnly Code = "CODE_IMPORTED_CSV_DELETE_STATUS_DONE_ONLY"

	// ERP Plan errors
	CodeErpPlanInsufficientPlan     Code = "CODE_ERP_PLAN_INSUFFICIENT_PLAN"
	CodeErpPlanExceedUserLimitation Code = "CODE_ERP_PLAN_EXCEED_USER_LIMITATION"

	// office usage errors
	CodeSalesforceLinkageRestriction Code = "CODE_SALESFORCE_LINKAGE_RESTRICTION"
	CodeSlackLinkageRestriction      Code = "CODE_SLACK_LINKAGE_RESTRICTION"
	CodeRelevantDocumentsRestriction Code = "CODE_RELEVANT_DOCUMENTS_RESTRICTION"
	CodeUserInsufficientLinkage      Code = "CODE_USER_INSUFFICIENT_LINKAGE"

	// contract template errors
	CodeContractTemplateDuplicatedName     Code = "CODE_CONTRACT_TEMPLATE_DUPLICATED_NAME"
	CodeContractTemplateInvalidInformation Code = "CODE_CONTRACT_TEMPLATE_INVALID_INFORMATION"
	CodeContractTemplateInvalidStampNumber Code = "CODE_CONTRACT_TEMPLATE_INVALID_STAMP_NUMBER"

	// email aggregation errors
	CodeEmailAggregationInvalidUserID    Code = "CODE_EMAIL_AGGREGATION_INVALID_USER_ID"
	CodeReceivedContractAlreadyConnected Code = "CODE_RECEIVED_CONTRACT_ALREADY_CONNECTED"

	// announcement errors
	CodeAnnouncementDuplicatedName Code = "CODE_ANNOUNCEMENT_DUPLICATED_NAME"

	// slack errors
	CodeSlackUserNotFound        Code = "CODE_SLACK_USER_NOT_FOUND"
	CodeSlackWorkspaceNotMatched Code = "CODE_SLACK_WORKSPACE_NOT_MATCHED"
	CodeSlackNoWorkspaceForUser  Code = "CODE_SLACK_NO_WORKSPACE_FOR_USER"

	// office connection errors
	CodeNavisOfficeConnectionNotFound    Code = "CODE_NAVIS_OFFICE_CONNECTION_NOT_FOUND"
	CodeNavisOfficeConnectionInvalidKind Code = "CODE_NAVIS_OFFICE_CONNECTION_INVALID_KIND"

	// locked errors
	CodeConcludedContractsExportLocked Code = "CODE_CONCLUDED_CONTRACTS_EXPORT_LOCKED"

	// custom contract field name
	CodeCustomContractFieldNameAlreadyUsed                        Code = "CODE_CUSTOM_CONTRACT_FIELD_NAME_ALREADY_USED"
	CodeCustomContractFieldNameAlreadyUsedByContract              Code = "CODE_CUSTOM_CONTRACT_FIELD_NAME_ALREADY_USED_BY_CONTRACT"
	CodeCustomContractFieldNameAlreadyUsedByContractFieldTemplate Code = "CODE_CUSTOM_CONTRACT_FIELD_NAME_ALREADY_USED_BY_CONTRACT_FIELD_TEMPLATE"
	CodeInvalidCustomContractFieldName                            Code = "CODE_INVALID_CUSTOM_CONTRACT_FIELD_NAME"

	// contract field template
	CodeContractFieldTemplateDeleteTheDefaultTemplate Code = "CODE_CONTRACT_FIELD_TEMPLATE_DELETE_THE_DEFAULT_TEMPLATE"
	CodeContractFieldTemplateOrderQuantityNotMatched  Code = "CODE_CONTRACT_FIELD_TEMPLATE_ORDER_QUANTITY_NOT_MATCHED"
	CodeContractFieldTemplateDeleteInUsed             Code = "CODE_CONTRACT_FIELD_TEMPLATE_DELETE_IN_USED"

	// office contract field
	CodeCustomContractFieldNameIDInvalid          Code = "CODE_CUSTOM_CONTRACT_FIELD_NAME_ID_INVALID"
	CodeOfficeContractFieldTemplateDuplicatedName Code = "CODE_OFFICE_CONTRACT_FIELD_TEMPLATE_DUPLICATED_NAME"
	CodeInvalidCustomContractFieldRadioMapping    Code = "CODE_INVALID_CUSTOM_CONTRACT_FIELD_RADIO_MAPPING"
	CodeLimitExceededCustomContractFields         Code = "CODE_LIMIT_EXCEEDED_CUSTOM_CONTRACT_FIELDS"
	CodeMissingDefaultContractFields              Code = "CODE_MISSING_DEFAULT_CONTRACT_FIELDS"

	// user group errors
	CodeUserGroupNotFound       Code = "CODE_USER_GROUP_NOT_FOUND"
	CodeUserGroupDuplicatedName Code = "CODE_USER_GROUP_DUPLICATED_NAME"

	// batch contract template
	CodeBatchContractTemplateNameDuplicated                Code = "CODE_BATCH_CONTRACT_TEMPLATE_NAME_DUPLICATED"
	CodeBatchContractTemplateStatusNotFinished             Code = "CODE_BATCH_CONTRACT_TEMPLATE_STATUS_NOT_FINISHED"
	CodeBatchContractTemplateExceedImprintersLimit         Code = "CODE_BATCH_CONTRACT_TEMPLATE_EXCEED_IMPRINTERS_LIMIT"
	CodeBatchContractTemplateUsedInContracts               Code = "CODE_BATCH_CONTRACT_TEMPLATE_USED_IN_CONTRACTS"
	CodeBatchContractTemplateUnlinkableInternalCustomField Code = "CODE_BATCH_CONTRACT_TEMPLATE_UNLINKABLE_INTERNAL_CUSTOM_FIELD"
	CodeBatchContractTemplateUnmatchedTypeCustomField      Code = "CODE_BATCH_CONTRACT_TEMPLATE_UNMATCHED_TYPE_CUSTOM_FIELD"
	CodeBatchContractTemplateExistedContractFieldName      Code = "CODE_BATCH_CONTRACT_TEMPLATE_EXISTED_CONTRACT_FIELD_NAME"

	// webhook
	CodeWebhookDuplicatedURL      Code = "CODE_WEBHOOK_DUPLICATED_URL"
	CodeWebhookExceedNumberOfURLs Code = "CODE_WEBHOOK_EXCEED_NUMBER_OF_URLS"

	// proposal
	CodeConflictProposal                                     Code = "CODE_CONFLICT_PROPOSAL"
	CodeConflictProposalDocumentVersion                      Code = "CODE_CONFLICT_PROPOSAL_DOCUMENT_VERSION"
	CodeProposalAlreadyLinkedAuto                            Code = "CODE_PROPOSAL_ALREADY_LINKED_AUTO"
	CodeProposalAlreadyLinkedManual                          Code = "CODE_PROPOSAL_ALREADY_LINKED_MANUAL"
	CodeProposalPersonInChargeNotBasedGeneralRole            Code = "CODE_PROPOSAL_PERSON_IN_CHARGE_NOT_BASED_GENERAL_ROLE"
	CodeProposalPersonInChargeAlreadyInvitedAsProposalMember Code = "CODE_PROPOSAL_PERSON_IN_CHARGE_ALREADY_INVITED_AS_PROPOSAL_MEMBER"
	CodeProposalInvalidMembersOrUserGroups                   Code = "CODE_PROPOSAL_INVALID_MEMBERS_OR_USER_GROUP"
	CodeProposalUnsupportedDocumentType                      Code = "CODE_PROPOSAL_UNSUPPORTED_DOCUMENT_TYPE"
	CodeProposalUnsupportedAttachmentType                    Code = "CODE_PROPOSAL_UNSUPPORTED_ATTACHMENT_TYPE"
	CodeProposalUploadingDocumentUnauthorized                Code = "CODE_PROPOSAL_UPLOADING_DOCUMENT_UNAUTHORIZED"
	CodeProposalDeletingAttachmentUnauthorized               Code = "CODE_PROPOSAL_DELETING_ATTACHMENT_UNAUTHORIZED"
	CodeProposalUserInsufficientSendMessage                  Code = "CODE_PROPOSAL_USER_INSUFFICIENT_SEND_MESSAGE"
	CodeMessageCannotBePushedPusher                          Code = "CODE_MESSAGE_CANNOT_BE_PUSHED_PUSHER"
	CodeProposalUserInsufficientDeleteMessage                Code = "CODE_PROPOSAL_USER_INSUFFICIENT_DELETE_MESSAGE"
	CodeProposalMessageDeleteOldDocument                     Code = "CODE_PROPOSAL_MESSAGE_DELETE_OLD_DOCUMENT"
	CodeProposalUserInsufficientReturnProposal               Code = "CODE_PROPOSAL_USER_INSUFFICIENT_RETURN_PROPOSAL"
	CodeProposalsInvalidContractStartDateTo                  Code = "CODE_PROPOSALS_INVALID_CONTRACT_START_DATE_TO"
	CodeProposalsInvalidCreatedDateFrom                      Code = "CODE_PROPOSALS_INVALID_CREATED_DATE_FROM"
	CodeProposalsInvalidCreatedDateTo                        Code = "CODE_PROPOSALS_INVALID_CREATED_DATE_TO"
	CodeProposalDuplicatedAttachmentDocumentPath             Code = "CODE_PROPOSAL_DUPLICATED_ATTACHMENT_DOCUMENT_PATH"
	CodeProposalInvalidMentionMemberRole                     Code = "CODE_PROPOSAL_INVALID_MENTION_MEMBER_ROLE"
	CodeProposalMissingAssigneeInformation                   Code = "CODE_PROPOSAL_MISSING_ASSIGNEE_INFORMATION"

	// legal check
	CodeLegalCheckRestriction Code = "CODE_LEGAL_CHECK_RESTRICTION"

	// address book
	CodeAddressBookDuplicatedPartnerEmail Code = "CODE_ADDRESS_BOOK_DUPLICATED_PARTNER_EMAIL"
	CodeAddressBookExceededMaxLimit       Code = "CODE_ADDRESS_BOOK_EXCEEDED_MAX_LIMIT"

	// contract expiration
	CodeExpiredContract Code = "CODE_EXPIRED_CONTRACT"

	// document extraction
	CodeDocumentExtractionInvalidStatus Code = "CODE_DOCUMENT_EXTRACTION_INVALID_STATUS"
	CodeDocumentExtractionNotFound      Code = "CODE_DOCUMENT_EXTRACTION_NOT_FOUND"
)

// SystemError define system error
type SystemError interface {
	Type() TypeError
	Code() Code
	Message() string
	Param() interface{}
	StatusCode() int
	Error() string
}

// SystemErrors an array of system errors
type SystemErrors []SystemError

// Error implements error interface
func (errs SystemErrors) Error() string {
	errsString := make([]string, 0, len(errs))
	for _, err := range errs {
		errsString = append(errsString, err.Error())
	}

	return strings.Join(errsString, "\n")
}
