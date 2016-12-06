all: sqlite.ll

%.ll: %.c
	clang -S -emit-llvm -o $@ $<

sqlite.ll: sqlite3.ll
	cp $< $@
	cat main.ll >> $@
	../strip.sh $@
