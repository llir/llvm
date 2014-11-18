prepare:
	clang -S -emit-llvm -o for.ll for.c
	opt -mem2reg -S -o for.ll for.ll
	# Manual cleanup of for.ll
