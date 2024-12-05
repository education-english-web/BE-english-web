package profiler

import (
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

type (
	// DatadogConfig hold configuration
	DatadogConfig struct {
		Env            string
		ServiceName    string
		ServiceVersion string
	}
)

// NewDatadogProfiler create profiler from config
func NewDatadogProfiler(cfg *DatadogConfig) DatadogProfiler {
	return DatadogProfiler{cfg: cfg}
}

// DatadogProfiler wrapper for datadog profiler
type DatadogProfiler struct {
	cfg *DatadogConfig
}

// Start tracer
func (t DatadogProfiler) Start() error {
	return profiler.Start(
		profiler.WithEnv(t.cfg.Env),
		profiler.WithService(t.cfg.ServiceName),
		profiler.WithVersion(t.cfg.ServiceVersion),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			profiler.BlockProfile,
			profiler.MutexProfile,
			profiler.GoroutineProfile,
		),
	)
}

// Stop tracer
func (t DatadogProfiler) Stop() {
	profiler.Stop()
}
