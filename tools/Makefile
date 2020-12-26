all: string2enum stringer

stringer:
	go install golang.org/x/tools/cmd/stringer

string2enum:
	go build github.com/mewspring/tools/cmd/string2enum

.PHONY: all clean

clean:
	$(RM) -v string2enum
