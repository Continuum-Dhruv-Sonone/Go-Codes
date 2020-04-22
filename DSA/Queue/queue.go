package queue

import (
	"errors"
	"fmt"
)

type MyQueue struct {
	a []int
}

var q MyQueue

func NewQueue() *MyQueue {
	return &MyQueue{a: make([]int, 0)}
}

func (q *MyQueue) Enqueue(val int) {
	q.a = append(q.a, val)
}

func (q *MyQueue) Dequeue() error {
	if len(q.a) == 0 {
		return errors.New("The queue is empty")
	}

	if len(q.a) == 1 {
		q.a = make([]int, 0)
		return nil
	}

	q.a = q.a[1:]
	return nil
}

func (q *MyQueue) isEmpty() bool {
	if len(q.a) == 0 {
		return true
	}
	return false
}

func (q *MyQueue) Display() {
	fmt.Println(q.a)
}
