package main

import (
	"github.com/going-down/bst/pkg/bst"
)

func main() {
	node := new(node)
	tree := new(Tree)
	val := new(ElementInt)
	val.value = 100
	node.value = val
	tree.root = node
	arr := []int{1, 16, 32, 2, 64, 128, 4, 8}
	for _, v := range arr {
		tree.Add(&ElementInt{value: v})
	}
	//tree.root.left = tree.root
	//tree.root.right = tree.root
	//	print(tree.root.value.(*ElementInt).value)
	tree.Print()
}
