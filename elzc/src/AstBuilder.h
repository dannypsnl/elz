#pragma once

#include "ElzBaseListener.h" // parser/ElzBaseListener.h
#include "ElzParser.h"       // parser/ElzParser.h

#include "./ast/expr.h"
#include "./ast/top_statement.h"
#include <iostream>

namespace elz {
class AstBuilder : public ElzBaseListener {
public:
  // TODO: Put context at here, and improve it.
  void enterDefine(ElzParser::DefineContext *ctx) override {
    ast::VariableDefine new_var{ctx->ID()->getText()};
    std::cout << "new var: " << ctx->ID()->getText() << std::endl;
  }
};
} // namespace elz
