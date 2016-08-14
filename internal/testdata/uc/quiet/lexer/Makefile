# *.c -> *.ll
CFILES = $(wildcard *.c)
LLFILES = $(CFILES:.c=.ll)

all: $(LLFILES)

%.ll: %.c
	clang -S -emit-llvm -o $@ $<
