package ast

import (
	"bytes"
	"fmt"

	"github.com/elz-lang/elz/errors"

	"llvm.org/llvm/bindings/go/llvm"
)

func NewContext() *Context {
	c := &Context{
		Parent:   nil,
		Reporter: errors.NewReporter(),
		Module:   llvm.NewModule("main"),
		Context:  llvm.NewContext(),
		Vars:     make(map[string]llvm.Value),
		VarsType: make(map[string]string),
		Types:    make(map[string]llvm.Type),
		Builder:  llvm.NewBuilder(),

		functions:        make(map[string]*Function),
		builtInOperators: make(map[string]string),
	}

	c.builtInOperators["+(i32,i32)"] = "i32"
	c.builtInOperators["+(i64,i64)"] = "i64"
	c.builtInOperators["-(i32,i32)"] = "i32"
	c.builtInOperators["-(i64,i64)"] = "i64"
	c.builtInOperators["*(i32,i32)"] = "i32"
	c.builtInOperators["*(i64,i64)"] = "i64"
	c.builtInOperators["/(i32,i32)"] = "i32"
	c.builtInOperators["/(i64,i64)"] = "i64"
	c.builtInOperators["==(i32,i32)"] = "bool"
	c.builtInOperators["==(i64,i64)"] = "bool"

	c.builtInOperators["+(f32,f32)"] = "f32"
	c.builtInOperators["+(f64,f64)"] = "f64"
	c.builtInOperators["-(f32,f32)"] = "f32"
	c.builtInOperators["-(f64,f64)"] = "f64"
	c.builtInOperators["*(f32,f32)"] = "f32"
	c.builtInOperators["*(f64,f64)"] = "f64"
	c.builtInOperators["/(f32,f32)"] = "f32"
	c.builtInOperators["/(f64,f64)"] = "f64"
	c.builtInOperators["==(f32,f32)"] = "bool"
	c.builtInOperators["==(f64,f64)"] = "bool"

	return c
}

type Context struct {
	Parent   *Context
	Reporter *errors.Reporter
	Module   llvm.Module
	Context  llvm.Context
	Vars     map[string]llvm.Value
	VarsType map[string]string
	Types    map[string]llvm.Type
	Builder  llvm.Builder

	functions        map[string]*Function
	builtInOperators map[string]string
}

type Function struct {
	value   llvm.Value
	retType string
}

func (c *Context) NewType(name string, t llvm.Type) {
	c.Types[name] = t
}

func (c *Context) Type(name string) llvm.Type {
	typ := LLVMType(name)
	if typ.String() != "VoidType" {
		return typ
	}
	if name == "()" {
		return llvm.VoidType()
	}
	return llvm.PointerType(c.Module.GetTypeByName(name), 0)
}

func (c *Context) NewFunc(signature string, retType string) {
	c.functions[signature] = &Function{
		value:   llvm.Value{},
		retType: retType,
	}
}

func (c *Context) FuncValue(signature string, val llvm.Value) {
	c.functions[signature].value = val
}

func (c *Context) signature(funcName string, exprs ...Expr) string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(funcName)
	buf.WriteRune('(')
	for i, e := range exprs {
		e.Type(c)
		buf.WriteString(e.Type(c))
		if len(exprs)-1 > i {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune(')')
	return buf.String()
}

// builtInOperation assumes signature is found already
func (c *Context) builtInOperation(signature string, args []llvm.Value) llvm.Value {
	switch signature {
	case "+(i32,i32)":
		fallthrough
	case "+(i64,i64)":
		return c.Builder.CreateAdd(args[0], args[1], "")
	case "-(i32,i32)":
		fallthrough
	case "-(i64,i64)":
		return c.Builder.CreateSub(args[0], args[1], "")
	case "*(i32,i32)":
		fallthrough
	case "*(i64,i64)":
		return c.Builder.CreateMul(args[0], args[1], "")
	case "/(i32,i32)":
		fallthrough
	case "/(i64,i64)":
		return c.Builder.CreateMul(args[0], args[1], "")
	case "==(i32,i32)":
		fallthrough
	case "==(i64,i64)":
		return c.Builder.CreateICmp(
			llvm.IntEQ,
			args[0],
			args[1],
			"",
		)
	case "+(f32,f32)":
		fallthrough
	case "+(f64,f64)":
		return c.Builder.CreateFAdd(args[0], args[1], "")
	case "-(f32,f32)":
		fallthrough
	case "-(f64,f64)":
		return c.Builder.CreateSub(args[0], args[1], "")
	case "*(f32,f32)":
		fallthrough
	case "*(f64,f64)":
		return c.Builder.CreateMul(args[0], args[1], "")
	case "/(f32,f32)":
		fallthrough
	case "/(f64,f64)":
		return c.Builder.CreateMul(args[0], args[1], "")
	case "==(f32,f32)":
		fallthrough
	case "==(f64,f64)":
		return c.Builder.CreateFCmp(
			llvm.FloatOEQ,
			args[0],
			args[1],
			"",
		)
	default:
		panic("Compiler bug at Context::Call, builtInOperation assumes signature is built in operation")
	}
}

func (c *Context) Call(funcName string, exprs ...Expr) llvm.Value {
	signature := c.signature(funcName, exprs...)

	args := []llvm.Value{}
	for _, e := range exprs {
		args = append(args, e.Codegen(c))
	}

	if _, ok := c.builtInOperators[signature]; ok {
		return c.builtInOperation(signature, args)
	}

	if c.funcRetTyp(signature) != nil {
		fn := c.funcRetTyp(signature).value
		return c.Builder.CreateCall(fn, args, "")
	}
	panic(fmt.Sprintf("`%s`, No this function or operator in current scope", signature))
}

func (c *Context) funcRetTyp(signature string) *Function {
	if retT, ok := c.builtInOperators[signature]; ok {
		return &Function{
			retType: retT,
		}
	}
	if f, ok := c.functions[signature]; ok {
		return f
	}
	if c.Parent == nil {
		return nil
	}
	return c.Parent.funcRetTyp(signature)
}

func (c *Context) VarValue(name string, value llvm.Value) {
	c.Vars[name] = value
}

func (c *Context) NewVar(name string, typ string) {
	// FIXME: let vars contains Var Node only, then Var Node contains more info is better.
	c.VarsType[name] = typ
	// FIXME: Missing export & mutable or not info
}

func (c *Context) Var(name string) (llvm.Value, bool) {
	v, ok := c.Vars[name]
	if ok {
		return v, true
	}
	if c.Parent != nil {
		return c.Parent.Var(name)
	}
	return llvm.Value{}, false
	// It will cause easy panic in llvm system
	// To match the type have to write down this line
	// p.s. Because var not found is common, we can't panic this
}

func (c *Context) VarType(name string) (string, bool) {
	v, ok := c.VarsType[name]
	if ok {
		return v, true
	}
	if c.Parent != nil {
		return c.Parent.VarType(name)
	}
	return "no this var", false
}
