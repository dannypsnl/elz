package collection

type Stack struct {
	stack []interface{}
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]interface{}, 0),
	}
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) Push(ele interface{}) {
	s.stack = append(s.stack, ele)
}
func (s *Stack) Pop() (res interface{}) {
	l := len(s.stack)
	res = s.stack[l-1]
	s.stack = s.stack[:l-1]
	return
}
