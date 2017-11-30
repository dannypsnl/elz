#pragma once

#include "ElzBaseListener.h" // parser/ElzBaseListener.h
#include "ElzParser.h"       // parser/ElzParser.h

#include "./ast/expr.h"

namespace elz {
class AstBuilder : public ElzBaseListener {
public:
  // TODO: Put context at here, and improve it.
  void enterVarDefine(ElzParser::DefineContext *ctx) { ; }
};
} // namespace elz
