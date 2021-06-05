package trees_test

import (
	"testing"

	"github.com/moreirathomas/algo/trees"
	"github.com/moreirathomas/algo/util"
)

func TestBinary(t *testing.T) {
	tree := trees.NewBinaryTree()

	// Left is two level deep.
	child := tree.AddNode(util.RandomNonZeroDigit(), trees.ChildLeft)
	if child != tree.Left {
		t.Errorf("expected child to be added to the tree")

	}

	grandChild := child.AddNode(util.RandomNonZeroDigit(), trees.ChildLeft)

	// Right is one level deep.
	tree.AddNode(util.RandomNonZeroDigit(), trees.ChildRight)

	root := grandChild.Root
	if root != tree {
		t.Errorf("expected root node be %s: got %s", tree.String(), root.String())
	}
}
