package tracer

import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type (
	// DatadogConfig hold configuration
	DatadogConfig struct {
		Env            string
		ServiceName    string
		ServiceVersion string
		Host           string
		AgentAMPPort   string
	}
)

// NewDatadogTracer create tracer from config
func NewDatadogTracer(cfg *DatadogConfig) DatadogTracer {
	return DatadogTracer{cfg: cfg}
}

// DatadogTracer wrapper for datadog tracer
type DatadogTracer struct {
	cfg *DatadogConfig
}

// Start tracer
func (t DatadogTracer) Start() {
	tracer.Start(
		tracer.WithEnv(t.cfg.Env),
		tracer.WithService(t.cfg.ServiceName),
		tracer.WithServiceVersion(t.cfg.ServiceVersion),
		tracer.WithAgentAddr(t.cfg.Host+":"+t.cfg.AgentAMPPort),
		tracer.WithAnalytics(true),
		tracer.WithRuntimeMetrics(),
	)
}

// Stop tracer
func (t DatadogTracer) Stop() {
	tracer.Stop()
}
