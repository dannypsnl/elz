#include <stdlib.h>

struct string {
  int len;
  char *raw_str;
};

void string_push(struct string *s, char c) {
  ++s->len;
  char *original_str = s->raw_str;
  s->raw_str = malloc(s->len);
  s->raw_str = original_str;
  s->raw_str[s->len - 1] = c;
  s->raw_str[s->len] = '\0';
}

int main() {
  int a = 0;
  int *b = &a;
  int c = *b;

  struct string dan = {4, "dan"};
  struct string *ref_dan = &dan;
  string_push(ref_dan, 'n');
  string_push(ref_dan, 'y');
}
