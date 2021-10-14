package huffman

import (
	"testing"
)

func TestBuild(t *testing.T) {
	r1 := Build(nil)
	if r1 != nil {
		t.Errorf("expected nil, got %v", r1)
	}

	r2 := Build([]*Node{})
	if r2 != nil {
		t.Errorf("expected nil, got %v", r2)
	}

	leaves := []*Node{
		{Value: "a", Count: 20},
		{Value: "b", Count: 40},
		{Value: "c", Count: 10},
		{Value: "d", Count: 7},
		{Value: "e", Count: 8},
		{Value: "f", Count: 15},
	}

	// [a, b, c, d, e, f]
	// [c, d, U, e, f]
	// [U, V, e, f]
	// [e, W, f]
	// [X, f]
	// [Y]

	root := Build(leaves)
	println(root.String())
}
