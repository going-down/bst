package main

import (
	"github.com/going-down/bst/pkg/assoc"
	"github.com/going-down/bst/pkg/bst"
	"github.com/going-down/bst/pkg/comparable"
)

func main() {
	intMapInit := [][]int{
		{10, 1},
		{11, 1},
		{12, 1},
		{14, 1},
		{15, 1},
		{16, 1},
	}
	intMap := bst.Tree{}
	intMap.Add(&assoc.Element{Key: &comparable.Integer{Value: 10}, Value: 97})
	for _, i := range intMapInit {
		intMap.Add(&assoc.Element{Key: &comparable.Integer{Value: i[0]}, Value: i[1]})
	}

	strMapInit := [][]interface{}{
		{"a", 1},
		{"b", 1},
		{"c", 1},
		{"d", 1},
		{"e", 1},
		{"f", 1},
	}
	strMap := bst.Tree{}
	strMap.Add(&assoc.Element{Key: &comparable.String{Value: "z"}, Value: 97}) //TODO: fix when key equals "a"
	for _, i := range strMapInit {
		strMap.Add(&assoc.Element{Key: &comparable.String{Value: i[0].(string)}, Value: i[1]})
	}
	// don't mix up key types if no comparison between said types
	// using comparable.Interface methods override
}
