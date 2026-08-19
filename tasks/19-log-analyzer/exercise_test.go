package task19

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestAnalyze(t *testing.T) {
	dir := t.TempDir()
	paths := []string{filepath.Join(dir, "a.jsonl"), filepath.Join(dir, "b.jsonl")}
	contents := []string{
		"{\"level\":\"INFO\",\"latency_ms\":12}\n\n{\"level\":\"ERROR\",\"latency_ms\":80}\n",
		"{\"level\":\"INFO\",\"latency_ms\":25}\n",
	}
	for i := range paths {
		if err := os.WriteFile(paths[i], []byte(contents[i]), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	got, err := Analyze(context.Background(), paths, 2)
	if err != nil {
		t.Fatal(err)
	}
	want := Report{Lines: 3, ByLevel: map[string]int{"INFO": 2, "ERROR": 1}, MaxLatencyMS: 80}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Analyze = %#v, want %#v", got, want)
	}
}

func TestAnalyzeReportsFileAndLine(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.jsonl")
	if err := os.WriteFile(path, []byte("{\"level\":\"INFO\",\"latency_ms\":1}\nnot-json\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := Analyze(context.Background(), []string{path}, 1)
	messageWithoutPath := ""
	if err != nil {
		messageWithoutPath = strings.ReplaceAll(err.Error(), path, "")
	}
	if err == nil || !strings.Contains(err.Error(), path) || !strings.Contains(messageWithoutPath, "2") {
		t.Fatalf("error = %v, want path and line 2", err)
	}
	if !reflect.DeepEqual(got, Report{}) {
		t.Fatalf("report on error = %#v, want zero value", got)
	}
}

func TestAnalyzeValidatesWorkersBeforeOpeningFiles(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "must-not-be-opened.jsonl")
	got, err := Analyze(context.Background(), []string{missing}, 0)
	if !errors.Is(err, ErrInvalidWorkers) {
		t.Fatalf("error = %v, want ErrInvalidWorkers", err)
	}
	if !reflect.DeepEqual(got, Report{}) {
		t.Fatalf("report = %#v, want zero value", got)
	}
}

func TestAnalyzeSupportsLongLines(t *testing.T) {
	level := strings.Repeat("X", 70*1024)
	path := filepath.Join(t.TempDir(), "long.jsonl")
	content := "{\"level\":\"" + level + "\",\"latency_ms\":7}\n"
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := Analyze(context.Background(), []string{path}, 1)
	if err != nil {
		t.Fatal(err)
	}
	if got.Lines != 1 || got.ByLevel[level] != 1 || got.MaxLatencyMS != 7 {
		t.Fatalf("Analyze(long line) = %#v", got)
	}
}
