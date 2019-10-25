package test

import (
	"github.com/going-down/bst/pkg/assoc"
	"github.com/going-down/bst/pkg/bst"
	"github.com/going-down/bst/pkg/comparable"
	//"os"
	"testing"

	"gotest.tools/assert"
)

func TestEverything(t *testing.T) {
	intMap := bst.Tree{}
	intMap.Add(&assoc.Element{Key: &comparable.Integer{Value: 9}, Value: 97})
	assert.Assert(t, (*intMap.Find(&comparable.Integer{Value: 14})).(int))
}
