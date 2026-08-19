package task17

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMiddlewareRequestIDAndLog(t *testing.T) {
	var logs bytes.Buffer
	h := Middleware(log.New(&logs, "", 0), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := RequestID(r.Context()); got != "test-id" {
			t.Errorf("RequestID = %q", got)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	req := httptest.NewRequest(http.MethodPost, "/items?token=secret", nil)
	req.Header.Set("X-Request-ID", "test-id")
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated || rr.Header().Get("X-Request-ID") != "test-id" {
		t.Fatalf("response status/header = %d/%q", rr.Code, rr.Header().Get("X-Request-ID"))
	}
	line := logs.String()
	for _, part := range []string{"POST", "/items", "201", "test-id"} {
		if !strings.Contains(line, part) {
			t.Errorf("log %q does not contain %q", line, part)
		}
	}
	if strings.Contains(line, "secret") {
		t.Errorf("log leaked query string: %q", line)
	}
}

func TestMiddlewareRecoversAndGeneratesID(t *testing.T) {
	var logs bytes.Buffer
	h := Middleware(log.New(&logs, "", 0), http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		panic("boom")
	}))
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/panic", nil))
	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d", rr.Code)
	}
	if rr.Header().Get("X-Request-ID") == "" {
		t.Fatal("generated request ID is empty")
	}
	if !strings.Contains(logs.String(), "boom") {
		t.Fatalf("panic not logged: %q", logs.String())
	}
}

func TestMiddlewareKeepsCommittedStatusAfterPanic(t *testing.T) {
	var logs bytes.Buffer
	h := Middleware(log.New(&logs, "", 0), http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		panic("after write")
	}))
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/committed", nil))
	if rr.Code != http.StatusAccepted {
		t.Fatalf("status = %d, want already committed %d", rr.Code, http.StatusAccepted)
	}
	for _, part := range []string{"after write", "202"} {
		if !strings.Contains(logs.String(), part) {
			t.Errorf("log %q does not contain %q", logs.String(), part)
		}
	}
}

func TestMiddlewareGeneratesUniqueIDsConcurrently(t *testing.T) {
	h := Middleware(log.New(io.Discard, "", 0), http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	const requests = 32
	type response struct {
		id   string
		code int
	}
	responses := make(chan response, requests)
	for range requests {
		go func() {
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/generated", nil))
			responses <- response{id: rr.Header().Get("X-Request-ID"), code: rr.Code}
		}()
	}
	seen := make(map[string]struct{}, requests)
	for range requests {
		result := <-responses
		if result.code != http.StatusOK {
			t.Errorf("implicit status = %d, want 200", result.code)
		}
		if result.id == "" {
			t.Fatal("generated request ID is empty")
		}
		if _, exists := seen[result.id]; exists {
			t.Fatalf("duplicate generated request ID %q", result.id)
		}
		seen[result.id] = struct{}{}
	}
}
