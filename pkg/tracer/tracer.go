package tracer

var trace Tracer = NopTracer{}

type (
	// Tracer interface
	Tracer interface {
		Start()
		Stop()
	}

	Span interface {
		// SetTag sets a key/value pair as metadata on the span.
		SetTag(key string, value interface{})
	}
)

// SetTracer overwrite the default tracer by another
func SetTracer(t Tracer) {
	trace = t
}

// Start tracer
func Start() {
	trace.Start()
}

// Stop tracer
func Stop() {
	trace.Stop()
}
