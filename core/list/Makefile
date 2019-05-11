all: install

.PHONY: install build test clean
install: build
	cp list.o $(HOME)/.elz/core/
build: list.o
test:
	clang++ -g -c *.cc
	clang -g -c *.c
	clang++ -g *.o -lgtest -lgtest_main
	./a.out
clean:
	rm -f *.out
	rm -f *.o

list.o:
	clang -c list.c
