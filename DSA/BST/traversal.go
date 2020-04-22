package bst

import (
	"errors"
	"fmt"
)

//InorderTraversal traverses the bst in inorder
func (t *Tree) InorderTraversal() {
	inorder(t.Root)
}

func inorder(node *Node) {
	if node != nil {
		inorder(node.Left)
		fmt.Println(node.Data)
		inorder(node.Right)
	}
}

//PreorderTraversal traverses the bst in preorder
func (t *Tree) PreorderTraversal() {
	preorder(t.Root)
}

func preorder(node *Node) {

	if node != nil {
		fmt.Println(node.Data)
		preorder(node.Left)
		preorder(node.Right)
	}
}

//PostOrderTraversal traverses the bst in post order
func (t *Tree) PostOrderTraversal() {
	postorder(t.Root)
}

func postorder(node *Node) {
	if node == nil {
		return
	}
	postorder(node.Left)
	postorder(node.Right)
	fmt.Println(node.Data)
}

type queue struct {
	array []*Node
}

func newQueue() queue {
	return queue{
		array: make([]*Node, 0),
	}
}

func (q *queue) enqueue(value *Node) {
	q.array = append(q.array, value)
}

func (q *queue) dequeue() (*Node, error) {
	if q.isEmpty() {
		return nil, errors.New("Empty queue")
	}

	removeElement := q.array[0]
	q.array[0] = nil
	q.array = q.array[1:]
	return removeElement, nil
}

func (q *queue) isEmpty() bool {
	if len(q.array) <= 0 {
		return true
	}

	return false
}

//LevelOrderTraversal traverses the bst in level order
func (t *Tree) LevelOrderTraversal() {
	if t.Root == nil {
		return
	}

	q := newQueue()
	q.enqueue(t.Root)
	for !q.isEmpty() {
		n, _ := q.dequeue()
		fmt.Println(n.Data)

		if n.Left != nil {
			q.enqueue(n.Left)
		}

		if n.Right != nil {
			q.enqueue(n.Right)
		}
	}
}
