#pragma once

#include "llvm/ADT/APFloat.h"
#include "llvm/IR/Constants.h"
#include "llvm/IR/IRBuilder.h"
#include "llvm/IR/LLVMContext.h"
#include "llvm/IR/Module.h"
#include "llvm/IR/Value.h"
#include <iostream>

#include <map>
#include <memory>
#include <string>

using llvm::ConstantFP;
using llvm::IRBuilder;
using llvm::LLVMContext;
using llvm::Module;
using llvm::Value;

namespace ast {

static LLVMContext TheContext;
static IRBuilder<> Builder(TheContext);
static std::unique_ptr<Module> TheModule =
    llvm::make_unique<Module>("main", TheContext);
static std::map<std::string, Value *> NamedValues;

class Expr {
public:
  virtual ~Expr() {}
  virtual Value *codegen() = 0;
};

class Number : public Expr {
  double Val;

public:
  // Maybe should use string val and transport inside?
  Number(double val) : Val{val} {}
  virtual Value *codegen() {
    return ConstantFP::get(TheContext, llvm::APFloat(Val));
  }
};

class Variable : public Expr {
  std::string Val;

public:
  Variable(std::string val) : Val{val} {}
  virtual Value *codegen() {
    Value *V = NamedValues[Val];
    if (!V) {
      std::cout << "We don't have a variable call[" << Val << ']' << std::endl;
    }
    return V;
  }
};

class BinaryExpr : public Expr {};

} // namespace ast
