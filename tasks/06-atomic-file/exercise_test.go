package task06

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFileAtomic(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.txt")
	if err := os.WriteFile(path, []byte("old"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := WriteFileAtomic(path, []byte("new state"), 0o640); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil || string(got) != "new state" {
		t.Fatalf("ReadFile = (%q, %v), want new state", got, err)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].Name() != "state.txt" {
		t.Fatalf("temporary files left behind: %v", entries)
	}
}

func TestWriteFileAtomicMissingDirectory(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing", "state.txt")
	if err := WriteFileAtomic(path, []byte("data"), 0o600); err == nil {
		t.Fatal("expected an error for a missing parent directory")
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("target unexpectedly exists: %v", err)
	}
}
