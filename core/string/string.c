#include "string.h"

#include <stdlib.h>
#include <stdio.h>

struct String *new_string(int size, char *str) {
  struct String *head = (struct String *)malloc(sizeof(struct String));
  struct String str_v = {
      .c = *(str + 0),
      .tail = NULL,
  };
  *head = str_v;

  struct String *next = head;
  int index = 0;
  while (index < size) {
    struct String str_v = {
        .c = *(str + index),
        .tail = NULL,
    };
    next->tail = (struct String *)malloc(sizeof(struct String));
    *(next->tail) = str_v;
    next = next->tail;
    index++;
  }
  return head;
}

uint64_t string_length(struct String *string) {
  uint64_t count = -1;
  struct String **next;
  for (next = &string; *next != NULL; *next = (*next)->tail) {
    count++;
  }
  return count;
}

char string_index(struct String *string, uint64_t index) {
  uint64_t count = 0;
  struct String **next;
  for (next = &string; *next != NULL; *next = (*next)->tail) {
    if (count == index + 1) {
      return (*next)->c;
    }
    count++;
  }
  return '\0';
}

void delete_string(struct String *string) {
  struct String *next = string;
  while (next != NULL) {
    struct String *prev = next;
    next = next->tail;
    free(prev);
  }
}

void elz_println(struct String *str) {
  char *cstr = (char *)malloc(sizeof(char) * string_length(str));
  printf(cstr);
  printf("\n");
  delete[] cstr;
}