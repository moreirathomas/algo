package trees

import (
	"encoding/json"

	"github.com/moreirathomas/algo/util"
)

type treeNode struct {
	Val    int       `json:"value"`
	Parent *treeNode `json:"-"`
	Left   *treeNode `json:"left,omitempty"`
	Right  *treeNode `json:"right,omitempty"`
	Root   *treeNode `json:"-"`
}

type child int

const (
	ChildLeft child = iota
	ChildRight
)

// New returns a new binary tree given the root node's value.
func NewBinaryTree(rootVal int) *treeNode {
	return &treeNode{Val: rootVal}
}

// New returns a new binary tree with a randomized initial value.
func NewBinaryTreeRandomized() *treeNode {
	return &treeNode{Val: util.RandomNonZeroDigit()}
}

// AddNode adds a new child to a node. It returns the node just added.
func (node *treeNode) AddNode(val int, child child) *treeNode {
	nextNode := newNode(val, node)
	switch child {
	case ChildLeft:
		node.Left = nextNode
	case ChildRight:
		node.Right = nextNode
	}
	return nextNode
}

// newNode returns a new node given a value and its parent.
// The root node is inferred by the parent.
func newNode(val int, parent *treeNode) *treeNode {
	// Infer the root node. If there is no root is defined on the parent, it is the root.
	root := parent.Root
	if root == nil {
		root = parent
	}

	return &treeNode{
		Val:    val,
		Parent: parent,
		Left:   nil,
		Right:  nil,
		Root:   root,
	}
}

func (node *treeNode) Sibling() *treeNode {
	// eh.
	if node.Parent == nil {
		return nil
	} else if node.Parent.Left == node {
		return node.Parent.Right
	} else {
		return node
	}
}

// Invert inverts a tree in place.
func (tree *treeNode) Invert() {
	invert(tree)
}

func invert(node *treeNode) {
	if node != nil {
		node.Left, node.Right = node.Right, node.Left
		invert(node.Left)
		invert(node.Right)
	}
}

// String returns the JSON stringified version of a tree.
func (tree treeNode) String() string {
	str, err := json.Marshal(tree)
	if err != nil {
		panic(err)
	}
	return string(str)
}
