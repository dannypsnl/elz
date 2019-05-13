all: install

.PHONY: install build test clean
install: build
	@echo "installing lib: list..."
	@cp list.o $(HOME)/.elz/core/
	@echo "lib: list installed"
build: list.o
test:
	@clang++ -g -c *.cc
	@clang -g -c *.c
	@clang++ -g *.o -lgtest -lgtest_main
	@./a.out
clean:
	@rm -f *.out
	@rm -f *.o

list.o:
	@clang -c list.c
