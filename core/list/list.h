#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>

struct List {
  void *head;
  struct List *tail;
};

// APIs
struct List *new_list(int size, void **elems);
uint64_t list_length(struct List *list);
void *list_index(struct List *list, uint64_t index);
void delete_list(struct List *list);

#ifdef __cplusplus
} // extern "C"
#endif
