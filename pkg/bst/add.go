package bst

import (
	"github.com/going-down/bst/pkg/binary_tree/node"
	"github.com/going-down/bst/pkg/element"
)

func (tree *Tree) Add(element element.Interface) {
	var ref node.Interface = &Node{element: element}
	found := tree.findNode(element.GetKey())
	if *found != nil {
		panic("Element key must be unique when add")
	}
	*found = &ref
}
