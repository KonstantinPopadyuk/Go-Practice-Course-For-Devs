package task10

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFetchJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"gopher"}`))
	}))
	defer srv.Close()
	var got struct {
		Name string `json:"name"`
	}
	if err := FetchJSON(context.Background(), srv.Client(), srv.URL, 1024, &got); err != nil {
		t.Fatal(err)
	}
	if got.Name != "gopher" {
		t.Fatalf("decoded name = %q", got.Name)
	}
}

func TestFetchJSONStatusAndLimit(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/large" {
			_, _ = w.Write([]byte(strings.Repeat("x", 50)))
			return
		}
		http.Error(w, "backend is sad", http.StatusBadGateway)
	}))
	defer srv.Close()
	var dst any
	if err := FetchJSON(context.Background(), srv.Client(), srv.URL+"/large", 10, &dst); !errors.Is(err, ErrResponseTooLarge) {
		t.Fatalf("large response error = %v", err)
	}
	if err := FetchJSON(context.Background(), srv.Client(), srv.URL+"/bad", 1024, &dst); err == nil || !strings.Contains(err.Error(), "502") {
		t.Fatalf("status error = %v", err)
	}
}

func TestFetchJSONCancellation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	}))
	defer srv.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	var dst any
	err := FetchJSON(ctx, srv.Client(), srv.URL, 1024, &dst)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("cancellation error = %v", err)
	}
}
