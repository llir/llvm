# *.c -> *.ll
CFILES = $(wildcard *.c)
LLFILES = $(CFILES:.c=.ll)

CFLAGS?=-O0

all: $(LLFILES)

%.ll: %.c
	clang $(CFLAGS) -S -emit-llvm -o $@ $<
