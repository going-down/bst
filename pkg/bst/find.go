package bst

import "github.com/going-down/bst/pkg/comparable"

func (tree *Tree) Find(key comparable.Interface) *interface{} {
	found := tree.findNode(key)
	if *found == nil || (**found).Data() == nil {
		return nil
	}
	return (*(**found).Data()).GetValue()
}
