GODIR=`echo ${GOPATH} | awk '{split($$0, arr, ":"); print arr[1]}'`
TM_DIR=${GODIR}/src/github.com/inspirer/textmapper

all: gen

gen: ll.tm
	@${TM_DIR}/tm-tool/libs/textmapper.sh $<
	@go fmt ./... > /dev/null
	@go install ./... > /dev/null

clean:
	$(RM) -v listener.go lexer.go lexer_tables.go parser.go parser_tables.go token.go
	$(RM) -rf -v ast/
	$(RM) -rf -v selector/

.PHONY: all gen clean
