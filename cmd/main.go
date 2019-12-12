package main

import (
	"fmt"
	"github.com/going-down/bst/pkg/assoc"
	"github.com/going-down/bst/pkg/binary_tree/visit"
	"github.com/going-down/bst/pkg/binary_tree/walk"
	"github.com/going-down/bst/pkg/bst"
	"github.com/going-down/bst/pkg/comparable"
	"github.com/going-down/bst/pkg/element"
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
	intMap.Add(&assoc.Element{Key: &comparable.Integer{Value: 9}, Value: 97})
	for _, i := range intMapInit {
		intMap.Add(&assoc.Element{Key: &comparable.Integer{Value: i[0]}, Value: i[1]})
	}
	intMap.Iterate(&bst.Walker{
		Visitor: &visit.SideEffect{Callable: func(element element.Interface) {
			fmt.Printf("%d: %d\n", element.GetKey().(*comparable.Integer).Value, (*element.GetValue()).(int))
		}},
		WalkSequence: walk.Sequence{walk.Left, walk.Data, walk.Right},
	})

	strMapInit := [][]interface{}{
		{"a", 1},
		{"b", 1},
		{"c", 1},
		{"d", 1},
		{"e", 1},
		{"f", 1},
	}
	strMap := bst.Tree{}
	for _, i := range strMapInit {
		strMap.Add(&assoc.Element{Key: &comparable.String{Value: i[0].(string)}, Value: i[1]})
	}
	//strMap.Add(&assoc.Element{Key: &comparable.String{Value: "a"}, Value: 97}) // panic
	strMap.Set(&comparable.String{Value: "a"}, 97)
	// don't mix up key types if comparison between them is not provided
	// using comparable.Interface methods overriding
}
