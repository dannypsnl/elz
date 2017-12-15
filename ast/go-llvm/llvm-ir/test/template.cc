template <typename T> class Test {
public:
  T element;
};

// template <> class Test<double> {
// public:
// const char *element;
//};

// const char *c = "hello llvm";

Test<int> t{1};
Test<float> t1{1};
Test<double> t2{3.14};

template <typename T> void foo(T) {}
void foo(int) {}
