package task10

import (
	"context"
	"errors"
	"net/http"
)

var ErrResponseTooLarge = errors.New("response too large")

func FetchJSON(ctx context.Context, client *http.Client, url string, maxBytes int64, dst any) error {
	// TODO: implement.
	return nil
}
