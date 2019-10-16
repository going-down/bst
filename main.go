package main

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
	node struct {
		value       Element
		left, right *BinaryTreeNode
	}
	Tree struct {
		root *BinaryTreeNode
	}
	PreOrder interface {
		Traverser
	}
	TreeElementPrinter interface {
		Visitor
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

func (self *Tree) find(key Element) **BinaryTreeNode {
	current := &self.root
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

func (self *Tree) Add(key Element) {
	var ref BinaryTreeNode = &node{value: key}
	(*self.find(key)) = &ref
}

type ElementInt struct {
	value int
}

func (self *ElementInt) Less(other Comparable) bool {
	return self.value < other.(*ElementInt).value
}

func (self *ElementInt) Greater(other Comparable) bool {
	return self.value > other.(*ElementInt).value
}

func (self *ElementInt) Equal(other Comparable) bool {
	return self.value == other.(*ElementInt).value
}

type X interface {
	foo() **X
}

type Y struct {
	x *X
	i int
}

func (y *Y) foo() **X {
	return &y.x
}

func nofun() Y {
	return Y{x:nil, i:100}
}

func main() {
	y:=nofun()
	var ref X
	ref = &y
	ref.foo()
	var rref *Y
	rref = ref.(*Y)
	print(rref.i)
	t := Tree{}
	//var e Element
	//e = &ElementInt{104}
	t.Add(&ElementInt{104})
	//var nref *Node
	//nref = (*t.root).(*Node)
	//var vref *ElementInt
	//vref =
	print((*t.root).(*node).value.(*ElementInt).value)
		//(**(Y{x:&ref}).foo()).(*Y)

	//(**y.foo()).(Y) = Y{x:nil}}
}
/*



func label (current *node, indent int) {
	if (current != nil) {
		for i := 0; i < indent; i += 1 {
			print(" ")
		}
		print(current.value.(*ElementInt).value)
		print("\n")
		label(current.left, indent+4)
		label(current.right, indent+4)
	}
}


func (self *Tree) Print() {
	label(self.root, 0)
}

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
*/