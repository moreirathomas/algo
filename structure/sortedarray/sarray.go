package sortedarray

import (
	"errors"

	"github.com/moreirathomas/algo/sort"
)

// TODO make it generic or name it as specialized for ints.

type SortedArray struct {
	values []int
}

func (a *SortedArray) Insert(val int) {
	i := sort.Binary(a.values, val)

	a.values = append(a.values, 0)
	copy(a.values[i+1:], a.values[i:])
	a.values[i] = val
}

func (a SortedArray) Values() []int {
	return a.values
}

func (a SortedArray) At(i int) (int, error) {
	if i > len(a.values)-1 {
		return 0, errors.New("out of range")
	}
	return a.values[i], nil
}
