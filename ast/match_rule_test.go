package ast

import (
	"testing"
	"github.com/dannypsnl/assert"

	"github.com/elz-lang/elz/util"

	"llvm.org/llvm/bindings/go/llvm"
)

func Test_match_rule(t *testing.T) {
	assert := assert.NewTester(t)
	c := NewContext()

	p1 := Pattern{
		E: &I32{Val: "10"},
		S: &Return{
			&I32{Val: "1"},
		},
	}
	rest := Pattern{
		E: nil,
		S: &Return{
			&I32{Val: "0"},
		},
	}
	m := NewMatch(&I32{Val: "10"}, []*Pattern{&p1}, &rest)

	f := &FnDef{
		Export:    false,
		Name:      "foo",
		Params:    []*Param{},
		Body:      []Stat{m},
		RetType:   "i32",
		Ctx:       c,
		Notations: []util.Notation{},
	}

	f.Check(c)
	f.Codegen(c)

	testEngine, err := llvm.NewExecutionEngine(c.Module)
	if err != nil {
		panic(err)
	}
	result := testEngine.RunFunction(f.Codegen(c), []llvm.GenericValue{})
	assert.Eq(result.Int(true), uint64(1))
}
