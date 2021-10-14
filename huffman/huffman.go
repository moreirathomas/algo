// https://en.wikipedia.org/wiki/Huffman_coding

package huffman

import (
	"encoding/json"
	"sort"
)

type Node struct {
	Value  string `json:"value"`
	Count  int    `json:"count"`
	Parent *Node  `json:"-"`
	Left   *Node  `json:"left,omitempty"`
	Right  *Node  `json:"right,omitempty"`
}

// SortNodes implements sort.Interface.
type SortNodes []*Node

func (n SortNodes) Len() int {
	return len(n)
}

func (n SortNodes) Less(i, j int) bool {
	return n[i].Count < n[j].Count
}

func (n SortNodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func BuildSorted(leaves []*Node) *Node {
	if len(leaves) == 0 {
		return nil
	}

	for len(leaves) > 1 {
		// Take the two leaves with smallest probability and
		// create a new node having these two nodes as children.

		left, right := leaves[0], leaves[1]
		count := left.Count + right.Count
		node := &Node{
			Left:  left,
			Right: right,
			Count: count,
		}

		left.Parent = node
		right.Parent = node

		// Exclude the two leaf nodes.
		leaves = leaves[2:]

		// Insert.
		i := sort.Search(len(leaves), func(i int) bool { return leaves[i].Count >= count })
		leaves = append(leaves, &Node{})
		copy(leaves[i+1:], leaves[i:])
		leaves[i] = node
	}

	return leaves[0]
}

func Build(leaves []*Node) *Node {
	sort.Stable((SortNodes(leaves)))
	return BuildSorted(leaves)
}

func (n Node) String() string {
	str, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	return string(str)
}
