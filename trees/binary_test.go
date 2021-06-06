package trees_test

import (
	"testing"

	"github.com/moreirathomas/algo/trees"
	"github.com/moreirathomas/algo/util"
)

func TestAddNode(t *testing.T) {
	tree := trees.NewBinaryTreeRandomized()
	child := tree.AddNode(util.RandomNonZeroDigit(), trees.ChildLeft)

	if child != tree.Left {
		t.Error("expected child to be added to the tree")
	}
}

func TestRootNode(t *testing.T) {
	tree := trees.NewBinaryTreeRandomized()
	child := tree.AddNode(util.RandomNonZeroDigit(), trees.ChildLeft)
	grandChild := child.AddNode(util.RandomNonZeroDigit(), trees.ChildLeft)
	root := grandChild.Root

	if root == nil || root != tree {
		t.Errorf("expected to get root node: want %s, got %s", tree.String(), root.String())
	}
}

func TestSiblingNode(t *testing.T) {
	tree := trees.NewBinaryTreeRandomized()
	tree.AddNode(util.RandomNonZeroDigit(), trees.ChildLeft)
	tree.AddNode(util.RandomNonZeroDigit(), trees.ChildLeft)
	sibling := tree.Left.Sibling()

	if sibling != tree.Right {
		t.Errorf("expected to get sibling node: want %s, got %s", tree.Right.String(), sibling.String())
	}
}

func TestInvertTree(t *testing.T) {

	// Given this tree:
	// ⁠     4
	// ⁠   /  \
	// ⁠  2    7
	// ⁠ / \  / \
	// 1  3 6  9

	// The inverted tree shoudl be:
	// ⁠     4
	// ⁠   /  \
	// ⁠  7    2
	// ⁠ / \  / \
	// 9  6 3  1

	tree := trees.NewBinaryTree(4)

	leftChild := tree.AddNode(2, trees.ChildLeft)

	leftChild.AddNode(1, trees.ChildLeft)
	leftChild.AddNode(3, trees.ChildRight)

	rightChild := tree.AddNode(7, trees.ChildRight)
	rightChild.AddNode(6, trees.ChildLeft)
	rightChild.AddNode(9, trees.ChildRight)

	tree.Invert()

	badInversionError := func(node string, want, got int) {
		t.Errorf("expected tree to be inverted at node %s: want %d, got %d", node, want, got)
	}

	if tree.Left.Val != 7 {
		badInversionError("left child", 7, tree.Left.Val)
	}
	if tree.Left.Left.Val != 9 {
		badInversionError("left child's left child", 9, tree.Left.Left.Val)
	}
	if tree.Left.Right.Val != 6 {
		badInversionError("left child's right child", 6, tree.Left.Right.Val)
	}
	if tree.Right.Val != 2 {
		badInversionError("right child", 2, tree.Right.Val)
	}
	if tree.Right.Left.Val != 3 {
		badInversionError("right child's left child", 3, tree.Right.Left.Val)
	}
	if tree.Right.Right.Val != 1 {
		badInversionError("right child's right child", 1, tree.Right.Right.Val)
	}
}
