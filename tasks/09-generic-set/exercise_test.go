package task09

import "testing"

func TestSetZeroValueAndOperations(t *testing.T) {
	var s Set[string]
	s.Add("go")
	s.Add("go")
	s.Add("python")
	if s.Len() != 2 || !s.Contains("go") || s.Contains("rust") {
		t.Fatalf("unexpected set state: len=%d values=%v", s.Len(), s.Values())
	}
	s.Remove("go")
	if s.Contains("go") || s.Len() != 1 {
		t.Fatalf("Remove failed: %v", s.Values())
	}
}

func TestSetCloneIsIndependent(t *testing.T) {
	var original Set[int]
	original.Add(1)
	clone := original.Clone()
	clone.Add(2)
	original.Remove(1)
	if original.Contains(2) || !clone.Contains(1) || !clone.Contains(2) {
		t.Fatalf("clone shares state: original=%v clone=%v", original.Values(), clone.Values())
	}
}

func TestSetSupportsComparableStruct(t *testing.T) {
	type key struct{ A, B int }
	var s Set[key]
	s.Add(key{1, 2})
	if !s.Contains(key{1, 2}) {
		t.Fatal("struct key not found")
	}
}
