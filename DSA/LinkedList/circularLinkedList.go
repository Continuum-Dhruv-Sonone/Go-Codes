package LinkedList

import "fmt"

type CircularLinkedList struct {
	tail   *Node
	length int
}

func NewCirurcularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{length: 0}
}

func (cll *CircularLinkedList) Front() *Node {
	return cll.tail.next
}

func (cll *CircularLinkedList) Append(n *Node) {
	if cll.tail == nil {
		cll.tail = n
		cll.tail.next = cll.tail
	} else {

		n.next = cll.tail.next
		cll.tail.next = n
		cll.tail = n
	}

	cll.length++
}

func (cll *CircularLinkedList) Prepend(n *Node) {
	if cll.tail == nil {
		cll.tail = n
		cll.tail.next = cll.tail
	} else {

		n.next = cll.tail.next
		cll.tail.next = n

	}

	cll.length++
}

func (cll *CircularLinkedList) Display() {

	if cll.tail == nil {
		return
	}

	curr := cll.tail.next

	for {

		fmt.Print(curr.data)
		curr = curr.next
		if curr == cll.tail {
			break
		}
	}

}

func (cll *CircularLinkedList) Length() int {
	return cll.length
}
