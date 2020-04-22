package bst

//Delete is used to delete node in bst
func (t *Tree) Delete(data int) {
	deleteNode(t.Root, data)
}

func deleteNode(node *Node, data int) *Node {
	if node == nil {
		return node
	}

	if data < node.Data {
		node.Left = deleteNode(node.Left, data)
		return node
	} else if data > node.Data {
		node.Right = deleteNode(node.Right, data)
		return node
	} else {
		if node.Left == nil && node.Right == nil {
			node = nil
			return nil
		}

		if node.Left == nil {
			node = node.Right
			return node
		}

		if node.Right == nil {
			node = node.Right
			return node
		}

		temp := findMin(node.Right)

		node.Data = temp.Data
		node.Right = deleteNode(node.Right, temp.Data)

		return node
	}

}

func findMin(node *Node) *Node {
	curr := node

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}
