package profiler

var profile Profiler = NopProfiler{}

type (
	// Profiler interface
	Profiler interface {
		Start() error
		Stop()
	}
)

// SetProfiler overwrite the default profiler by another
func SetProfiler(p Profiler) {
	profile = p
}

// Start profiler
func Start() error {
	return profile.Start()
}

// Stop profiler
func Stop() {
	profile.Stop()
}
