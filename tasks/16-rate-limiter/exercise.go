package task16

import (
	"context"
	"errors"
	"time"
)

var ErrInvalidInterval = errors.New("interval must be positive")

type Limiter struct {
	// TODO: add synchronization and scheduling state.
	interval time.Duration
}

func NewLimiter(interval time.Duration) (*Limiter, error) {
	// TODO: implement.
	return nil, nil
}

func (l *Limiter) Wait(ctx context.Context) error {
	// TODO: implement.
	return nil
}
