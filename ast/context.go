package ast

import (
	"bytes"
	"fmt"

	"github.com/elz-lang/elz/collection/stack"
	"github.com/elz-lang/elz/errors"

	"llvm.org/llvm/bindings/go/llvm"
)

var builtInOps = map[string]string{
	"+(i32,i32)":  "i32",
	"+(i64,i64)":  "i64",
	"-(i32,i32)":  "i32",
	"-(i64,i64)":  "i64",
	"*(i32,i32)":  "i32",
	"*(i64,i64)":  "i64",
	"/(i32,i32)":  "i32",
	"/(i64,i64)":  "i64",
	"==(i32,i32)": "bool",
	"==(i64,i64)": "bool",
	"<(i32,i32)":  "bool",
	"<(i64,i64)":  "bool",
	"<=(i32,i32)": "bool",
	"<=(i64,i64)": "bool",
	">(i32,i32)":  "bool",
	">(i64,i64)":  "bool",
	">=(i32,i32)": "bool",
	">=(i64,i64)": "bool",

	"+(f32,f32)":  "f32",
	"+(f64,f64)":  "f64",
	"-(f32,f32)":  "f32",
	"-(f64,f64)":  "f64",
	"*(f32,f32)":  "f32",
	"*(f64,f64)":  "f64",
	"/(f32,f32)":  "f32",
	"/(f64,f64)":  "f64",
	"==(f32,f32)": "bool",
	"==(f64,f64)": "bool",
}

type elzTypeField struct {
	name, typ string
	export    bool
	offset    int
}
type elzType struct {
	llvmType llvm.Type
	typeName string
	// llvm use offset instead of field name to save fields
	field map[string]*elzTypeField // offset map of field, fieldName => offset
}

func newElzType(typeName string, t llvm.Type, attrs []TypeAttr) *elzType {
	newT := &elzType{
		typeName: typeName,
		llvmType: t,
		field:    make(map[string]*elzTypeField),
	}

	for i, a := range attrs {
		newT.field[a.Name] = &elzTypeField{name: a.Name, typ: a.Typ, export: a.Export, offset: i}
	}

	return newT
}

func NewContext() *Context {
	c := &Context{
		Parent:    nil,
		Reporter:  errors.NewReporter(),
		Module:    llvm.NewModule("main"),
		Context:   llvm.NewContext(),
		variables: make(map[string]llvm.Value),
		VarsType:  make(map[string]string),
		types:     make(map[string]*elzType),
		Builder:   llvm.NewBuilder(),

		breaks:    stack.New(),
		functions: make(map[string]*Function),
	}

	return c
}

type Context struct {
	Parent    *Context
	Reporter  *errors.Reporter
	Module    llvm.Module
	Context   llvm.Context
	variables map[string]llvm.Value
	VarsType  map[string]string
	types     map[string]*elzType
	Builder   llvm.Builder

	breaks    *stack.Stack
	functions map[string]*Function
}

type Function struct {
	value   llvm.Value
	retType string
}

func (c *Context) Type(elzTypeName string) llvm.Type {
	typ := LLVMType(elzTypeName)
	if typ.String() != "VoidType" {
		return typ
	}
	if elzTypeName == "()" {
		return llvm.VoidType()
	}
	return llvm.PointerType(c.Module.GetTypeByName(elzTypeName), 0)
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

func (c *Context) calcSignature(funcName string, exprs ...Expr) string {
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
	case ">(i32,i32)":
		fallthrough
	case ">(i64,i64)":
		return c.Builder.CreateICmp(
			llvm.IntSGT,
			args[0],
			args[1],
			"",
		)
	case ">=(i32,i32)":
		fallthrough
	case ">=(i64,i64)":
		return c.Builder.CreateICmp(
			llvm.IntSGE,
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
	signature := c.calcSignature(funcName, exprs...)

	args := make([]llvm.Value, 0)
	for _, e := range exprs {
		args = append(args, e.Codegen(c))
	}

	if _, ok := builtInOps[signature]; ok {
		return c.builtInOperation(signature, args)
	}

	if c.Func(signature) != nil {
		fn := c.Func(signature).value
		return c.Builder.CreateCall(fn, args, "")
	}
	panic(fmt.Sprintf("`%s`, No this function or operator in current scope", signature))
}

func (c *Context) Func(signature string) *Function {
	// check is it is the built-in operator?
	if retT, ok := builtInOps[signature]; ok {
		return &Function{
			retType: retT,
		}
	}
	// Have function in this context
	if f, ok := c.functions[signature]; ok {
		return f
	}
	// if it's root & not found ~> global not found
	if c.Parent == nil {
		return nil
	}
	// search to parent
	return c.Parent.Func(signature)
}

func (c *Context) VarValue(name string, value llvm.Value) {
	c.variables[name] = value
}

func (c *Context) NewVar(name string, typ string) {
	// FIXME: let vars contains Var Node only, then Var Node contains more info is better.
	c.VarsType[name] = typ
	// FIXME: Missing export & mutable or not info
}

func (c *Context) LLVMValueOfVar(name string) (llvm.Value, bool) {
	v, ok := c.variables[name]
	if ok {
		return v, true
	}
	if c.Parent != nil {
		return c.Parent.LLVMValueOfVar(name)
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

func (c *Context) NewType(name string, t llvm.Type, attrs []TypeAttr) {
	if c.Parent != nil {
		panic("Elz only allow you add type at root scope")
	}
	c.types[name] = newElzType(name, t, attrs)
}

func (c *Context) FindFieldIndexOfType(typeName, fieldName string) int {
	root := c
	for root.Parent != nil {
		root = c.Parent
	}
	t := root.types[typeName]
	return t.field[fieldName].offset
}
func (c *Context) FieldTypeOf(typeName, fieldName string) string {
	root := c
	for root.Parent != nil {
		root = c.Parent
	}
	return root.types[typeName].field[fieldName].typ
}
