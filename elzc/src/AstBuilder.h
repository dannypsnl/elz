#pragma once

#include "ElzBaseListener.h" // parser/ElzBaseListener.h
#include "ElzParser.h"       // parser/ElzParser.h

#include "./ast/expr.h"

namespace elz {
class AstBuilder : public ElzBaseListener {
  void enterVarDefine(ElzParser::DefineContext *ctx) {}
};
} // namespace elz
