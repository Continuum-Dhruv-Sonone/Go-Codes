package bst

import stack "github.com/ContinuumLLC/GO/DSA/Stack"

// Inorder ...
func Inorder(root *Node) []interface{} {
	var list []interface{}

	st := stack.NewStack()

	curr := root

	for curr != nil {

		for curr != nil {
			st.Push(curr)
			curr = curr.Left
		}

		node := st.Peek()

		n := node.(*Node)
		list = append(list, n.Data)

		st.Pop()

		if n.Right != nil {
			curr = curr.Left
		}

		if st.IsEmpty() {
			break
		}

	}

	return list
}
