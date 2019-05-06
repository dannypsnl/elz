#include "list.h"

#include <stdio.h>

#include "gtest/gtest.h"

void print_element(void *element) { printf("element: %d\n", *(int *)element); }

TEST(List, test_new_list) {
  int one = 1;
  int two = 2;
  void *init[] = {&one, &two};
  struct List *list = new_list(2, init);
  ASSERT_EQ(2, length(list));
  delete_list(list);
}
