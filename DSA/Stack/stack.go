package stack

import (
	"errors"
	"fmt"
)

type MyStack struct {
	a []int
}

var stk MyStack

func NewStack() *MyStack {
	return &MyStack{a: make([]int, 0)}
}

func (s *MyStack) IsEmpty() bool {
	if len(s.a) > 0 {
		return false
	}

	return true
}

func (s *MyStack) Top() (int, error) {
	if stk.IsEmpty() {
		return 0, errors.New("The stack is empty")
	}
	return s.a[len(s.a)-1], nil
}

func (s *MyStack) Push(val int) {
	s.a = append(s.a, val)
}

func (s *MyStack) Pop() error {
	if s.IsEmpty() {
		return errors.New("The stack is empty")
	}

	if len(s.a) == 1 {
		s.a = make([]int, 0)
		return nil
	}
	s.a = s.a[:len(s.a)-1]
	return nil
}

func (s *MyStack) Display() {
	fmt.Println(s.a)
}
