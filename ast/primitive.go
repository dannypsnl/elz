package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Id struct {
	Val string
}

func (i *Id) Codegen(ctx *Context) llvm.Value {
	// FIXME: id should not only be a global var
	return ctx.GlobalVars[i.Val].v
}

// At here we can see, ident's type need to logging in Context
// So Context should send into Type method and Context::Vars
// need a new structure for usage
func (i *Id) Type(ctx *Context) string {
	// FIXME: id should not only be a global var
	return ctx.VarsType[i.Val]
}

type Str struct {
	Val string
}

func (s *Str) Codegen(ctx *Context) llvm.Value {
	return llvm.ConstString(s.Val, false)
}

func (s *Str) Type(*Context) string {
	return "str"
}

type Number struct {
	Val string
}

func (n *Number) Codegen(*Context) llvm.Value {
	return llvm.ConstFloatFromString(llvm.FloatType(), n.Val)
}

func (n *Number) Type(*Context) string {
	return "num"
}

type F32 struct {
	Val string
}

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

func (f64 *F64) Type(*Context) string { return "f64" }

type I32 struct {
	Val string
}

func (i32 *I32) Codegen(*Context) llvm.Value {
	return llvm.ConstIntFromString(llvm.Int32Type(), i32.Val, 10)
}
func (i32 *I32) Type(*Context) string { return "i32" }

type I64 struct {
	Val string
}

func (i64 *I64) Codegen(*Context) llvm.Value {
	return llvm.ConstIntFromString(llvm.Int64Type(), i64.Val, 10)
}

func (i64 *I64) Type(*Context) string { return "i64" }
