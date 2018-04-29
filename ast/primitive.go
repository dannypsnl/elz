package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

type Id struct {
	Val string
}

func (i *Id) Codegen(c *Context) llvm.Value {
	// FIXME: id should not only be a global var
	v, ok := c.Var(i.Val)
	if ok {
		return v
	}
	return llvm.Value{}
}

func (i *Id) Check(c *Context) {
	_, ok := c.VarType(i.Val)
	if !ok {
		c.Reporter.Emit(fmt.Sprintf("var: %s not found", i.Val))
	}
}

// At here we can see, ident's type need to logging in Context
// So Context should send into Type method and Context::Vars
// need a new structure for usage
func (i *Id) Type(c *Context) string {
	// FIXME: id should not only be a global var
	v, _ := c.VarType(i.Val)
	return v
}

type Str struct {
	Val string
}

func (s *Str) Check(*Context) {}
func (s *Str) Codegen(ctx *Context) llvm.Value {
	return llvm.ConstString(s.Val, false)
}

func (s *Str) Type(*Context) string {
	return "str"
}

type F32 struct {
	Val string
}

func (f32 *F32) Check(*Context) {}
func (f32 *F32) Codegen(*Context) llvm.Value {
	return llvm.ConstFloatFromString(llvm.FloatType(), f32.Val)
}

func (f32 *F32) Type(*Context) string { return "f32" }

type F64 struct {
	Val string
}

func (f64 *F64) Codegen(*Context) llvm.Value {
	return llvm.ConstFloatFromString(llvm.DoubleType(), f64.Val)
}

func (f64 *F64) Check(*Context)       {}
func (f64 *F64) Type(*Context) string { return "f64" }

type I8 struct {
	Val string
}

func (i8 *I8) Check(*Context) {}
func (i8 *I8) Codegen(*Context) llvm.Value {
	return llvm.ConstIntFromString(LLVMType("i8"), i8.Val, 10)
}
func (i8 *I8) Type(*Context) string { return "i8" }

type I32 struct {
	Val string
}

func (i32 *I32) Check(*Context) {}
func (i32 *I32) Codegen(*Context) llvm.Value {
	return llvm.ConstIntFromString(LLVMType("i32"), i32.Val, 10)
}
func (i32 *I32) Type(*Context) string { return "i32" }

type I64 struct {
	Val string
}

func (i64 *I64) Check(*Context) {}
func (i64 *I64) Codegen(*Context) llvm.Value {
	return llvm.ConstIntFromString(llvm.Int64Type(), i64.Val, 10)
}

func (i64 *I64) Type(*Context) string { return "i64" }
