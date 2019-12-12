package assoc

import (
	"github.com/going-down/bst/pkg/comparable"
	"github.com/going-down/bst/pkg/element"
)

type Element struct {
	Key   comparable.Interface
	Value interface{}
}

func (element *Element) GetKey() comparable.Interface { return element.Key }
func (element *Element) GetValue() *interface{}       { return &element.Value }

func MakeElement(p comparable.Interface, v interface{}) element.Interface {
	return &Element{Key: p, Value: v}
}
