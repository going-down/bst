package element

import "github.com/going-down/bst/pkg/comparable"

type Interface interface {
	GetKey() comparable.Interface
	GetValue() *interface{}
}
