#include "list.h"

#include "gtest/gtest.h"

TEST(List, test_new_list) {
  int one = 1;
  int two = 2;
  void *init[] = {&one, &two};
  struct List *list = new_list(2, init);
  ASSERT_EQ(2, list_length(list));
  delete_list(list);
}

TEST(List, test_list_index) {
  int one = 1;
  int two = 2;
  void *init[] = {&one, &two};
  struct List *list = new_list(2, init);
  void *first = list_index(list, 0);
  ASSERT_EQ(1, *(int *)first);
  void *second = list_index(list, 1);
  ASSERT_EQ(2, *(int *)second);
  delete_list(list);
}