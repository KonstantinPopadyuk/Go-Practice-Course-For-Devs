package task17

import (
	"context"
	"log"
	"net/http"
)

type requestIDKey struct{}

func RequestID(ctx context.Context) string {
	value, _ := ctx.Value(requestIDKey{}).(string)
	return value
}

func Middleware(logger *log.Logger, next http.Handler) http.Handler {
	// TODO: implement.
	return next
}
