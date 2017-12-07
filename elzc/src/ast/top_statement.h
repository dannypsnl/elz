#pragma once

#include "llvm/ADT/APFloat.h"
#include "llvm/IR/Constants.h"
#include "llvm/IR/GlobalValue.h"
#include "llvm/IR/GlobalVariable.h"

#include <string>

using llvm::ConstantFP;
using llvm::GlobalValue;
using llvm::GlobalVariable;

namespace ast {

class VariableDefine {
  std::string name;
  std::string type;

public:
  VariableDefine(std::string name) : name{name} {}
  void codegen() {
    TheModule->getOrInsertGlobal(name, Builder.getInt32Ty());
    GlobalVariable *gvar = TheModule->getNamedGlobal(name);
    gvar->setInitializer(
        ConstantFP::get(TheModule->getContext(), llvm::APFloat(3.1415926)));
    gvar->setLinkage(GlobalValue::CommonLinkage);
  }
};

} // namespace ast
