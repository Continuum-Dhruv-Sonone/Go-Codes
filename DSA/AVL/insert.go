package avl

// Insert ...
func (t *Tree) Insert(data int) {
	if t.Root == nil {
		t.Root = NewNode(data)
		return
	}

	t.Root = insert(t.Root, data)

}

func insert(root *Node, data int) *Node {
	if root == nil {
		return NewNode(data)
	} else if root.Data > data {
		root.Left = insert(root.Left, data)
	} else if root.Data < data {
		root.Right = insert(root.Right, data)
	} else {
		return root
	}

	balance := getBalance(root)

	if balance > 1 && data < root.Left.Data {
		return rotateRight(root)
	}

	if balance < -1 && data > root.Right.Data {
		return rotateLeft(root)
	}

	if balance > 1 && data > root.Left.Data {
		root.Left = rotateLeft(root.Left)
		return rotateRight(root)
	}

	if balance > 1 && data < root.Right.Data {
		root.Right = rotateRight(root.Right)
		return rotateLeft(root)
	}

	return root
}
