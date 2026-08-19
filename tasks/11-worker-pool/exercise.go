package task11

import (
	"context"
	"errors"
)

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap(ctx context.Context, workers int, values []int, fn func(context.Context, int) (int, error)) ([]int, error) {
	// TODO: implement.
	return nil, nil
}
