package bst

import "github.com/going-down/bst/pkg/comparable"

func (tree *Tree) Find(key comparable.Interface) *interface{} {
	return (*(**tree.findNode(key)).Data()).GetValue()
}
