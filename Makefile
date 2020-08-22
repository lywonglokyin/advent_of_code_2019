GCC := g++ -Wall -std=c++0x
DIR := bin build

target_cpps := $(wildcard src/day*.cpp)
target_binary := $(patsubst src/%.cpp, bin/%, $(target_cpps))


all: $(target_binary)

$(target_binary): $(target_cpps) build/utils.o | $(DIR)
	$(GCC) -o $@ $(patsubst bin/%, src/%.cpp, $@) build/utils.o

build/utils.o: | $(DIR)
	$(GCC) -c src/utils.cpp -o build/utils.o

$(DIR):
	mkdir $(DIR)

.PHONY: clean
clean :
	rm -f bin/*
	rm -f build/utils.o