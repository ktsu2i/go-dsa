package set_test

import (
	"go-dsa/sets/set"
	"testing"
)

func TestAdd(t *testing.T) {
	s := set.New[int]()

	// 1. Add one value
	s.Add(1)
	if s.Size() != 1 {
		t.Errorf("Expected size 1, but got %d", s.Size())
	}
	if !s.Contains(1) {
		t.Errorf("Expected to have 1")
	}

	// 2. Add two values
	s.Add(2)
	if s.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", s.Size())
	}
	if !s.Contains(2) {
		t.Errorf("Expected to have 2")
	}

	// 3. Add same number
	s.Add(1)
	if s.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", s.Size())
	}
	if !s.Contains(1) {
		t.Errorf("Expected to have 1")
	}
}
