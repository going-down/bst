package bst

import (
	"github.com/going-down/bst/pkg/assoc"
	"github.com/going-down/bst/pkg/binary_tree/node"
	"github.com/going-down/bst/pkg/comparable"
)

func (tree *Tree) Set(key comparable.Interface, value interface{}) {
	found := tree.findNode(key)
	if *found == nil {
		var ref node.Interface = &Node{element: &assoc.Element{Key: key, Value: value}}
		*found = &ref
	} else {
		*(*(**found).Data()).GetValue() = value
	}
}
