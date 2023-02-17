package set

import "testing"

func TestNewSet(t *testing.T) {
	set := NewSet[int]()
	if len(set) != 0 {
		t.Error("Expected empty set created. Received", set)
	}
}

func TestSetInsert(t *testing.T) {
	set := NewSet[string]()
	set.Insert("first")
	set.Insert("second")
	set.Insert("first")

	if len(set) != 2 {
		t.Error("Expected 2 entries in set. Receieved", len(set))
	}
	if !set.Contains("first") || !set.Contains("second") {
		t.Error("Expected 'first' and 'second' in set", set)
	}
}

func TestSetContains(t *testing.T) {
	set := NewSet[int]()
	set.Insert(34)
	set.Insert(7)

	if !set.Contains(34) {
		t.Error("Expected 34 to be found in set", set)
	}
	if set.Contains(96) {
		t.Error("Expected 96 not to be found in set", set)
	}
}

func TestSetRemove(t *testing.T) {
	set := NewSet[int]()
	set.Insert(23)
	set.Insert(46)

	if !set.Remove(23) || set.Contains(23) {
		t.Error("Expected 23 to be removed from set", set)
	}

	if set.Remove(96) {
		t.Error("Expected failure to remove 96 from set", set)
	}
}

func TestSetUnion(t *testing.T) {
	fst := NewSet[int]()
	fst.Insert(1)
	fst.Insert(33)

	scd := NewSet[int]()
	scd.Insert(1)
	scd.Insert(54)

	union := Union(fst, scd)
	if len(union) != 3 {
		t.Error("Expected union to have length 3. Receieved", union)
	}

	if !union.Contains(1) || !union.Contains(33) || !union.Contains(54) {
		t.Error("Expected union to contain 1, 33 and 54. Received,", union)
	}

}

func TestSetIntersection(t *testing.T) {
	fst := NewSet[[2]int]()
	fst.Insert([2]int{1, 2})
	fst.Insert([2]int{2, 3})

	scd := NewSet[[2]int]()
	scd.Insert([2]int{1, 2})
	scd.Insert([2]int{3, 4})

	intersection := Intersection(fst, scd)
	if len(intersection) != 1 {
		t.Error("Expected intersection to have 1 element. Received", intersection)
	}
	if !intersection.Contains([2]int{1, 2}) {
		t.Error("Expected intersection to contain 1. Received", intersection)
	}
}
