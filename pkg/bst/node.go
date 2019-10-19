package bst

import (
	"github.com/going-down/bst/pkg/binary_tree/node"
	"github.com/going-down/bst/pkg/element"
)

type Node struct {
	element     element.Interface
	left, right *node.Interface
}

func (node *Node) Left() **node.Interface   { return &node.left }
func (node *Node) Right() **node.Interface  { return &node.right }
func (node *Node) Data() *element.Interface { return &node.element }
