package bst

import stack "github.com/ContinuumLLC/GO/DSA/Stack"

// Postorder ...
func Postorder(root *Node) []interface{} {
	var list []interface{}
	prev := root

	if root == nil {
		return list
	}

	st := stack.NewStack()
	st.Push(root)

	for !st.IsEmpty() {
		node := st.Peek()

		curr := node.(*Node)

		st.Pop()

		if curr == prev || curr == prev.Left || curr == prev.Right {
			if curr.Left != nil {
				st.Push(curr.Left)
			} else if curr.Right != nil {
				st.Push(curr.Right)
			} else {
				list = append(list, curr.Data)
			}
		}

		if curr.Left == prev {
			if curr.Right != nil {
				st.Push(curr.Right)
			} else {
				list = append(list, curr.Data)
			}
		}

		if curr.Right == prev {

			list = append(list, curr.Data)
		}

		prev = curr
	}

	return list
}
