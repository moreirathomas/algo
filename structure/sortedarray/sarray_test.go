package sortedarray

import (
	"reflect"
	"testing"
)

func TestHeapInsert(t *testing.T) {
	s := &SortedArray{}

	s.Insert(2)
	s.Insert(0)
	s.Insert(1)

	testOrder(t, []int{0, 1, 2}, s.Values())

	s.Insert(10)
	testOrder(t, []int{0, 1, 2, 10}, s.Values())
}

func testOrder(t *testing.T, want []int, got []int) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("bad insertion order, want %v, got %v", want, got)
	}
}
