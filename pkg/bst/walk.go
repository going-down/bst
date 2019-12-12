package bst

import (
	"github.com/going-down/bst/pkg/binary_tree/node"
	"github.com/going-down/bst/pkg/binary_tree/visit"
	"github.com/going-down/bst/pkg/binary_tree/walk"
)

type Walker struct {
	Visitor      visit.Interface
	WalkSequence walk.Sequence
}

func (walker *Walker) GetVisitor() visit.Interface { return walker.Visitor }

func (walker *Walker) Walk(node *node.Interface) {
	if node != nil {
		for _, callable := range walker.WalkSequence {
			callable(walker, node)
		}
	}
}
