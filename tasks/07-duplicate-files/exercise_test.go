package task07

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	root := t.TempDir()
	paths := []string{filepath.Join(root, "a.txt"), filepath.Join(root, "nested", "b.txt")}
	if err := os.Mkdir(filepath.Dir(paths[1]), 0o755); err != nil {
		t.Fatal(err)
	}
	for _, p := range paths {
		if err := os.WriteFile(p, []byte("same content"), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, "unique.txt"), []byte("different!!!"), 0o600); err != nil {
		t.Fatal(err)
	}

	got, err := FindDuplicates(root, 5)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte("same content"))
	want := map[string][]string{hex.EncodeToString(sum[:]): paths}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("FindDuplicates = %#v, want %#v", got, want)
	}
}

func TestFindDuplicatesHonorsMinimumSize(t *testing.T) {
	root := t.TempDir()
	for _, name := range []string{"a", "b"} {
		if err := os.WriteFile(filepath.Join(root, name), []byte("x"), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	got, err := FindDuplicates(root, 2)
	if err != nil || len(got) != 0 {
		t.Fatalf("FindDuplicates = (%v, %v), want empty map", got, err)
	}
}
