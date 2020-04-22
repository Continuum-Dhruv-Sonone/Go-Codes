package bst

//Height returns the height of bst
func (t *Tree) Height() int {
	if t.Root == nil {
		return -1
	}

	return t.Root.heightNode()
}

func (node *Node) heightNode() int {
	if node == nil {
		return -1
	}

	leftHeight := node.Left.heightNode()
	rightHeight := node.Right.heightNode()

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}
