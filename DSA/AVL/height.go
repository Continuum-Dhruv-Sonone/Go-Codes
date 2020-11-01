package avl

func height(root *Node) int {
	if root == nil {
		return 0
	}

	left := height(root.Left)
	right := height(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}

func getBalance(root *Node) int {
	if root == nil {
		return 0
	}

	return height(root.Left) - height(root.Right)
}
