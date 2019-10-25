package main

import (
	"fmt"
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
	intMap.Add(&assoc.Element{Key: &comparable.Integer{Value: 9}, Value: 97})
	for _, i := range intMapInit {
		intMap.Add(&assoc.Element{Key: &comparable.Integer{Value: i[0]}, Value: i[1]})
	}
	fmt.Print((*intMap.Find(&comparable.Integer{Value: 14})).(int))
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
	//strMap.Set(&comparable.Integer{Value: 3}, 97)
	// don't mix up key types if comparison between them is not provided
	// using comparable.Interface methods overriding
}
