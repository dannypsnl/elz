#include "./expr.h"
#include "llvm/Support/raw_ostream.h"
#include "gtest/gtest.h"

TEST(CodegenOfAST, Number) {
  ast::Number num{10};
  ast::TheModule->print(llvm::errs(), nullptr);
}
