package task18

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

var ErrInvalidShutdownTimeout = errors.New("shutdown timeout must be positive")

func Run(ctx context.Context, server *http.Server, listener net.Listener, shutdownTimeout time.Duration) error {
	// TODO: implement.
	return nil
}
