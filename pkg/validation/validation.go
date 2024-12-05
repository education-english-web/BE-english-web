package validation

import (
	"math"
	"regexp"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
)

var (
	globalValidator    *validator.Validate
	enUniTrans         *ut.UniversalTranslator
	englishTranslator  ut.Translator
	enLocales          locales.Translator
	jaUniTrans         *ut.UniversalTranslator
	japaneseTranslator ut.Translator
	jpLocales          locales.Translator
)

const (
	japanese = "ja"
	english  = "en"
)

func init() {
	globalValidator = validator.New()
	_ = globalValidator.RegisterValidation("only_upper_case", customOnlyUpperCase)
	_ = globalValidator.RegisterValidation("renew_duration", renewDuration)
	_ = globalValidator.RegisterValidation("cancel_auto_renew_duration", cancelAutoRenewDuration)
	_ = globalValidator.RegisterValidation("contract_value", contractValue)

	enLocales = en.New()
	enUniTrans = ut.New(enLocales, enLocales)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	englishTranslator, _ = enUniTrans.GetTranslator(english)
	_ = en_translations.RegisterDefaultTranslations(globalValidator, englishTranslator)

	jpLocales = ja.New()
	jaUniTrans = ut.New(jpLocales, jpLocales)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	japaneseTranslator, _ = jaUniTrans.GetTranslator(japanese)
	_ = ja_translations.RegisterDefaultTranslations(globalValidator, japaneseTranslator)
}

// GetTranslator to get translator with language
func GetTranslator(lang string) ut.Translator {
	switch lang {
	case japanese:
		return japaneseTranslator
	default:
		return englishTranslator
	}
}

// GetInstance global validator
func GetInstance() *validator.Validate {
	return globalValidator
}

func customOnlyUpperCase(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^([A-Z])\w+`)

	return regex.MatchString(fl.Field().String())
}

func renewDuration(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^([1-9]|[12][0-9]|3[01])\|D$|^([1-9]|1[0-2])\|M$|^([1-9]|[1-9][0-9])\|Y$`)

	return regex.MatchString(fl.Field().String())
}

func cancelAutoRenewDuration(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^([1-9]|[12]\d|3[01])\|D$|^([1-9]|1[012])\|M$`)

	return regex.MatchString(fl.Field().String())
}

func contractValue(fl validator.FieldLevel) bool {
	v := fl.Field().Float()

	return v == math.Trunc(v) && (v >= 0 && v < 1000000000000)
}
