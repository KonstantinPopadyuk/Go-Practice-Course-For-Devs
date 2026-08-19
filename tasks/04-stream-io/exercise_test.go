package task04

import (
	"errors"
	"strings"
	"testing"
)

var errRead = errors.New("read failed")

type failingReader struct{ sent bool }

func (r *failingReader) Read(p []byte) (int, error) {
	if !r.sent {
		r.sent = true
		copy(p, "first\n")
		return len("first\n"), nil
	}
	return 0, errRead
}

func TestCountLines(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"", 0}, {"\n", 1}, {"a\nb\n", 2}, {"a\nb", 2},
		{strings.Repeat("x", 70*1024) + "\nlast", 2},
	}
	for _, tt := range tests {
		got, err := CountLines(strings.NewReader(tt.in))
		if err != nil || got != tt.want {
			t.Errorf("CountLines(input of %d bytes) = (%d, %v), want (%d, nil)", len(tt.in), got, err, tt.want)
		}
	}
}

func TestCountLinesPreservesReadError(t *testing.T) {
	_, err := CountLines(&failingReader{})
	if !errors.Is(err, errRead) {
		t.Fatalf("error = %v, want wrapped errRead", err)
	}
}
