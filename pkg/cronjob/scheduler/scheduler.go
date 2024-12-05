package scheduler

import (
	"fmt"
	"os"
	"os/signal"
	"slices"
	"syscall"

	"github.com/robfig/cron/v3"

	"github.com/education-english-web/BE-english-web/pkg/cronjob"
	"github.com/education-english-web/BE-english-web/pkg/timeutil"
)

type scheduledJob struct {
	schedules []string
	job       cronjob.Job
}

func (sj *scheduledJob) Schedules() []string {
	return sj.schedules
}

func (sj *scheduledJob) Job() cronjob.Job {
	return sj.job
}

type scheduler struct {
	c                *cron.Cron
	scheduledJobsMap map[string]cronjob.ScheduledJob // used to run job manually
}

func Setup(env string, jobConfig cronjob.JobMapping, availableJobs map[string]cronjob.Job) (cronjob.Scheduler, error) {
	c := cron.New(
		cron.WithLocation(timeutil.JST),
	)
	scheduledJobsMap := make(map[string]cronjob.ScheduledJob)

	// build scheduledJobsMap
	for jobName, config := range jobConfig {
		if config.Environments != nil && env != "" && !slices.Contains(config.Environments, env) {
			// skip if environment is not in the list
			continue
		}

		job, ok := availableJobs[jobName]
		if !ok {
			continue
		}

		for _, schedule := range config.Schedules {
			_, err := c.AddJob(schedule, job)
			if err != nil {
				return nil, err
			}
		}

		scheduledJobsMap[jobName] = &scheduledJob{
			schedules: config.Schedules,
			job:       job,
		}
	}

	return &scheduler{
		c:                c,
		scheduledJobsMap: scheduledJobsMap,
	}, nil
}

func (s *scheduler) Run(jobName string) error {
	var err error

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error running job: %v", r)
		}
	}()

	job, ok := s.scheduledJobsMap[jobName]
	if !ok {
		return fmt.Errorf("job %s not found", jobName)
	}

	job.Job().Run()

	return err
}

func (s *scheduler) ScheduledJobs() map[string][]string {
	scheduledJobs := make(map[string][]string)

	for name, sj := range s.scheduledJobsMap {
		scheduledJobs[name] = sj.Schedules()
	}

	return scheduledJobs
}

func (s *scheduler) Start() {
	stopSignal := make(chan os.Signal, 1)

	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)
	s.c.Start()
	<-stopSignal
	s.c.Stop()
}

func (s *scheduler) Stop() {
	s.c.Stop()
}
