package set

type set[E comparable] []E

func NewSet[E comparable]() set[E] {
	return make(set[E], 0)
}

func (s *set[E]) Insert(entry E) {
	if !s.Contains(entry) {
		*s = append(*s, entry)
	}
}

func (s set[E]) find(entry E) (place int, found bool) {
	for idx, element := range s {
		if element == entry {
			return idx, true
		}
	}
	return -1, false
}

func (s set[E]) Contains(entry E) bool {
	_, found := s.find(entry)
	return found
}

func (s *set[E]) Remove(entry E) bool {
	place, found := s.find(entry)
	if !found {
		return false
	}

	if len(*s) == 1 {
		*s = NewSet[E]()
	} else {
		(*s)[place] = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
	}
	return true
}

func Union[E comparable](fst set[E], scd set[E]) (union set[E]) {
	for _, entry := range fst {
		union.Insert(entry)
	}
	for _, entry := range scd {
		union.Insert(entry)
	}
	return union
}

func Intersection[E comparable](fst set[E], scd set[E]) (intersection set[E]) {
	for _, fstEntry := range fst {
		for _, scdEntry := range scd {
			if fstEntry == scdEntry {
				intersection.Insert(fstEntry)
			}
		}
	}
	return intersection
}
