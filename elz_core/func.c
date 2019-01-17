#include <stdio.h>
#include <stdlib.h>

typedef struct elz_func_struct {
  int counter;
  void **parameters;
} elz_func;

elz_func *new_elz_func(int parameter_length) {
  elz_func *ret = (elz_func *)malloc(sizeof(elz_func));
  struct elz_func_struct f = {
      .counter = 0,
      .parameters = (void **)malloc(sizeof(void *) * parameter_length),
  };
  *ret = f;
  return ret;
}

void elz_func_append_argument(elz_func *f, void *arg) {
  *(f->parameters + f->counter) = arg;
  f->counter++;
}

int add(int x, int y) { return x + y; }

int main() {
  elz_func *add_func = new_elz_func(2);
  int x = 1, y = 2;
  // Compiler can handle how to store arguments
  elz_func_append_argument(add_func, (void *)&x);
  elz_func_append_argument(add_func, (void *)&y);
  if (add_func->counter == 2) {
    // Compiler should know how many parameters function would take and the
    // types of them, so this can be handle by compiler
    int x = *(int *)add_func->parameters[0];
    int y = *(int *)add_func->parameters[1];
    int result = add(x, y);
    printf("result: %d\n", result);
  }
}
