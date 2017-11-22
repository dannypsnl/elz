template <typename T> class Test {
public:
  T element;
};

template <> class Test<double> {
public:
  char *element;
};

Test<int> t{1};
Test<double> t2{};

template <typename T> void foo(T) {}
void foo(int) {}
