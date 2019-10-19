package walk

import (
	"github.com/going-down/bst/pkg/binary_tree/node"
)

type Interface interface {
	Walk(node *node.Interface)
}

type Controller struct {
	callOrder []func(node *node.Interface)
}

func (walker *Controller) Walk(node *node.Interface) {
	for _, callable := range walker.callOrder {
		callable(node)
	}
}
