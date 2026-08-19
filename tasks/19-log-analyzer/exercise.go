package task19

import (
	"context"
	"errors"
)

var ErrInvalidWorkers = errors.New("workers must be positive")

type Report struct {
	Lines        int
	ByLevel      map[string]int
	MaxLatencyMS int
}

func Analyze(ctx context.Context, paths []string, workers int) (Report, error) {
	// TODO: implement.
	return Report{}, nil
}
