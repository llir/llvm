all: gen

gen: ll.bnf
	gocc $<
	find . -type f -name '*.go' | xargs goimports -w

debug_lexer: ll.bnf
	gocc -debug_lexer -v $<
	find . -type f -name '*.go' | xargs goimports -w

debug_parser: ll.bnf
	gocc -debug_parser -v $<
	find . -type f -name '*.go' | xargs goimports -w

clean:
	rm -f errors/errors.go
	rm -f lexer/acttab.go
	rm -f lexer/lexer.go
	rm -f lexer/transitiontable.go
	rm -f parser/action.go
	rm -f parser/actiontable.go
	rm -f parser/gototable.go
	rm -f parser/parser.go
	rm -f parser/productionstable.go
	rm -f token/token.go
	rm -f util/litconv.go
	rm -f util/rune.go
	-rmdir --ignore-fail-on-non-empty errors
	-rmdir --ignore-fail-on-non-empty lexer
	-rmdir --ignore-fail-on-non-empty parser
	-rmdir --ignore-fail-on-non-empty token
	-rmdir --ignore-fail-on-non-empty util
	rm -f terminals.txt LR1_conflicts.txt LR1_sets.txt first.txt lexer_sets.txt

.PHONY: gen debug_lexer debug_parser clean
