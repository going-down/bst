package bst

type (
	Comparable interface {
		Less(other Comparable) bool
		Greater(other Comparable) bool
		Equal(other Comparable) bool
	}
	Element interface {
		Comparable
	}
	BinaryTreeNode interface {
		Left() **BinaryTreeNode
		Right() **BinaryTreeNode
		Value() *Element
	}
	Visitor interface {
		Visit(element *Element)
	}
	Traverser interface {
		Traverse(node *BinaryTreeNode, visitor *Visitor)
	}
	BinaryTree interface {
		Add(element Element)
		Walk(visitor *Visitor, traverser *Traverser)
	}
	PreOrder interface {
		Traverser
	}
	TreeElementPrinter interface {
		Visitor
	}
)
