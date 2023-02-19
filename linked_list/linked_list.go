package linked_list

import "fmt"

type linkedList[E comparable] struct {
	head   *node[E]
	length int
}

type node[E comparable] struct {
	value E
	next  *node[E]
}

func (list *linkedList[E]) Push(value E) (count int) {
	newNode := node[E]{value: value, next: list.head}
	list.head = &newNode
	list.length++
	return list.length
}

// Not sure whether this helps or hinders
func (list *linkedList[E]) iterate(cb func(current *node[E]) (stop bool)) {
	for link := list.head; link != nil; link = link.next {
		if stop := cb(link); stop {
			return
		}
	}
}

func (list *linkedList[E]) find(value E) (link *node[E], found bool) {
	list.iterate(func(current *node[E]) (stop bool) {
		if current.value == value {
			link = current
			return true
		}
		return false
	})
	return link, link != nil
}

func (list *linkedList[E]) Pop(value E) (removed *node[E], found bool) {
	var prev *node[E]
	list.iterate(func(current *node[E]) (stop bool) {
		if current.value == value {
			removed = current
			return true
		} else {
			prev = current
			return false
		}
	})
	if removed == nil {
		return nil, false
	}

	if prev == nil {
		list.head = removed.next
	} else {
		prev.next = removed.next
	}
	list.length--
	return removed, true
}

func (list *linkedList[E]) Contains(value E) bool {
	_, found := list.find(value)
	return found
}

func (list *linkedList[E]) ToString() (toString string) {
	list.iterate(func(current *node[E]) (stop bool) {
		toString += fmt.Sprint(current.value)
		if current.next != nil {
			toString += "->"
		}
		return false
	})
	return toString
}
