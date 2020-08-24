package dllist

import (
	"testing"
)

func TestDLList(t *testing.T) { // TestDLList is a function that tests working of our list
	a := NewList() // Creating a new list a
	x1 := 0
	if ct := a.Len(); ct != x1 { // Checking start length, if it isn't right it breaks test and returns information about a failed line of code
		t.Fatalf("Wrong length for new list: got %d expected %d", ct, x1)
	}
	a.PushFront(1) // Using PushFront method, we add new elements to front part of list
	a.PushFront(2)
	a.PushFront(3)
	e4 := a.PushFront(4)
	e5 := a.PushFront(5)
	x2 := 5
	if ct := a.Len(); ct != x2 { // Checking list length after edding five new elements, if it isn't right it breaks test and returns information about a failed line of code
		t.Fatalf("Wrong length for new list: got %d expected %d", ct, x2)
	}
	i := 5
	x3 := 0
	for e := a.First(); e != nil; e = e.Next() { // Checking list elements, if they aren't right it breaks test and returns information about a failed line of code
		if ct := e.Value; ct != i {
			t.Fatalf("Wrong value for new list in position %d: got %d expected %d", x3, ct, i)
		}
		i--
		x3++
	}
	a.Remove(e4) // Using Remove method, we remove two elements from the list
	a.Remove(e5)
	i = 3
	x3 = 0
	for e := a.First(); e != nil; e = e.Next() { // Checking list elements after removing two elements, if they aren't right it breaks test and returns information about a failed line of code
		if ct := e.Value; ct != i {
			t.Fatalf("Wrong value for new list in position %d: got %d expected %d", x3, ct, i)
		}
		i--
		x3++
	}
}
