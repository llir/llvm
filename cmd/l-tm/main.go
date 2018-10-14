package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/kr/pretty"
	"github.com/mewmew/l-tm/asm"
)

func main() {
	flag.Parse()
	for _, llPath := range flag.Args() {
		fmt.Println("path:", llPath)
		start := time.Now()
		module, err := asm.ParseFile(llPath)
		if err != nil {
			log.Fatalf("%q: %+v", llPath, err)
		}
		fmt.Println("took:", time.Since(start))
		pretty.Println("module:", module.Text())
		for _, entity := range module.TopLevelEntity() {
			pretty.Println("entity:", entity.Text())
		}
		fmt.Println()
	}
}
