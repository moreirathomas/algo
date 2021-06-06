package trees_test

import (
	"testing"

	"github.com/moreirathomas/algo/trees"
	"github.com/moreirathomas/algo/util"
)

// This file uses a shared tree for all its tests.
var valuesDepth2 = []int{4, 2, 7, 1, 3, 6, 9}

// It is represented as this:
// ⁠     4
// ⁠   /  \
// ⁠  2    7
// ⁠ / \  / \
// 1  3 6  9

// In its inverted form:
// ⁠     4
// ⁠   /  \
// ⁠  7    2
// ⁠ / \  / \
// 9  6 3  1

// A simpler tree with only a depth of 1.
var valuesDepth1 = valuesDepth2[:3]

func TestNewTree(t *testing.T) {
	tree := trees.NewBinaryTree(valuesDepth2...)

	badNodeValueErrorFmt := func(node string, want, got int) {
		t.Errorf("expected tree to have correct at %s node: want %d, got %d", node, want, got)
	}

	if tree.Val != 4 {
		badNodeValueErrorFmt("root", 4, tree.Val)
	}
	if tree.Left.Val != 2 {
		badNodeValueErrorFmt("left child", 2, tree.Left.Val)
	}
	if tree.Left.Left.Val != 1 {
		badNodeValueErrorFmt("left child's left child", 1, tree.Left.Left.Val)
	}
	if tree.Left.Right.Val != 3 {
		badNodeValueErrorFmt("left child's right child", 3, tree.Left.Right.Val)
	}
	if tree.Right.Val != 7 {
		badNodeValueErrorFmt("right child", 7, tree.Right.Val)
	}
	if tree.Right.Left.Val != 6 {
		badNodeValueErrorFmt("right child's left child", 6, tree.Right.Left.Val)
	}
	if tree.Right.Right.Val != 9 {
		badNodeValueErrorFmt("right child's right child", 9, tree.Right.Right.Val)
	}
}

func TestInsertNode(t *testing.T) {
	tree := trees.NewBinaryTree(util.RandomNonZeroDigit())
	node := trees.NewNode(util.RandomNonZeroDigit(), tree)
	tree.InsertNode(node)

	insertionErrorFmt := func(where string) {
		t.Errorf("expected node to be inserted to the %s: got tree %s, want tree with node %s",
			where, tree.String(), node.String())
	}

	switch {
	case node.Val < tree.Val:
		if node != tree.Left {
			insertionErrorFmt("left")
		}
	case node.Val > tree.Val:
		if node != tree.Right {
			insertionErrorFmt("right")
		}
	default:
		t.Errorf("unexpected error: unable to compare %v and %v", tree, node)
	}
}

func TestRootNode(t *testing.T) {
	tree := trees.NewBinaryTree(valuesDepth2...)
	root := tree.Left.Left.Root

	if root == nil || root != tree {
		t.Errorf("expected to get root node: want %s, got %s", tree.String(), root.String())
	}
}

func TestSiblingNode(t *testing.T) {
	tree := trees.NewBinaryTree(valuesDepth1...)
	sibling := tree.Left.Sibling()

	if sibling != tree.Right {
		t.Errorf("expected to get sibling node: want %s, got %s", tree.Right.String(), sibling.String())
	}
}

func TestInvertTree(t *testing.T) {

	tree := trees.NewBinaryTree(valuesDepth2...)

	tree.Invert()

	badNodeValueErrorFmt := func(node string, want, got int) {
		t.Errorf("expected tree to be inverted at node %s: want %d, got %d", node, want, got)
	}

	if tree.Left.Val != 7 {
		badNodeValueErrorFmt("left child", 7, tree.Left.Val)
	}
	if tree.Left.Left.Val != 9 {
		badNodeValueErrorFmt("left child's left child", 9, tree.Left.Left.Val)
	}
	if tree.Left.Right.Val != 6 {
		badNodeValueErrorFmt("left child's right child", 6, tree.Left.Right.Val)
	}
	if tree.Right.Val != 2 {
		badNodeValueErrorFmt("right child", 2, tree.Right.Val)
	}
	if tree.Right.Left.Val != 3 {
		badNodeValueErrorFmt("right child's left child", 3, tree.Right.Left.Val)
	}
	if tree.Right.Right.Val != 1 {
		badNodeValueErrorFmt("right child's right child", 1, tree.Right.Right.Val)
	}
}
