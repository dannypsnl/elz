#include <stdint.h>
#include <stdlib.h>

struct List {
  void *head;
  struct List *tail;
};

// APIs
struct List *new_list(int size, void **elems);
void for_each(struct List *list, void f(void *));
uint64_t length(struct List *list);
void delete_list(struct List *list);
