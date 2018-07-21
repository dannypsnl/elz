// Package stack is stack implementation
package stack

// Stack is data structure: stack's implement
type Stack struct {
	stack []interface{}
}

// New create a new stack
func New() *Stack {
	return &Stack{
		stack: make([]interface{}, 0),
	}
}

// Len return the length of stack
func (s *Stack) Len() int {
	return len(s.stack)
}

// Empty return the stack is empty or not
func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Last() interface{} {
	if len(s.stack) == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}

// Push push new element into stack
func (s *Stack) Push(element interface{}) {
	s.stack = append(s.stack, element)
}

// Pop pop off element from stack and return it
func (s *Stack) Pop() (res interface{}) {
	l := len(s.stack)
	if l <= 0 {
		return
	}
	res = s.stack[l-1]
	s.stack = s.stack[:l-1]
	return
}
