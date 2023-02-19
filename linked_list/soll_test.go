package linked_list

import "testing"

func TestSelfOrderingAndFrequency(t *testing.T) {
	soll := selfOrderingLinkedList[int]{}
	soll.Push(3)
	soll.Push(4)
	soll.Push(-2)
	if soll.head.value != -2 {
		t.Error("Expect -2 at the head of list. Received", soll.head)
	}

	soll.Contains(3)
	if soll.head.value != 3 {
		t.Error("Expected 3 at the head of list. Received", soll.head)
	}

	if frequency, found := soll.Frequency(3); !found || frequency != 1 {
		t.Error("Expected 3 to be found with frequency 1. Received", frequency, found)
	}

	if frequency, found := soll.Frequency(-2); !found || frequency != 0 {
		t.Error("Expected -2 to be found with frequency 0. Received", frequency, found)
	}
}
