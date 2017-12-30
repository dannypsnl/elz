package collection

// Stack is data structure: stack's implement
type Stack struct {
	stack []interface{}
}

// NewStack create a new stack
func NewStack() *Stack {
	return &Stack{
		stack: make([]interface{}, 0),
	}
}

// Len return the length of stack
func (s *Stack) Len() int {
	return len(s.stack)
}

// Push push new element into stack
func (s *Stack) Push(element interface{}) {
	s.stack = append(s.stack, element)
}

// Pop pop off element from stack and return it
func (s *Stack) Pop() (res interface{}) {
	l := len(s.stack)
	res = s.stack[l-1]
	s.stack = s.stack[:l-1]
	return
}
