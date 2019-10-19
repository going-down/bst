package bst

import (
	"github.com/going-down/bst/pkg/binary_tree/node"
	"github.com/going-down/bst/pkg/element"
)

func (tree *Tree) Add(element element.Interface) {
	var ref node.Interface = &Node{element: element}
	*tree.findNode(element.GetKey()) = &ref
}
