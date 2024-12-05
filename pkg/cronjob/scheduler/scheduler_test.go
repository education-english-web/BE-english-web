package scheduler

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/robfig/cron/v3"

	"github.com/education-english-web/BE-english-web/app/config"
	"github.com/education-english-web/BE-english-web/pkg/cronjob"
	"github.com/education-english-web/BE-english-web/pkg/timeutil"
)

func Test_scheduledJob_Schedule(t *testing.T) {
	type fields struct {
		schedule []string
		job      cronjob.Job
	}

	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "ok",
			fields: fields{
				schedule: []string{"*/1 * * * *"},
				job:      nil,
			},
			want: []string{"*/1 * * * *"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sj := &scheduledJob{
				schedules: tt.fields.schedule,
				job:       tt.fields.job,
			}
			got := sj.Schedules()

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("scheduledJob.Schedules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scheduledJob_Job(t *testing.T) {
	type fields struct {
		schedule []string
		job      cronjob.Job
	}

	tests := []struct {
		name   string
		fields fields
		want   cronjob.Job
	}{
		{
			name: "Ok",
			fields: fields{
				schedule: []string{},
				job:      nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sj := &scheduledJob{
				schedules: tt.fields.schedule,
				job:       tt.fields.job,
			}

			if got := sj.Job(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scheduledJob.Job() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetup(t *testing.T) {
	var (
		c = cron.New(
			cron.WithLocation(timeutil.JST),
		)
		schedules = []string{
			"0 1 28 2 *",
			"0 1 29 2 *",
		}
	)

	type args struct {
		env           string
		jobConfig     cronjob.JobMapping
		availableJobs map[string]cronjob.Job
	}

	tests := []struct {
		name    string
		args    args
		want    cronjob.Scheduler
		wantErr bool
	}{
		{
			name: "Ok",
			args: args{
				env: config.ENVHeroku,
				jobConfig: map[string]cronjob.Config{
					"export": {
						Schedules:    schedules,
						Environments: nil,
					},
				},
				availableJobs: map[string]cronjob.Job{
					"export": nil,
				},
			},
			want: &scheduler{
				c: c,
				scheduledJobsMap: map[string]cronjob.ScheduledJob{
					"export": &scheduledJob{
						schedules: schedules,
						job:       nil,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "env is not supported",
			args: args{
				env: config.ENVDevelopment,
				jobConfig: map[string]cronjob.Config{
					"export": {
						Schedules:    schedules,
						Environments: []string{"production", "staging"},
					},
				},
				availableJobs: map[string]cronjob.Job{
					"export": nil,
				},
			},
			want: &scheduler{
				c:                c,
				scheduledJobsMap: map[string]cronjob.ScheduledJob{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Setup(tt.args.env, tt.args.jobConfig, tt.args.availableJobs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(scheduler{})); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func Test_scheduler_Run(t *testing.T) {
	type fields struct {
		c                *cron.Cron
		scheduledJobsMap map[string]cronjob.ScheduledJob
	}

	type args struct {
		jobName string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Ok",
			fields: fields{
				c: &cron.Cron{},
				scheduledJobsMap: map[string]cronjob.ScheduledJob{
					"jobName": nil,
				},
			},
			args: args{
				jobName: "jobName",
			},
			wantErr: nil,
		},
		{
			name: "job",
			fields: fields{
				c: &cron.Cron{},
				scheduledJobsMap: map[string]cronjob.ScheduledJob{
					"jobName": nil,
				},
			},
			args: args{
				jobName: "jobName1",
			},
			wantErr: fmt.Errorf("job %s not found", "jobName1"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scheduler{
				c:                tt.fields.c,
				scheduledJobsMap: tt.fields.scheduledJobsMap,
			}

			err := s.Run(tt.args.jobName)
			if err == nil && tt.wantErr != nil {
				t.Errorf("scheduler.Run() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("scheduler.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scheduler_Start(t *testing.T) {
	type fields struct {
		c                *cron.Cron
		scheduledJobsMap map[string]cronjob.ScheduledJob
	}

	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scheduler{
				c:                tt.fields.c,
				scheduledJobsMap: tt.fields.scheduledJobsMap,
			}

			s.Start()
		})
	}
}

func Test_scheduler_Stop(t *testing.T) {
	type fields struct {
		c                *cron.Cron
		scheduledJobsMap map[string]cronjob.ScheduledJob
	}

	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scheduler{
				c:                tt.fields.c,
				scheduledJobsMap: tt.fields.scheduledJobsMap,
			}

			s.Stop()
		})
	}
}
