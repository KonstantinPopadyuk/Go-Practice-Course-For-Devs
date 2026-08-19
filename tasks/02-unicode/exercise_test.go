package task02

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	tests := map[string]string{
		"":       "",
		"hello":  "olleh",
		"Привет": "тевирП",
		"Go🙂世界":  "界世🙂oG",
	}
	for in, want := range tests {
		got := Reverse(in)
		if got != want {
			t.Errorf("Reverse(%q) = %q, want %q", in, got, want)
		}
		if !utf8.ValidString(got) {
			t.Errorf("Reverse(%q) returned invalid UTF-8", in)
		}
	}
}

func TestReverseTwice(t *testing.T) {
	in := strings.Repeat("аб🙂", 20)
	if got := Reverse(Reverse(in)); got != in {
		t.Fatalf("double reverse = %q, want original", got)
	}
}
