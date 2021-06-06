package trees

import (
	"encoding/json"
)

type treeNode struct {
	Val    int       `json:"value"`
	Left   *treeNode `json:"left,omitempty"`
	Right  *treeNode `json:"right,omitempty"`
	parent *treeNode `json:"-"`
	root   *treeNode `json:"-"`
}

// NewBinaryTree returns a binary tree. The first value passed is always the root.
// All the others are inserted to the left or right depending on their value compared
// to the previous node.
func NewBinaryTree(values ...int) *treeNode {
	root := NewNode(values[0], nil)
	for _, v := range values[1:] {
		node := NewNode(v, root)
		root.InsertNode(node)
	}
	return root
}

// InsertNode inserts the given node on the tree. A node whose value is greater
// than the previous one is inserted to the right, and to the left otherwise.
func (tree *treeNode) InsertNode(node *treeNode) {
	if tree == nil {
		panic("cannot add a node to a nil parent")
	}
	if node.Val < tree.Val {
		if tree.Left == nil {
			tree.Left = node
		}
		tree.Left.InsertNode(node)
	}
	if node.Val > tree.Val {
		if tree.Right == nil {
			tree.Right = node
		}
		tree.Right.InsertNode(node)
	}
}

// NewNode returns a new node given a value and its parent.
// The root node is inferred by the parent.
func NewNode(val int, parent *treeNode) *treeNode {
	// Infer the root node: either this is the root,
	// or the parent has no root and thus is the root, or it knows the root.
	root := parent
	if parent != nil && parent.root != nil {
		root = parent.root
	}
	return &treeNode{Val: val, parent: parent, root: root}
}

// Invert inverts a tree in place.
func (tree *treeNode) Invert() {
	if tree != nil {
		tree.Left, tree.Right = tree.Right, tree.Left
		tree.Left.Invert()
		tree.Right.Invert()
	}
}

func (node *treeNode) Parent() *treeNode {
	return node.parent
}

func (node *treeNode) Root() *treeNode {
	return node.root
}

// Sibling returns the other node at the same level of depth from the same parent.
func (node *treeNode) Sibling() *treeNode {
	if node.parent == nil {
		return nil
	}
	if node.parent.Left == node {
		return node.parent.Right
	}
	return node
}

// String returns the JSON stringified version of a tree.
func (tree treeNode) String() string {
	str, err := json.Marshal(tree)
	if err != nil {
		panic(err)
	}
	return string(str)
}
