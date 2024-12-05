package worker

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

// Job that is submitted to the workers
type Job []byte

// Workers that runs the job
type Workers interface {
	Run(job Job)
}
