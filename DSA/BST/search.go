package bst

//Search return true if node with given data exists in bst
func (t *Tree) Search(data int) bool {
	if t.Root == nil {
		return false
	}
	return t.Root.searchNode(data)
}

//Search return true if node with given data exists in bst
func (n *Node) searchNode(data int) bool {
	if n == nil {
		return false
	}

	switch {
	case data < n.Data:
		return n.Left.searchNode(data)
	case data > n.Data:
		return n.Right.searchNode(data)
	default:
		return true
	}
}
