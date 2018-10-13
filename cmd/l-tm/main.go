package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/mewmew/l-tm/parser"
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
		l := &parser.Lexer{}
		l.Init(source)
		p := &parser.Parser{}
		p.Init()
		if err := p.Parse(l); err != nil {
			log.Fatalf("%q: %+v", llPath, err)
		}
		fmt.Println("took:", time.Since(start))
		fmt.Println()
	}
}
