package node

import "github.com/going-down/bst/pkg/element"

type (
	Interface interface {
		Left() **Interface
		Right() **Interface
		Data() *element.Interface
	}
)
