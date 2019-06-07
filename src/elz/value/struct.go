package value

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"
	"github.com/elz-lang/elz/src/irutil"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmtypes "github.com/llir/llvm/ir/types"
	llvmvalue "github.com/llir/llvm/ir/value"
)

type Struct struct {
	// just for complete interface
	llvmvalue.Value
	// just for putting it into binding as return expression
	ast.Expr
	*ast.TypeDefine
	initBlock  *ir.Block
	mallocFunc *ir.Func
	modName    string
	mod        *ir.Module

	cacheT *types.Struct
}

func NewStruct(mod *ir.Module, modName string, mallocFunc *ir.Func, typeDef *ast.TypeDefine) *Struct {
	return &Struct{
		TypeDefine: typeDef,
		mallocFunc: mallocFunc,
		mod:        mod,
		modName:    modName,
	}
}

func (s *Struct) Constructor() *ir.Func {
	typName := s.Name
	if s.modName != "" {
		typName = fmt.Sprintf("%s::%s", s.modName, s.Name)
	}
	// emit new type
	structT := s.mod.NewTypeDef(
		typName,
		s.ElzType().LLVMType(),
	)
	// create constructor
	params := make([]*ir.Param, 0)
	for _, field := range s.Fields {
		params = append(params,
			ir.NewParam(field.Name, types.FromString(field.Type.(*ast.ExistType).Name).LLVMType()),
		)
	}
	constructor := s.mod.NewFunc(typName, structT, params...)
	b := constructor.NewBlock("")
	structMalloc := b.NewCall(s.mallocFunc, constant.NewInt(llvmtypes.I64, irutil.SizeOf(structT)))
	casted := b.NewBitCast(structMalloc, llvmtypes.NewPointer(structT))
	for i, _ := range params {
		to := b.NewGetElementPtr(casted,
			constant.NewInt(llvmtypes.I64, 0),
			constant.NewInt(llvmtypes.I32, int64(i)),
		)
		b.NewStore(params[i], to)
	}
	b.NewRet(b.NewLoad(casted))
	return constructor
}

func (s *Struct) ElzType() *types.Struct {
	if s.cacheT == nil {
		s.cacheT = &types.Struct{
			TypeDefine: s.TypeDefine,
		}
	}
	return s.cacheT
}
