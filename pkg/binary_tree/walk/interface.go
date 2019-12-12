package walk

import (
	"github.com/going-down/bst/pkg/binary_tree/node"
	"github.com/going-down/bst/pkg/binary_tree/visit"
	"github.com/going-down/bst/pkg/element"
)

type Interface interface {
	GetVisitor() visit.Interface
	Walk(node *node.Interface)
}

type partialWalker func(walker Interface, node *node.Interface)
type Sequence []partialWalker

func MakeDataWalker(accessor func(node.Interface) *element.Interface) partialWalker {
	return func(walker Interface, node *node.Interface) {
		walker.GetVisitor().Visit(*accessor(*node))
	}
}
func MakeBranchWalker(accessor func(node.Interface) **node.Interface) partialWalker {
	return func(walker Interface, node *node.Interface) {
		walker.Walk(*accessor(*node))
	}
}

var Data = MakeDataWalker(node.Interface.Data)
var Left = MakeBranchWalker(node.Interface.Left)
var Right = MakeBranchWalker(node.Interface.Right)
