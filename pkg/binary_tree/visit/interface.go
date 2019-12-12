package visit

import (
	"github.com/going-down/bst/pkg/element"
)

type Interface interface {
	Visit(element element.Interface)
}

type SideEffect struct {
	Callable func(element element.Interface)
}

func (visitor *SideEffect) Visit(element element.Interface) {
	visitor.Callable(element)
}
