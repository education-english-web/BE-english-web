package datadog

import (
	"context"
	"runtime"
	"strings"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type ContextKey string

const (
	KeyTenantUID ContextKey = "tenant_uid"
	KeyOfficeID  ContextKey = "office_id"
)

type spanSource struct {
	parentFuncName string
	funcPath       string
	lineNumber     int
}

func StartSpanFromCtx(ctx context.Context) (tracer.Span, context.Context) {
	sourceInfo := getParentFuncCallerInfo()
	span, spanCtx := tracer.StartSpanFromContext(ctx, sourceInfo.parentFuncName)

	span.SetTag("source.filename", sourceInfo.parentFuncName)
	span.SetTag("source.func_path", sourceInfo.funcPath)
	span.SetTag("source.line_number", sourceInfo.lineNumber)

	officeID := ctx.Value(KeyOfficeID)
	if officeID != nil {
		span.SetTag("source.office_id", officeID)
	}

	tenantUID := ctx.Value(KeyTenantUID)
	if tenantUID != nil {
		span.SetTag("source.tenant_uid", tenantUID)
	}

	return span, spanCtx
}

func getParentFuncCallerInfo() spanSource {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	funcPath := frame.Func.Name()
	parentFuncName := funcPath

	lastDotIdx := strings.LastIndex(funcPath, "/")
	if lastDotIdx > 0 {
		parentFuncName = funcPath[lastDotIdx+1:]
	}

	return spanSource{
		parentFuncName: parentFuncName,
		funcPath:       funcPath,
		lineNumber:     frame.Line,
	}
}

func SetError(ctx context.Context, err error) {
	span, _ := tracer.SpanFromContext(ctx)
	span.SetTag("error", true)
	span.SetTag("error.message", err.Error())
}
