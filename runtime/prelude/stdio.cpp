#include "string.hpp"
#include <iostream>

extern "C" {
void println(String *s) { std::cout << s->value << std::endl; }
}
