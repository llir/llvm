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
		fmt.Printf("=== [ %v ] =======================\n", llPath)
		fmt.Println()
		start := time.Now()
		module, err := asm.ParseFile(llPath)
		if err != nil {
			log.Fatalf("%q: %+v", llPath, err)
		}
		fmt.Println("took:", time.Since(start))
		fmt.Println()
		m := asm.Translate(module)
		pretty.Println(m)
	}
}
