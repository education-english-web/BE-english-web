package registry

import (
	"context"
	"encoding/base64"
	"encoding/pem"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/wire"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"github.com/education-english-web/BE-english-web/app/domain/service"
	"github.com/education-english-web/BE-english-web/app/external/adapter/oidcclient"

	//"github.com/education-english-web/BE-english-web/app/external/adapter/restclient/zendesk"
	//"github.com/education-english-web/BE-english-web/app/external/adapter/slackclient"
	"github.com/education-english-web/BE-english-web/app/services"
	//"github.com/education-english-web/BE-english-web/app/usecases/emailfactory"
	//"github.com/education-english-web/BE-english-web/app/usecases/emailfactory/emailtemplates"
	"github.com/education-english-web/BE-english-web/pkg/api"
	"github.com/education-english-web/BE-english-web/pkg/config"
	"github.com/education-english-web/BE-english-web/pkg/encoder"
	"github.com/education-english-web/BE-english-web/pkg/gormutil"
	"github.com/education-english-web/BE-english-web/pkg/hashid"
	"github.com/education-english-web/BE-english-web/pkg/hashing"
	"github.com/education-english-web/BE-english-web/pkg/hashing/md5"
	"github.com/education-english-web/BE-english-web/pkg/htmltemplate"
	"github.com/education-english-web/BE-english-web/pkg/htmltemplate/gohtmltemplate"
	"github.com/education-english-web/BE-english-web/pkg/htmltoimage"
	"github.com/education-english-web/BE-english-web/pkg/htmltoimage/chromedp"
	"github.com/education-english-web/BE-english-web/pkg/htmltopdf"
	"github.com/education-english-web/BE-english-web/pkg/htmltopdf/wkhtmltopdf"
	appLog "github.com/education-english-web/BE-english-web/pkg/log"
	"github.com/education-english-web/BE-english-web/pkg/logkeeper"
	"github.com/education-english-web/BE-english-web/pkg/notifier/event"
	"github.com/education-english-web/BE-english-web/pkg/notifier/eventfactory"
	"github.com/education-english-web/BE-english-web/pkg/notifier/slack"
	"github.com/education-english-web/BE-english-web/pkg/pdfutil"
	"github.com/education-english-web/BE-english-web/pkg/pdfutil/pdfcpu"
	"github.com/education-english-web/BE-english-web/pkg/pdfutil/pdfhelper"
	"github.com/education-english-web/BE-english-web/pkg/pusherclient"
	"github.com/education-english-web/BE-english-web/pkg/redis"
	"github.com/education-english-web/BE-english-web/pkg/spreadsheet"
	"github.com/education-english-web/BE-english-web/pkg/spreadsheet/googlesheet"
	"github.com/education-english-web/BE-english-web/pkg/timeutil"
	"github.com/education-english-web/BE-english-web/pkg/uuidint/sonyflake"
	"github.com/education-english-web/BE-english-web/pkg/uuidstring"
	"github.com/education-english-web/BE-english-web/pkg/uuidstring/ulid"
)

// Dependency Injection: All singleton set for wire generate
var (
	singletonSet = wire.NewSet(
		config.GetConfig,
		gormutil.GetDB,
		encoder.NewEncoder,
		NewUUIDGenerator,
		NewJWT,
		NewHashPassword,
		//NewEmailFactory,
		//NewSendGridMailer,
		NewTimeFactory,
		//NewHTTPFactory,
		NewPDFFactory,
		NewPDFHelper,
		NewMD5Hashing,
		NewHTMLToPDFConverter,
		NewHTMLTemplateParser,
		NewHTMLToImageGenerator,
		redis.Get,
		NewEventFactory,
		slack.New,
		logkeeper.GetCloudWatch,
		hashid.GetIDHasher,
		//NewZendeskService,
		//ProvideSlackService,
		//ProvideVerifiers,
		api.NewClient,
		pusherclient.Get,
		sonyflake.Get,
	)
)

