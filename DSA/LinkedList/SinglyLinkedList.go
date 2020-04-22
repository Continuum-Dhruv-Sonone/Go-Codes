package LinkedList

import "fmt"

type SinglyLinkedList struct {
	front  *Node
	length int
}

func NewLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{length: 0}
}

func (sl *SinglyLinkedList) Front() *Node {
	return sl.front
}

func (sl *SinglyLinkedList) Append(n *Node) {
	if sl.front == nil {
		sl.front = n
	} else {
		currentNode := sl.front

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = n
	}

	sl.length++
}

func (sl *SinglyLinkedList) Prepend(n *Node) {
	if sl.front == nil {
		sl.front = n
	} else {
		n.next = sl.front
		sl.front = n
	}

	sl.length++
}

func (sl *SinglyLinkedList) Display() {
	if sl.length == 0 {
		fmt.Println("The LinkedList is empty.")
	} else {
		currentNode := sl.front
		for currentNode != nil {
			fmt.Println(currentNode.data)
			currentNode = currentNode.next
		}
	}
}

func (sl *SinglyLinkedList) Length() int {
	return sl.length
}

func (sl *SinglyLinkedList) Back() *Node {
	currentNode := sl.front

	for currentNode != nil && currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}
