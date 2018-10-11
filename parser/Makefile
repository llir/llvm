TM_DIR=/home/u/goget/src/github.com/inspirer/textmapper

all: gen

gen: ll.tm
	java -jar ${TM_DIR}/tm-tool/libs/textmapper-0.9.21.jar $<
	go fmt

clean:
	$(RM) -v lexer.go lexer_tables.go parser.go parser_tables.go token.go

.PHONY: all gen clean
