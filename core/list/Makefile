.PHONY: build test clean
build:
	clang -c list.c
test:
	clang++ -g -c *.cc
	clang -g -c *.c
	clang++ -g *.o -lgtest -lgtest_main
	./a.out
clean:
	rm -f *.out
	rm -f *.o
