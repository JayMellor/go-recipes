package linked_list

import "testing"

func TestLinkedListPush(t *testing.T) {
	list := linkedList[int]{}
	if ret := list.Push(1); ret != 1 {
		t.Error("Expected 1 returned from Push. Recevied", ret)
	}
	if ret := list.Push(3); ret != 2 {
		t.Error("Expected 2 returned from Push. Recevied", ret)
	}
	if list.length != 2 {
		t.Error("Expected list to have length 2. Recevied", list.length)
	}
	if list.head.value != 3 {
		t.Error("Expected head value to be 3. Received", list.head.value)
	}
	if list.head.next == nil {
		t.Error("Expected head to have next", list.head)
	}
	if list.head.next.value != 1 {
		t.Error("Expected head next to have value 1. Received", list.head.next)
	}
	if list.head.next.next != nil {
		t.Error("Expected head next to not have a next. Received", list.head.next)
	}
}

func TestLinkedListPop(t *testing.T) {
	list := linkedList[int]{}
	list.Push(3)
	list.Push(4)
	list.Push(5)
	list.Push(6)

	removed, found := list.Pop(7)
	if found || removed != nil {
		t.Error("Expected 7 not to be found. Received", removed, found)
	}

	removed, found = list.Pop(5)
	if !found || removed.value != 5 {
		t.Error("Expected node 5 to be removed. Received", removed, found)
	}
	if list.length != 3 {
		t.Error("Expected list to have length 3. Received", list)
	}
	if list.head.value != 6 {
		t.Error("Expected head value of 6. Received", list.head)
	}
	if list.head.next == nil {
		t.Error("Expected head to have next. Received", list.head)
	}
	if list.head.next.value != 4 {
		t.Error("Expected head next to have value 4. Received", list.head.next)
	}

	removed, found = list.Pop(6)
	if !found || removed.value != 6 {
		t.Error("Expected node 6 to be removed. Received", removed, found)
	}
	if list.length != 2 {
		t.Error("Expected list to have length 2. Received", list)
	}
	if list.head.value != 4 {
		t.Error("Expected head to have value 4. Received", list.head)
	}

	removed, found = list.Pop(3)
	if !found || removed.value != 3 {
		t.Error("Expected node 3 to be removed. Received", removed, found)
	}
	if list.length != 1 {
		t.Error("Expected list to have length 1. Received", list)
	}
	if list.head.value != 4 {
		t.Error("Expected head to have value 4. Received", list.head)
	}
	if list.head.next != nil {
		t.Error("Expected head to not have next. Received", list.head)
	}
}

func TestLinkedListToString(t *testing.T) {
	list := linkedList[string]{}
	list.Push("Elden Ring")
	list.Push("Dark Souls")
	list.Push("Armored Core")

	if toString := list.ToString(); toString != "Armored Core->Dark Souls->Elden Ring" {
		t.Error("Expected Armored Core->Dark Souls->Elden Ring. Received", toString)
	}
}