//// NewEmailFactory returns new instance of email factory
//func NewEmailFactory() emailfactory.EmailFactory {
//	fqdn := config.GetConfig().FQDN
//
//	return emailfactory.NewEmailFactory(
//		emailtemplates.NewInternalApproverTemplate(config.GetConfig().SendGrid.InternalApprovalEmailTemplateID, fqdn),
//		emailtemplates.NewPartnerApproverTemplate(config.GetConfig().SendGrid.PartnerApprovalEmailTemplateID, fqdn),
//		emailtemplates.NewContractRejectedTemplate(config.GetConfig().SendGrid.ContractRejectedEmailTemplateID, fqdn),
//		emailtemplates.NewPartnerApproveAddedTemplate(config.GetConfig().SendGrid.PartnerUserDelegatedEmailTemplateID, fqdn),
//		emailtemplates.NewInternalConclusionTemplate(config.GetConfig().SendGrid.ContractConcludedInternalEmailTemplateID, fqdn),
//		emailtemplates.NewPartnerConclusionTemplate(config.GetConfig().SendGrid.ContractConcludedPartnerEmailTemplateID, fqdn),
//		emailtemplates.NewUserInvitationTemplate(config.GetConfig().SendGrid.UserInvitationEmailTemplateID, fqdn),
//		emailtemplates.NewPaperContractImprinterApproved(config.GetConfig().SendGrid.PaperContractImprinterApprovedEmailTemplateID, fqdn),
//		emailtemplates.NewPaperContractConclusionTemplate(config.GetConfig().SendGrid.TemplateIDPaperContractConcludedEmail, fqdn),
//		emailtemplates.NewContractEndDateAlertTemplate(config.GetConfig().SendGrid.TemplateIDContractEndDateAlert, fqdn),
//		emailtemplates.NewContractAutoRenewAlertTemplate(config.GetConfig().SendGrid.TemplateIDContractAutoRenewAlert, fqdn),
//		emailtemplates.NewHandoverLastAdminTemplate(config.GetConfig().SendGrid.TemplateIDHandoverLastAdmin, fqdn),
//		emailtemplates.NewEmailAggregationTemplate(config.GetConfig().SendGrid.TemplateIDEmailAggregation, fqdn),
//		emailtemplates.NewEmailAggregationInvitationTemplate(config.GetConfig().SendGrid.TemplateIDEmailAggregationInvitation, fqdn),
//		emailtemplates.NewInternalUserRemandedTemplate(config.GetConfig().SendGrid.InternalUserRemandedEmailTemplateID, fqdn),
//		emailtemplates.NewApplicantWithdrawContractTemplate(config.GetConfig().SendGrid.TemplateIDApplicantWithdrawContract, fqdn),
//		emailtemplates.NewPartnerApproverEnglishTemplate(config.GetConfig().SendGrid.TemplateIDPartnerApprovalEnglish, fqdn),
//		emailtemplates.NewPartnerApproveAddedEnglishTemplate(config.GetConfig().SendGrid.TemplateIDPartnerUserDelegatedEnglish, fqdn),
//		emailtemplates.NewPartnerConclusionEnglishTemplate(config.GetConfig().SendGrid.TemplateIDContractConcludedPartnerEnglish, fqdn),
//		emailtemplates.NewPartnerRejectionInternalTemplate(config.GetConfig().SendGrid.TemplateIDPartnerRejectedInternalEmail, fqdn),
//		emailtemplates.NewPartnerRejectionTemplate(config.GetConfig().SendGrid.TemplateIDPartnerRejected, fqdn),
//		emailtemplates.NewPartnerRejectionEnglishTemplate(config.GetConfig().SendGrid.TemplateIDPartnerRejectedEnglish, fqdn),
//		emailtemplates.NewProposalMemberAddedTemplate(config.GetConfig().SendGrid.TemplateIDProposalMemberAdded, fqdn),
//		emailtemplates.NewProposalMemberMentionedTemplate(config.GetConfig().SendGrid.TemplateIDProposalMemberMentioned, fqdn),
//		emailtemplates.NewProposalStatusDoneTemplate(config.GetConfig().SendGrid.TemplateIDProposalStatusDone, fqdn),
//		emailtemplates.NewAllInternalsApprovedTemplate(config.GetConfig().SendGrid.TemplateIDAllInternalsApproved, fqdn),
//		emailtemplates.NewProposalAssigneeEmailTemplate(config.GetConfig().SendGrid.TemplateIDProposalAssigneeEmail, fqdn),
//		emailtemplates.NewRemindContractExpirationTemplate(config.GetConfig().SendGrid.TemplateIDRemindContractExpiration),
//		emailtemplates.NewCreatorSFTemplateChangedTemplate(config.GetConfig().SendGrid.TemplateIDCreatorSFTemplateChanged),
//	)
//}

