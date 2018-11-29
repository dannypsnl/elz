package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Builder interface {
	NewAShr(x, y value.Value) *ir.InstAShr
	NewAdd(x, y value.Value) *ir.InstAdd
	NewAddrSpaceCast(from value.Value, to types.Type) *ir.InstAddrSpaceCast
	NewAlloca(elemType types.Type) *ir.InstAlloca
	NewAnd(x, y value.Value) *ir.InstAnd
	NewAtomicRMW(op enum.AtomicOp, dst, x value.Value, ordering enum.AtomicOrdering) *ir.InstAtomicRMW
	NewBitCast(from value.Value, to types.Type) *ir.InstBitCast
	NewBr(target *ir.BasicBlock) *ir.TermBr
	NewCall(callee value.Value, args ...value.Value) *ir.InstCall
	NewCatchPad(scope *ir.TermCatchSwitch, args ...value.Value) *ir.InstCatchPad
	NewCatchRet(from *ir.InstCatchPad, to *ir.BasicBlock) *ir.TermCatchRet
	NewCatchSwitch(scope ir.ExceptionScope, handlers []*ir.BasicBlock, unwindTarget ir.UnwindTarget) *ir.TermCatchSwitch
	NewCleanupPad(scope ir.ExceptionScope, args ...value.Value) *ir.InstCleanupPad
	NewCleanupRet(from *ir.InstCleanupPad, to ir.UnwindTarget) *ir.TermCleanupRet
	NewCmpXchg(ptr, cmp, new value.Value, successOrdering, failureOrdering enum.AtomicOrdering) *ir.InstCmpXchg
	NewCondBr(cond value.Value, targetTrue, targetFalse *ir.BasicBlock) *ir.TermCondBr
	NewExtractElement(x, index value.Value) *ir.InstExtractElement
	NewExtractValue(x value.Value, indices ...int64) *ir.InstExtractValue
	NewFAdd(x, y value.Value) *ir.InstFAdd
	NewFCmp(pred enum.FPred, x, y value.Value) *ir.InstFCmp
	NewFDiv(x, y value.Value) *ir.InstFDiv
	NewFMul(x, y value.Value) *ir.InstFMul
	NewFPExt(from value.Value, to types.Type) *ir.InstFPExt
	NewFPToSI(from value.Value, to types.Type) *ir.InstFPToSI
	NewFPToUI(from value.Value, to types.Type) *ir.InstFPToUI
	NewFPTrunc(from value.Value, to types.Type) *ir.InstFPTrunc
	NewFRem(x, y value.Value) *ir.InstFRem
	NewFSub(x, y value.Value) *ir.InstFSub
	NewFence(ordering enum.AtomicOrdering) *ir.InstFence
	NewGetElementPtr(src value.Value, indices ...value.Value) *ir.InstGetElementPtr
	NewICmp(pred enum.IPred, x, y value.Value) *ir.InstICmp
	NewIndirectBr(addr constant.Constant, validTargets ...*ir.BasicBlock) *ir.TermIndirectBr
	NewInsertElement(x, elem, index value.Value) *ir.InstInsertElement
	NewInsertValue(x, elem value.Value, indices ...int64) *ir.InstInsertValue
	NewIntToPtr(from value.Value, to types.Type) *ir.InstIntToPtr
	NewInvoke(invokee value.Value, args []value.Value, normal, exception *ir.BasicBlock) *ir.TermInvoke
	NewLShr(x, y value.Value) *ir.InstLShr
	NewLandingPad(resultType types.Type, clauses ...*ir.Clause) *ir.InstLandingPad
	NewLoad(src value.Value) *ir.InstLoad
	NewMul(x, y value.Value) *ir.InstMul
	NewOr(x, y value.Value) *ir.InstOr
	NewPhi(incs ...*ir.Incoming) *ir.InstPhi
	NewPtrToInt(from value.Value, to types.Type) *ir.InstPtrToInt
	NewResume(x value.Value) *ir.TermResume
	NewRet(x value.Value) *ir.TermRet
	NewSDiv(x, y value.Value) *ir.InstSDiv
	NewSExt(from value.Value, to types.Type) *ir.InstSExt
	NewSIToFP(from value.Value, to types.Type) *ir.InstSIToFP
	NewSRem(x, y value.Value) *ir.InstSRem
	NewSelect(cond, x, y value.Value) *ir.InstSelect
	NewShl(x, y value.Value) *ir.InstShl
	NewShuffleVector(x, y, mask value.Value) *ir.InstShuffleVector
	NewStore(src, dst value.Value) *ir.InstStore
	NewSub(x, y value.Value) *ir.InstSub
	NewSwitch(x value.Value, targetDefault *ir.BasicBlock, cases ...*ir.Case) *ir.TermSwitch
	NewTrunc(from value.Value, to types.Type) *ir.InstTrunc
	NewUDiv(x, y value.Value) *ir.InstUDiv
	NewUIToFP(from value.Value, to types.Type) *ir.InstUIToFP
	NewURem(x, y value.Value) *ir.InstURem
	NewUnreachable() *ir.TermUnreachable
	NewVAArg(vaList value.Value, argType types.Type) *ir.InstVAArg
	NewXor(x, y value.Value) *ir.InstXor
	NewZExt(from value.Value, to types.Type) *ir.InstZExt
}
