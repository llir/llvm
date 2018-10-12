TM_DIR=/home/u/goget/src/github.com/inspirer/textmapper

all: gen

gen: ll.tm
	@java -jar ${TM_DIR}/tm-tool/libs/textmapper-0.9.21.jar $<
	@go fmt > /dev/null

clean:
	$(RM) -v listener.go lexer.go lexer_tables.go parser.go parser_tables.go token.go
	$(RM) -rf -v ast/
	$(RM) -rf -v selector/

.PHONY: all gen clean