//// NewSendGridMailer return new instance of sendgrid mailer
//func NewSendGridMailer() mailer.Mailer {
//	return sendgrid.New(
//		config.GetConfig().Env,
//		config.GetConfig().SendGrid.AllowDomains,
//		config.GetConfig().SendGrid.SenderName,
//		config.GetConfig().SendGrid.SenderEmail,
//		config.GetConfig().SendGrid.APIKey,
//	)
//}

// NewTimeFactory return new instance of time factory
func NewTimeFactory() timeutil.TimeFactory {
	return timeutil.NewTimeFactory()
}

//// NewHTTPFactory return new instance of http factory
//func NewHTTPFactory() httputil.HTTPFactory {
//	return httputil.NewHTTPFactory()
//}

// NewUUIDGenerator return an instance of uuid generator
func NewUUIDGenerator() uuidstring.UUIDString {
	return ulid.New()
}

// NewPDFFactory return new instance of pdf factory
func NewPDFFactory() pdfutil.PDFFactory {
	return pdfcpu.New()
}

// NewPDFHelper return new instance of pdf helper
func NewPDFHelper() pdfutil.PDFHelper {
	return pdfhelper.New()
}

// NewJWT provides services.JWT
func NewJWT() services.JWT {
	return services.NewJWT(config.GetConfig().JWTSecret)
}

// NewHashPassword return new instance of password hashing
func NewHashPassword() services.HashPass {
	return services.NewHashPass(config.GetConfig().Salt)
}

// NewMD5Hashing return new instance of md5 hashing
func NewMD5Hashing() hashing.Hashing {
	return md5.New()
}

// NewHTMLToPDFConverter return new instance of html to pdf converter
func NewHTMLToPDFConverter() htmltopdf.Converter {
	return wkhtmltopdf.New()
}

// NewHTMLTemplateParser return new instance of html template parser
func NewHTMLTemplateParser() htmltemplate.Parser {
	return gohtmltemplate.New()
}

// NewHTMLToImageGenerator return new instance of html to image converter
func NewHTMLToImageGenerator() htmltoimage.Generator {
	return chromedp.New()
}

//func ProvideMFIDOAuth2Service() service.OIDCService {
//	issuer := config.GetConfig().MFID.BaseEndpoint
//
//	// new config
//	oauth2Cnf := oauth2.Config{
//		ClientID:     config.GetConfig().MFID.ClientID,
//		ClientSecret: config.GetConfig().MFID.ClientSecret,
//		Scopes:       []string{oidc.ScopeOpenID, "email"},
//		Endpoint: oauth2.Endpoint{
//			TokenURL: config.GetConfig().MFID.BaseEndpoint + "oauth/token",
//		},
//	}
//	// new verifier
//	verifier := oidc.NewVerifier(
//		issuer,
//		oidc.NewRemoteKeySet(context.Background(), issuer+"oauth/discovery/keys"),
//		&oidc.Config{
//			ClientID: config.GetConfig().MFID.ClientID,
//		},
//	)
//
//	return oidcclient.NewOIDCService(&oauth2Cnf, verifier, issuer, nil, nil)
//}

