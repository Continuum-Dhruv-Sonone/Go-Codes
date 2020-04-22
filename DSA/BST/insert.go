package bst

import "errors"

//Insert add a new node to bst tree
func (t *Tree) Insert(data int) error {
	if t.Root == nil {
		t.Root = NewNode(data)
		return nil
	}

	return t.Root.insertNode(data)
}

//Insert add node to a bst treer
func (n *Node) insertNode(data int) error {
	if n == nil {
		return errors.New("Cannot insert into a nil tree")
	}

	switch {
	case data < n.Data:
		if n.Left == nil {
			n.Left = NewNode(data)
			return nil
		}
		return n.Left.insertNode(data)
	case data > n.Data:
		if n.Right == nil {
			n.Right = NewNode(data)
			return nil
		}
		return n.Right.insertNode(data)
	default:
		return nil
	}
}
