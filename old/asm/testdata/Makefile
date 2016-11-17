all: for c4

for:
	clang -S -emit-llvm -o for.ll for.c
	opt -mem2reg -S -o for.ll for.ll

c4:
	clang -S -emit-llvm -o c4.ll c4.c
	opt -mem2reg -S -o c4.ll c4.ll

prepare:
	wget -O c4.c https://raw.githubusercontent.com/rswier/c4/master/c4.c
