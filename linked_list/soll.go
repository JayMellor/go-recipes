package linked_list

type selfOrderingLinkedList[E comparable] struct {
	linkedList[E]
	frequency map[E]int
}

func (soll *selfOrderingLinkedList[E]) Push(value E) (count int) {
	if soll.frequency == nil {
		soll.frequency = make(map[E]int)
	}
	soll.frequency[value] = 0
	return soll.linkedList.Push(value)
}

func (soll *selfOrderingLinkedList[E]) Contains(value E) bool {
	found := soll.linkedList.Contains(value)
	if found {
		soll.Pop(value)
		soll.Push(value)
		soll.frequency[value]++
	}
	return found
}

func (soll *selfOrderingLinkedList[E]) Frequency(value E) (frequency int, found bool) {
	found = soll.linkedList.Contains(value)
	if !found {
		return 0, false
	}
	return soll.frequency[value], true
}
