package visit

import (
	"github.com/going-down/bst/pkg/element"
)

type Interface interface {
	Visit(element *element.Interface)
}

type SideEffect struct {
	callable func(node *element.Interface)
}

func (visitor *SideEffect) Visit(node *element.Interface) {
	visitor.callable(node)
}
