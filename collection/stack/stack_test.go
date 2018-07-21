package stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	if s.Len() != 0 {
		t.Error(`initial state of stack is incorrect`)
	}
}

func TestPopInitStack(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(`pop nil stack should not panic, just not thing happened`)
		}
	}()
	s := New()
	s.Pop()
}

func TestPushPop(t *testing.T) {
	s := New()
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

func TestEmptyAtInit(t *testing.T) {
	s := New()
	if !s.Empty() {
		t.Error(`should not have size before push anything into stack`)
	}
}

func TestEmptyAfterPush(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(1)
	s.Push(1)
	if s.Empty() {
		t.Error(`stack should not be empty after push 3 things into it`)
	}
}

func TestLastElem(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	if s.Last() != 2 {
		t.Errorf("expect: %v, actual: %v", 2, s.Last())
	}
}

func Test_getLastFromEmptyStackShouldBeSafe(t *testing.T) {
	s := New()
	s.Last()
}

type foo interface {
	Foo()
}

type fooImp struct{}

func (fi *fooImp) Foo() {}

func Test_stackT(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("should panic when push a value do not implement T into stack")
		}
	}()
	s := New().WithT((*foo)(nil))
	s.Push(&fooImp{})
	s.Push(1)
}

func TestMockCalc(t *testing.T) {
	s := New()
	s.Push(1.1)
	s.Push(2.4)
	le := s.Pop().(float64)
	re := s.Pop().(float64)
	s.Push(le + re)
	res := s.Pop()
	if res != 3.5 {
		t.Error(`Mock calculator show error`)
	}
	s.Push(res)
	s.Push(3.0)
	le = s.Pop().(float64)
	re = s.Pop().(float64)
	s.Push(le * re)
	if s.Pop() != 10.5 {
		t.Error(`Mock calculator show error`)
	}
}
