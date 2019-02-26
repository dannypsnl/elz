// Package stack is stack implementation
package stack

import (
	"fmt"
	"reflect"
)

// Stack is data structure: stack's implement
type Stack struct {
	stack  []interface{}
	limitT *reflect.Type
}

// New create a new stack
func New() *Stack {
	return &Stack{
		stack: make([]interface{}, 0),
	}
}

// Just like Stack<T> in a language has generic
func (s *Stack) WithT(v interface{}) *Stack {
	t := reflect.TypeOf(v).Elem()
	s.limitT = &t
	return s
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
	if s.limitT != nil {
		if !reflect.TypeOf(element).Implements(*s.limitT) {
			panic(fmt.Sprintf("element must implement type: %s", *s.limitT))
		}
	}
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
