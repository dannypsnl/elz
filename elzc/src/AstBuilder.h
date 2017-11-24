#pragma once

#include "ElzBaseListener.h" // parser/ElzBaseListener.h
#include "ElzParser.h"       // parser/ElzParser.h

namespace elz {
class AstBuilder : public ElzBaseListener {
  void enterVarDefine(ElzParser::DefineContext *ctx) {}
};
} // namespace elz
