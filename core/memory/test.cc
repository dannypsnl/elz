#include "memory.h"

#include "gtest/gtest.h"

TEST(Memory, test_elz_malloc) {
  char *str = (char *)elz_malloc(5);
  strcpy(str, "abcd");
  ASSERT_EQ(*(str), 'a');
  ASSERT_EQ(*(str + 1), 'b');
  ASSERT_EQ(*(str + 2), 'c');
  ASSERT_EQ(*(str + 3), 'd');
}
