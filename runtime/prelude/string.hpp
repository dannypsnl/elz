//
// Created by dannypsnl on 2019/11/29.
//

#ifndef ELZ_LIB_RUNTIME_PRELUDE_STRING_H_
#define ELZ_LIB_RUNTIME_PRELUDE_STRING_H_

#include <string>

class String {
 public:
  std::string value;
  String(char *lit) : value(lit) {}
};

#endif //ELZ_LIB_RUNTIME_PRELUDE_STRING_H_
