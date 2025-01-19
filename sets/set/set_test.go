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

func TestRemove(t *testing.T) {
	s := set.New[int]()

	// 1. Empty set
	s.Remove(10)
	if s.Size() != 0 {
		t.Errorf("Expected size 0, but got %d", s.Size())
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	// 2. Remove one value
	s.Remove(1)
	if s.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", s.Size())
	}
	if s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("Expected to only have 2 & 3")
	}

	// 3. Remove number that doesn't contain
	s.Remove(9999)
	if s.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", s.Size())
	}
	if s.Contains(9999) {
		t.Errorf("Expected not to have 9999")
	}
}
