#include "list.h"

#include <stdlib.h>

struct List *empty_list() {
  struct List *list_ptr = (struct List *)malloc(sizeof(struct List));
  struct List list_v = {
      .head = NULL,
      .tail = NULL,
  };
  *list_ptr = list_v;
  return list_ptr;
}

struct List *push(struct List *list, void *element) {
  list->tail = empty_list();
  list->head = element;
  return list->tail;
}

struct List *new_list(int size, void **values) {
  struct List *head = empty_list();
  struct List *next = head;
  int index = 0;
  while (index < size) {
    next = push(next, *(values + index));
    index++;
  }
  return head;
}

uint64_t list_length(struct List *list) {
  struct List *next = list;
  uint64_t count = 0;
  while (next != NULL && next->head != NULL && next->tail != NULL) {
    count++;
    next = next->tail;
  }
  return count;
}

void for_each(struct List *list, void f(void *)) {
  struct List *next = list;
  while (next != NULL && next->head != NULL) {
    f(next->head);
    next = next->tail;
  }
}

void delete_list(struct List *list) {
  struct List *next = list;
  while (next != NULL) {
    struct List *prev = next;
    next = next->tail;
    free(prev);
  }
}
