#include "./expr.h"
#include "./top_statement.h"
#include "llvm/Support/raw_ostream.h"
#include "gtest/gtest.h"

TEST(CodegenOfAST, Number) {
  ast::Number num{10};
  ast::VariableDefine v{"foo"};
  v.codegen();
  ast::TheModule->print(llvm::errs(), nullptr);
}
