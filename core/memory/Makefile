all: install

.PHONY: install build test clean
install: build
	@echo "installing lib: memory..."
	@cp memory.o $(HOME)/.elz/core/
	@echo "lib: memory installed"
build: memory.o
test:
	@clang++ -g -c *.cc
	@clang -g -c *.c
	@clang++ -g *.o -lgtest -lgtest_main
	@./a.out
clean:
	@rm -f *.out
	@rm -f *.o

memory.o:
	@clang -c memory.c
