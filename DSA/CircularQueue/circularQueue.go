package main

import "fmt"

var (
	arr   [10]int
	front = -1
	rear  = -1
	size  = 10
)

func main() {
	isEmpty()
	enqueue(1)
	enqueue(2)
	enqueue(3)
	display()
	dequeue()
	dequeue()
	dequeue()
	display()
	isEmpty()
}

func isEmpty() bool {
	if front == -1 {
		return true
	}
	return false
}

func isFull() bool {
	if (front == 0 && rear == size-1) || (front == rear+1) {
		return true
	}
	return false
}

func enqueue(value int) {
	if isFull() {
		fmt.Println("Queue is full.")
		return
	}

	if front == -1 {
		front = 0
	}
	rear = (rear + 1) % size
	arr[rear] = value
}

func dequeue() {
	if isEmpty() {
		fmt.Println("Queue is empty.")
		return
	}
	if front == rear {
		front = -1
		rear = -1
		return
	}

	front = (front + 1) % size
}

func display() {

	if front > rear {
		for i := 0; i <= rear; i++ {
			fmt.Println(arr[i])
		}

		for i := front; i < size; i++ {
			fmt.Println(arr[i])
		}
		return
	}
	for i := front; i < rear; i++ {
		fmt.Println(arr[i])
	}
}
