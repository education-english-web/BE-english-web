package presenter

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	appErrors "github.com/education-english-web/BE-english-web/app/errors"
	appContext "github.com/education-english-web/BE-english-web/app/interface/api/context"
	appLog "github.com/education-english-web/BE-english-web/pkg/log"
	"github.com/education-english-web/BE-english-web/pkg/tracer"
)

type response struct {
	Data   interface{}     `json:"data,omitempty"`
	Paging interface{}     `json:"paging,omitempty"`
	Errors []responseError `json:"errors,omitempty"`
}

// ResponsePaging holds paging response information
type ResponsePaging struct {
	Total uint32 `json:"total"`
}

type ResponseCursor64Paging struct {
	BeforeID    string `json:"before_id,omitempty"`
	IsLastPage  bool   `json:"is_last_page,omitempty"`
	AfterID     string `json:"after_id,omitempty"`
	IsFirstPage bool   `json:"is_first_page,omitempty"`
}

type responseError struct {
	Type    string      `json:"type"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Param   interface{} `json:"param,omitempty"`
}

// RenderErrors returns error response
func RenderErrors(ctx *gin.Context, span tracer.Span, err error) {
	var errs appErrors.SystemErrors
	if errors.As(err, &errs) {
		// if err is a list of errors, we assume they are validation errors,
		// so we will always return http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response{
			Errors: fromSystemErrors(errs),
		})

		traceErrors(span, errs)

		return
	}

	var e appErrors.SystemError
	if errors.As(err, &e) {
		ctx.JSON(e.StatusCode(), response{
			Errors: fromSystemErrors(appErrors.SystemErrors{e}),
		})

		traceErrors(span, appErrors.SystemErrors{e})

		return
	}

	if errors.Is(err, appErrors.ErrContextCancelled) {
		appLog.
			WithError(err).
			WithField("url", ctx.Request.URL.String()).
			WithField("method", ctx.Request.Method).
			WithField("user_agent", ctx.Request.UserAgent()).
			WithField("ip", ctx.ClientIP()).
			WithField("user_id", appContext.GetUserID(ctx)).
			WithField("navis_office_id", appContext.GetNavisOfficeID(ctx)).Warningln()

		return
	}

	appLog.
		WithField("url", ctx.Request.URL.String()).
		WithField("method", ctx.Request.Method).
		WithField("user_agent", ctx.Request.UserAgent()).
		WithField("ip", ctx.ClientIP()).
		WithField("user_id", appContext.GetUserID(ctx)).
		WithField("navis_office_id", appContext.GetNavisOfficeID(ctx)).
		WithError(err).Errorln("internal server error")

	ctx.JSON(http.StatusInternalServerError, response{
		Errors: []responseError{{
			Type:    string(appErrors.TypeInternal),
			Code:    string(appErrors.CodeInternal),
			Message: "internal server error",
			Param:   nil,
		}},
	})
}

// RenderData returns data response
func RenderData(ctx *gin.Context, data, paging interface{}) {
	ctx.JSON(http.StatusOK, response{
		Data:   data,
		Paging: paging,
	})
}

// parse system errors to response errors
func fromSystemErrors(errs appErrors.SystemErrors) []responseError {
	if len(errs) == 0 {
		return nil
	}

	respErrors := make([]responseError, len(errs))

	for i, e := range errs {
		respErrors[i] = responseError{
			Type:    string(e.Type()),
			Code:    string(e.Code()),
			Message: e.Message(),
			Param:   e.Param(),
		}
	}

	return respErrors
}

/*
traceErrors

set tag for errors into DataDog span
*/
func traceErrors(span tracer.Span, appErrs appErrors.SystemErrors) {
	if len(appErrs) == 0 {
		return
	}

	type err struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	errs := make([]err, len(appErrs))

	for i := range appErrs {
		errs[i] = err{
			Code:    string(appErrs[i].Code()),
			Message: appErrs[i].Message(),
		}
	}

	byteErrs, e := json.Marshal(errs)
	if e == nil {
		span.SetTag("errors", string(byteErrs))
	}
}
