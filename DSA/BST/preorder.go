package bst

import stack "github.com/ContinuumLLC/GO/DSA/Stack"

// Preorder ...
func Preorder(root *Node) []interface{} {
	var list []interface{}

	if root == nil {
		return list
	}

	st := stack.NewStack()
	st.Push(root)

	for !st.IsEmpty() {

		node := st.Peek()

		n := node.(*Node)
		list = append(list, n.Data)

		st.Pop()

		if n.Right != nil {
			st.Push(n.Right)
		}

		if n.Left != nil {
			st.Push(n.Left)
		}

	}

	return list
}
