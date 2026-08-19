package task20

import (
	"context"
	"errors"
	"time"
)

var ErrNotStarted = errors.New("job was not started")

type Job struct {
	ID         string
	MaxRetries int
	Run        func(context.Context) error
}

type Result struct {
	ID       string
	Attempts int
	Err      error
}

func RunJobs(ctx context.Context, jobs []Job, workers int, attemptTimeout time.Duration) ([]Result, error) {
	// TODO: implement.
	return nil, nil
}
