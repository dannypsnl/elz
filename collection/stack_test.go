package collection

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s.Len() != 0 {
		t.Error(`initial state of stack is incorrect`)
	}
}

func TestPushPop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)
	if s.Len() != 5 {
		t.Error(`stack didn't get new element by Push`)
	}
	i := s.Pop()
	if i != 1 || s.Len() != 4 {
		t.Error(`Pop method didn't work as expected`, s.stack, i)
	}
}
