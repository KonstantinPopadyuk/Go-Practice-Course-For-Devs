package task05

import (
	"strings"
	"testing"
	"time"
)

func TestLoadConfig(t *testing.T) {
	got, err := LoadConfig(strings.NewReader(`{"address":"127.0.0.1:8080","workers":4,"timeout":"750ms"}`))
	if err != nil {
		t.Fatal(err)
	}
	if got.Address != "127.0.0.1:8080" || got.Workers != 4 || time.Duration(got.Timeout) != 750*time.Millisecond {
		t.Fatalf("unexpected config: %+v", got)
	}
}

func TestLoadConfigRejectsInvalidInput(t *testing.T) {
	inputs := []string{
		`{"address":"","workers":4,"timeout":"1s"}`,
		`{"address":"x","workers":0,"timeout":"1s"}`,
		`{"address":"x","workers":65,"timeout":"1s"}`,
		`{"address":"x","workers":1,"timeout":"0s"}`,
		`{"address":"x","workers":1,"timeout":"bad"}`,
		`{"address":"x","workers":1,"timeout":"1s","debug":true}`,
		`{"address":"x","workers":1,"timeout":"1s"} {}`,
	}
	for _, input := range inputs {
		if _, err := LoadConfig(strings.NewReader(input)); err == nil {
			t.Errorf("LoadConfig(%s) returned nil error", input)
		}
	}
}
