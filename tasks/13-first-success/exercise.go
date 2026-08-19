package task13

import (
	"context"
	"errors"
)

var ErrNoOperations = errors.New("no operations")

type Operation func(context.Context) (string, error)

func FirstSuccess(ctx context.Context, operations ...Operation) (string, error) {
	// TODO: implement.
	return "", nil
}
