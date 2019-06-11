all: install

.PHONY: install build test clean
install: build
	@echo "installing lib: string..."
	@cp string.o $(HOME)/.elz/core/
	@echo "lib: string installed"
build: string.o
test:
	@clang++ -std=c++11 -g -c *.cc
	@clang -g -c *.c
	@clang++ -g *.o -lgtest -lgtest_main
	@./a.out
clean:
	@rm -f *.out
	@rm -f *.o

string.o:
	@clang -c string.c
