package set_test

import (
	"go-dsa/set"
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

func TestContains(t *testing.T) {
	s := set.New[int]()

	// 1. Empty
	if s.Contains(1) {
		t.Errorf("Expected to not contain 1 on an empty set")
	}

	// 2. Set with size 1
	s.Add(10)
	if !s.Contains(10) {
		t.Errorf("Expected to have 10")
	}
}

func TestUnion(t *testing.T) {
	s1 := set.New[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := set.New[int]()
	s2.Add(4)
	s2.Add(5)
	s2.Add(6)

	union := s1.Union(s2)
	expected := []int{1, 2, 3, 4, 5, 6}

	if union.Size() != len(expected) {
		t.Errorf("Mismatch size, union size=%d and expected size=%d", union.Size(), len(expected))
	}
	for _, v := range expected {
		if !union.Contains(v) {
			t.Errorf("Expected union to contain %d, but it does not", v)
		}
	}
}

func TestDifference(t *testing.T) {
	s1 := set.New[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s1.Add(4)

	s2 := set.New[int]()
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)

	difference := s1.Difference(s2)
	expected := []int{1, 2}

	if difference.Size() != len(expected) {
		t.Errorf("Expected difference size=2, but got %d", difference.Size())
	}
	for _, v := range expected {
		if !difference.Contains(v) {
			t.Errorf("Expected difference to contain %d, but it does not", v)
		}
	}
}

func TestEquals(t *testing.T) {
	// 1. Same numbers in the same order
	s1 := set.New[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := set.New[int]()
	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	if !s1.Equals(s2) {
		t.Errorf("Expected s1 and s2 to be equal, but they are not")
	}

	// 2. Same numbers in a different order
	s2 = set.New[int]()
	s2.Add(3)
	s2.Add(1)
	s2.Add(2)

	if !s1.Equals(s2) {
		t.Errorf("Expected s1 and s2 to be equal, but they are not")
	}

	// 3. Different numbers
	s2 = set.New[int]()
	s2.Add(10)
	s2.Add(20)
	s2.Add(30)

	if s1.Equals(s2) {
		t.Errorf("Expected s1 and s2 to be not equal, but they are")
	}
}
