#include "string.h"

#include <string>

#include "gtest/gtest.h"

TEST(String, test_new_string) {
  std::string s{"abc"};
  char *cstr = new char[s.length() + 1];
  std::strcpy(cstr, s.c_str());
  struct String *string = new_string(s.length(), cstr);
  ASSERT_EQ(3, string_length(string));
  delete_string(string);
  delete[] cstr;
}

TEST(String, test_string_index) {
  std::string s{"abc"};
  char *cstr = new char[s.length() + 1];
  std::strcpy(cstr, s.c_str());
  struct String *string = new_string(s.length(), cstr);
  ASSERT_EQ('a', string_index(string, 0));
  ASSERT_EQ('b', string_index(string, 1));
  ASSERT_EQ('c', string_index(string, 2));
  delete[] cstr;
}
