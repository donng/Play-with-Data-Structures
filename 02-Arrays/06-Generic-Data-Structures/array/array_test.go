package array

import "testing"

func TestRemoveAllElementAdjacentMatches(t *testing.T) {
	values := []int{1, 2, 2, 3, 2, 2}
	a := New(8)
	for _, value := range values {
		a.AddLast(value)
	}

	if !a.RemoveAllElement(2) {
		t.Fatal("RemoveAllElement(2) returned false")
	}
	if got := a.GetSize(); got != 2 {
		t.Fatalf("size = %d, want 2", got)
	}
	for index, want := range []int{1, 3} {
		if got := a.Get(index); got != want {
			t.Errorf("element %d = %v, want %d", index, got, want)
		}
	}
	if got := a.GetCapacity(); got != 8 {
		t.Errorf("capacity = %d, want 8", got)
	}
	for index := a.GetSize(); index < len(values); index++ {
		if a.data[index] != nil {
			t.Errorf("vacated slot %d was not cleared", index)
		}
	}
	if a.RemoveAllElement(2) {
		t.Error("RemoveAllElement(2) returned true after all matches were removed")
	}
}

func TestRemoveAllElementKeepsUnprocessedSuffixOnComparePanic(t *testing.T) {
	a := New(5)
	for _, value := range []interface{}{1, 2, 3, "incompatible", 4} {
		a.AddLast(value)
	}

	defer func() {
		if recover() != "cannot compare different type params" {
			t.Fatal("expected comparison of different types to panic")
		}
		want := []interface{}{1, 3, "incompatible", 4}
		if got := a.GetSize(); got != len(want) {
			t.Fatalf("size = %d, want %d", got, len(want))
		}
		for index, value := range want {
			if got := a.Get(index); got != value {
				t.Errorf("element %d = %v, want %v", index, got, value)
			}
		}
		if a.data[len(want)] != nil {
			t.Error("vacated slot was not cleared")
		}
	}()

	a.RemoveAllElement(2)
}
