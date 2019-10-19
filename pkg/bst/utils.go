package bst

import (
	"github.com/going-down/bst/pkg/binary_tree/node"
	"github.com/going-down/bst/pkg/comparable"
)

func (tree *Tree) findNode(key comparable.Interface) **node.Interface {
	current := &tree.root
	for *current != nil {
		currentKey := (*(**current).Data()).GetKey()
		if key.Less(currentKey) {
			current = (**current).Left()
		} else if key.Greater(currentKey) {
			current = (**current).Right()
		} else {
			return current
		}
	}
	return current
}