func ProvideGoogleOAuth2Service() service.OIDCService {
	issuer := "https://accounts.google.com"

	// new config
	oauth2Cnf := oauth2.Config{
		ClientID:     config.GetConfig().GSuite.ClientID,
		ClientSecret: config.GetConfig().GSuite.ClientSecret,
		Scopes:       []string{oidc.ScopeOpenID, "email"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}
	// new verifier
	verifier := oidc.NewVerifier(
		issuer,
		oidc.NewRemoteKeySet(context.Background(), "https://www.googleapis.com/oauth2/v3/certs"),
		&oidc.Config{
			ClientID: config.GetConfig().GSuite.ClientID,
		},
	)

	return oidcclient.NewOIDCService(&oauth2Cnf, verifier, issuer, config.GetConfig().OIDCAllowDomains, nil)
}

func ProvideSalesforceOAuth2Service() service.OIDCService {
	issuer := "https://login.salesforce.com"

	// new config
	oauth2Cnf := oauth2.Config{
		ClientID:     config.GetConfig().Salesforce.ClientID,
		ClientSecret: config.GetConfig().Salesforce.ClientSecret,
		Scopes:       []string{oidc.ScopeOpenID, "id api web lightning email refresh_token"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://login.salesforce.com/services/oauth2/token",
		},
	}
	// new verifier
	verifier := oidc.NewVerifier(
		issuer,
		oidc.NewRemoteKeySet(context.Background(), "https://login.salesforce.com/id/keys"),
		&oidc.Config{
			ClientID: config.GetConfig().Salesforce.ClientID,
		},
	)

	privateKey, err := base64.StdEncoding.DecodeString(config.GetConfig().Salesforce.PrivateKey)
	if err != nil {
		appLog.WithError(err).
			WithField("message", "error parsing salesforce private key").
			Errorln()
	}

	block, _ := pem.Decode(privateKey)
	if block == nil {
		appLog.WithField("message", "invalid salesforce private key provided").
			Errorln()
	}

	jwtConfig := jwt.Config{
		Email:      config.GetConfig().Salesforce.ClientID,
		PrivateKey: privateKey,
		Expires:    3 * time.Minute,
		Audience:   issuer,
		TokenURL:   "https://login.salesforce.com/services/oauth2/token",
	}

	return oidcclient.NewOIDCService(&oauth2Cnf, verifier, issuer, nil, &jwtConfig)
}

//func ProvideSlackService() service.SlackService {
//	clientID := config.GetConfig().SlackApp.ClientID
//	clientSecret := config.GetConfig().SlackApp.ClientSecret
//
//	return slackclient.NewSlackService(clientID, clientSecret)
//}
//
//func ProvideAzureADOauth2Service() service.OIDCService {
//	tenantID := config.GetConfig().AzureAD.TenantID
//	issuer := fmt.Sprintf("https://login.microsoftonline.com/%s/v2.0", tenantID)
//
//	oauth2Cnf := oauth2.Config{
//		ClientID:     config.GetConfig().AzureAD.ClientID,
//		ClientSecret: config.GetConfig().AzureAD.ClientSecret,
//		Scopes:       []string{oidc.ScopeOpenID, "email"},
//		Endpoint: oauth2.Endpoint{
//			TokenURL: fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantID),
//		},
//	}
//
//	verifier := oidc.NewVerifier(
//		issuer,
//		oidc.NewRemoteKeySet(context.Background(), fmt.Sprintf("https://login.microsoftonline.com/%s/discovery/v2.0/keys", tenantID)),
//		&oidc.Config{
//			ClientID: config.GetConfig().AzureAD.ClientID,
//		},
//	)
//
//	return oidcclient.NewOIDCService(&oauth2Cnf, verifier, issuer, nil, nil)
//}
//
//func ProvideVerifiers() partnerverifier.PartnerVerifier {
//	// initialize KitAlive signature verifier
//	kitAliveVerifier, err := rsapkcs1v15.New(config.GetConfig().Partner.KitAlivePublicKey)
//	if err != nil {
//		panic(fmt.Errorf("init kitalive verifier: %w", err))
//	}
//
//	// initialize Payable signature verifier
//	payableSigVerifier, err := rsapkcs1v15.New(config.GetConfig().Partner.PayablePublicKey)
//	if err != nil {
//		panic(fmt.Errorf("init payable verifier: %w", err))
//	}
//
//	payableIPVerifier := whitelistips.New(config.GetConfig().Partner.PayableAllowedIPs)
//
//	// initialize partner verifiers
//	return partnerverifier.InitPartnerVerifiers(
//		map[partnerverifier.PartnerName]partnerverifier.Verifier{
//			partnerverifier.PartnerNameKitAlive: {
//				SigVerifier: kitAliveVerifier,
//			},
//			partnerverifier.PartnerNamePayable: {
//				SigVerifier: payableSigVerifier,
//				IPVerifier:  payableIPVerifier,
//			},
//		},
//	)
//}

// NewEventFactory return an instance of event factory
func NewEventFactory() eventfactory.EventFactory {
	env := config.GetConfig().Env
	eventConfig := eventfactory.GetEventConfig()

	return eventfactory.NewEventFactory(
		event.NewEventInternalUserLogin(
			env,
			eventConfig[event.EventNameInternalUserLogin.String()].MessageTemplate,
			eventConfig[event.EventNameInternalUserLogin.String()].Channels,
		),
		event.NewEventInternalUserAdd(
			env,
			eventConfig[event.EventNameInternalUserAdd.String()].MessageTemplate,
			eventConfig[event.EventNameInternalUserAdd.String()].Channels,
		),
		event.NewEventInternalUserUpdateRole(
			env,
			eventConfig[event.EventNameInternalUserUpdateRole.String()].MessageTemplate,
			eventConfig[event.EventNameInternalUserUpdateRole.String()].Channels,
		),
		event.NewEventDeleteOffice(
			env,
			eventConfig[event.EventNameDeleteOffice.String()].MessageTemplate,
			eventConfig[event.EventNameDeleteOffice.String()].Channels,
		),
		event.NewEventForcefullyAddUser(
			env,
			eventConfig[event.EventNameForcefullyAddUser.String()].MessageTemplate,
			eventConfig[event.EventNameForcefullyAddUser.String()].Channels,
		),
		event.NewEventCreateOfficeUsage(
			env,
			eventConfig[event.EventNameCreateOfficeUsage.String()].MessageTemplate,
			eventConfig[event.EventNameCreateOfficeUsage.String()].Channels,
		),
		event.NewEventUpdateUserSetting(
			env,
			eventConfig[event.EventNameUpdateUserSetting.String()].MessageTemplate,
			eventConfig[event.EventNameUpdateUserSetting.String()].Channels,
		),
		event.NewEventFailConclusionJob(
			env,
			eventConfig[event.EventNameFailConclusionJob.String()].MessageTemplate,
			eventConfig[event.EventNameFailConclusionJob.String()].Channels,
		),
		event.NewEventFailMFCBoxJob(
			env,
			eventConfig[event.EventNameFailMFCBoxJob.String()].MessageTemplate,
			eventConfig[event.EventNameFailMFCBoxJob.String()].Channels,
		),
		event.NewEventProxyLogin(
			env,
			eventConfig[event.EventNameProxyLogin.String()].MessageTemplate,
			eventConfig[event.EventNameProxyLogin.String()].Channels,
		),
		event.NewEventTurnOffLegalCheck(
			env,
			eventConfig[event.EventNameTurnOffLegalCheck.String()].MessageTemplate,
			eventConfig[event.EventNameTurnOffLegalCheck.String()].Channels,
		),
	)
}

func ProvideOfficesReportSpreadSheetClient(spreadSheetID string) spreadsheet.Client {
	privateKey, err := base64.StdEncoding.DecodeString(config.GetConfig().GServiceAccount.PrivateKey)
	if err != nil {
		appLog.WithError(err).
			WithField("message", "error parsing google service account private key").
			Errorln()
	}

	block, _ := pem.Decode(privateKey)
	if block == nil {
		appLog.WithField("message", "invalid google service account private key provided").
			Errorln()
	}

	cfg := &jwt.Config{
		Email:        config.GetConfig().GServiceAccount.Email,
		PrivateKey:   privateKey,
		PrivateKeyID: config.GetConfig().GServiceAccount.PrivateKeyID,
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets",
		},
		TokenURL: "https://oauth2.googleapis.com/token",
	}
	client := cfg.Client(context.Background())

	sheetService, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		appLog.WithError(err).WithField("message", "error providing sheet client").
			Errorln()
	}

	return googlesheet.New(sheetService, spreadSheetID)
}

//func NewZendeskService() service.ZendeskService {
//	return zendesk.New(
//		config.GetConfig().Zendesk.Subdomain,
//		config.GetConfig().Zendesk.Email,
//		config.GetConfig().Zendesk.Token,
//		config.GetConfig().Zendesk.ScheduleID,
//	)
//}
