package task18

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestRunServesAndShutsDown(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	started := make(chan struct{})
	server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-started:
		default:
			close(started)
		}
		_, _ = io.WriteString(w, "ok")
	})}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- Run(ctx, server, listener, time.Second) }()

	client := &http.Client{Timeout: 300 * time.Millisecond}
	resp, err := client.Get("http://" + listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()
	<-started
	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Run: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Run did not finish after cancellation")
	}
}

func TestRunValidatesTimeout(t *testing.T) {
	if err := Run(context.Background(), &http.Server{}, nil, 0); !errors.Is(err, ErrInvalidShutdownTimeout) {
		t.Fatalf("error = %v", err)
	}
}
