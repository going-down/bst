package assoc

import "github.com/going-down/bst/pkg/comparable"

type Element struct {
	Key   comparable.Interface
	Value interface{}
}

func (element *Element) GetKey() comparable.Interface { return element.Key }
func (element *Element) GetValue() *interface{}       { return &element.Value }
