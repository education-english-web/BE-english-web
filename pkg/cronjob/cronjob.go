package cronjob

//go:generate mockgen -destination=./mock/$GOFILE -source=$GOFILE -package=mock

// Scheduler run the scheduled jobs
type Scheduler interface {
	Run(jobName string) error
	Start()
	Stop()
	ScheduledJobs() map[string][]string
}

// ScheduledJob Job with schedules
type ScheduledJob interface {
	Schedules() []string
	Job() Job
}

// Job is the execution
type Job interface {
	Name() string
	Run()
}
