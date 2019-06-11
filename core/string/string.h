#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>

struct String {
  char c;
  struct String *tail;
};

// APIs
struct String *new_string(int size, char *c_str);
uint64_t string_length(struct String *string);
char string_index(struct String *string, uint64_t index);
void delete_string(struct String *string);

#ifdef __cplusplus
} // extern "C"
#endif
