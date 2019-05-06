.PHONY: build test clean
build:
	clang -c list.c
test:
	clang++ -g *.c -lgtest -lgtest_main
	./a.out
clean:
	rm -f *.out
	rm -f *.o
