package avl

//Node contains the actual data and adress to left and right node
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

//Tree contains the pointer to root node of tree
type Tree struct {
	Root *Node
}

//NewTree creates a empty bst tree
func NewTree() *Tree {
	return &Tree{nil}
}

//NewNode creates a new node for bst tree
func NewNode(data int) *Node {
	return &Node{data, nil, nil}
}
