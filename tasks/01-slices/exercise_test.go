package task01

import (
	"reflect"
	"slices"
	"testing"
)

func TestUniqueStable(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{"nil", nil, nil},
		{"empty", []string{}, []string{}},
		{"stable", []string{"go", "py", "go", "rust", "py"}, []string{"go", "py", "rust"}},
		{"case-sensitive", []string{"Go", "go", "Go"}, []string{"Go", "go"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := slices.Clone(tt.in)
			got := UniqueStable(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("UniqueStable(%q) = %#v, want %#v", tt.in, got, tt.want)
			}
			if !reflect.DeepEqual(tt.in, input) {
				t.Fatalf("input was modified: got %q, original %q", tt.in, input)
			}
			if len(got) > 0 {
				got[0] += "-changed-by-test"
				if !reflect.DeepEqual(tt.in, input) {
					t.Fatalf("result reuses input storage: input changed to %q", tt.in)
				}
			}
		})
	}
}
