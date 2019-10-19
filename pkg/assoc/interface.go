package assoc

import (
	"github.com/going-down/bst/pkg/comparable"
	"github.com/going-down/bst/pkg/element"
)

type Interface interface {
	Add(element element.Interface)
	Find(key comparable.Interface) *interface{}
	Set(key comparable.Interface, value interface{})
}
