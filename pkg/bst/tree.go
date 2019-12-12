package bst

import (
	"github.com/going-down/bst/pkg/binary_tree/node"
	"github.com/going-down/bst/pkg/binary_tree/walk"
)

type Tree struct {
	root *node.Interface
}

func (tree *Tree) Iterate(p walk.Interface) {
	p.Walk(tree.root)
}
