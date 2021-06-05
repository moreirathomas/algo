package main

import (
	"github.com/moreirathomas/algo/trees"
)

func main() {
	binaryTree := trees.NewBinaryTree()
	child := binaryTree.AddNode(50, trees.ChildLeft)
	grandChild := child.AddNode(25, trees.ChildLeft)
	binaryTree.AddNode(50, trees.ChildRight)

	println(binaryTree.String())
	binaryTree.Invert()
	println(binaryTree.String())

	root := grandChild.Root
	println("root value:", root.Val)
}
