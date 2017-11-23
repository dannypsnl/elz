#include <stdio.h>

extern double nestedinstruction();

int main() { printf("Our Global value: %f\n", nestedinstruction()); }
