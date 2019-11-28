#include "string.hpp"

extern "C" {
String *new_string(char *lit) { return new String(lit); }
}
