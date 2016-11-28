# *.c -> *.ll
CFILES = $(wildcard *.c)
LLFILES = $(CFILES:.c=.ll)

all: sar $(LLFILES)

%.ll: %.c
	clang -S -emit-llvm -o $@ $<
	opt -S --mem2reg -o $@ $@
	./strip.sh $@

sar:
	@if ! which $@ > /dev/null 2>&1 ; then \
		echo "Please install the \"sar\" tool"; \
		echo; \
		echo "   go get github.com/mewkiz/cmd/sar"; \
		echo; \
		exit 1; \
	fi

.PHONY: all sar
