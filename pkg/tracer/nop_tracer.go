package tracer

// NopTracer do nothing to help run local
type NopTracer struct{}

// Start do nothing
func (t NopTracer) Start() {}

// Stop do nothing
func (t NopTracer) Stop() {}

type NoopSpan struct{}

// SetTag do nothing
func (t NoopSpan) SetTag(_ string, _ interface{}) {
	// no operation
}
