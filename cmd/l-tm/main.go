package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/mewmew/l-tm/asm/ll"
)

func main() {
	flag.Parse()
	for _, llPath := range flag.Args() {
		fmt.Println("path:", llPath)
		buf, err := ioutil.ReadFile(llPath)
		if err != nil {
			log.Fatalf("%q: %+v", llPath, err)
		}
		source := string(buf)
		start := time.Now()
		lex := &ll.Lexer{}
		lex.Init(source)
		listener := func(t ll.NodeType, offset, endoffset int) {
			fmt.Printf("t: %v (start=%v, end=%v)\n", t, offset, endoffset)
		}
		p := &ll.Parser{}
		p.Init(listener)
		if err := p.Parse(lex); err != nil {
			log.Fatalf("%q: %+v", llPath, err)
		}
		fmt.Println("took:", time.Since(start))
		fmt.Println()
	}
}
