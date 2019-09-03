package LinkedList

type Node struct {
	data interface{}
	next *Node
}

func (n *Node) Next() *Node {
	return n.next
}
