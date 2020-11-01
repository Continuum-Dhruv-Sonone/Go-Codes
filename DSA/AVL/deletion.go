package avl

// Delete ...
func (t *Tree) delete(data int) {
	if t.Root == nil {
		return
	}

	t.Root = delete(t.Root, data)

}

func delete(node *Node, data int) *Node {
	if node == nil {
		return node
	}

	if data < node.Data {
		node.Left = delete(node.Left, data)
	} else if data > node.Data {
		node.Right = delete(node.Right, data)
		return node
	} else {
		if node.Left == nil && node.Right == nil {
			node = nil
		}

		if node.Left == nil {
			node = node.Right
		}

		if node.Right == nil {
			node = node.Right
		}

		temp := findMin(node.Right)

		node.Data = temp.Data
		node.Right = delete(node.Right, temp.Data)
	}

	if node == nil {
		return node
	}

	balance := getBalance(node)

	if balance > 1 && getBalance(node.Left) >= 0 {
		return rotateRight(node)
	}

	if balance < -1 && getBalance(node.Right) <= 0 {
		return rotateLeft(node)
	}

	if balance > 1 && getBalance(node.Left) < 0 {
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}

	if balance < -1 && getBalance(node.Left) > 0 {
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	return node
}

func findMin(node *Node) *Node {
	curr := node

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}
