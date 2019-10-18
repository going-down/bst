package bst

type (
	node struct {
		value       Element
		left, right *BinaryTreeNode
	}
	Tree struct {
		root *BinaryTreeNode
	}
)

func (node *node) Left() **BinaryTreeNode {
	return &node.left
}

func (node *node) Right() **BinaryTreeNode {
	return &node.right
}

func (node *node) Value() *Element {
	return &node.value
}

func (tree *Tree) Walk(visitor *Visitor, traverser *Traverser) {
	(*traverser).Traverse(tree.root, visitor)
}

func (tree *Tree) find(key Element) **BinaryTreeNode {
	current := &tree.root
	for *current != nil {
		if key.Less(*(**current).Value()) {
			current = (**current).Left()
		} else if key.Greater(*(**current).Value()) {
			current = (**current).Right()
		} else {
			return current
		}
	}
	return current
}

func (tree *Tree) Add(key Element) {
	var ref BinaryTreeNode = &node{value: key}
	*tree.find(key) = &ref
}

type ElementInt struct {
	value int
}

func (ref *ElementInt) Less(other Comparable) bool {
	return ref.value < other.(*ElementInt).value
}

func (ref *ElementInt) Greater(other Comparable) bool {
	return ref.value > other.(*ElementInt).value
}

func (ref *ElementInt) Equal(other Comparable) bool {
	return ref.value == other.(*ElementInt).value
}
