package workerpool

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/education-english-web/BE-english-web/pkg/worker"
	"github.com/education-english-web/BE-english-web/pkg/worker/workerpool/mock"
)

func TestNew(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mJobManager := mock.NewMockJobManager(mockCtrl)
	got := New(10, mJobManager)

	expected := &workerPool{
		jobsQueue:  make(chan worker.Job),
		jobManager: mJobManager,
	}

	if diff := cmp.Diff(
		expected,
		got,
		cmpopts.IgnoreFields(workerPool{}, "jobsQueue", "jobManager"),
	); diff != "" {
		t.Errorf(diff)
	}
}
