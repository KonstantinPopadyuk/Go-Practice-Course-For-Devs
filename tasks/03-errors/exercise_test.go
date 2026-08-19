package task03

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func TestParsePort(t *testing.T) {
	for _, raw := range []string{"1", "80", "65535"} {
		got, err := ParsePort(raw)
		if err != nil {
			t.Fatalf("ParsePort(%q): %v", raw, err)
		}
		want, _ := strconv.Atoi(raw)
		if got != want {
			t.Errorf("ParsePort(%q) = %d, want %d", raw, got, want)
		}
	}
}

func TestParsePortErrors(t *testing.T) {
	for _, raw := range []string{"0", "65536", "-1"} {
		_, err := ParsePort(raw)
		if !errors.Is(err, ErrOutOfRange) {
			t.Errorf("ParsePort(%q) error = %v, want ErrOutOfRange", raw, err)
		}
		if err == nil || !strings.Contains(err.Error(), raw) {
			t.Errorf("error %v does not contain input %q", err, raw)
		}
	}

	_, err := ParsePort("eighty")
	var numErr *strconv.NumError
	if !errors.As(err, &numErr) {
		t.Fatalf("syntax error %v does not wrap *strconv.NumError", err)
	}
	if !strings.Contains(err.Error(), "eighty") {
		t.Fatalf("error %v does not contain original input", err)
	}
}
