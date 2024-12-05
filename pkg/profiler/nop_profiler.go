package profiler

// NopProfiler do nothing to help run local
type NopProfiler struct{}

// Start do nothing
func (t NopProfiler) Start() error {
	return nil
}

// Stop do nothing
func (t NopProfiler) Stop() {}
