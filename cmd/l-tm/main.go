package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/llir/llvm/asm"
)

func main() {
	flag.Parse()
	for _, llPath := range flag.Args() {
		fmt.Printf("=== [ %v ] =======================\n", llPath)
		fmt.Println()
		fileStart := time.Now()
		m, err := asm.ParseFile(llPath)
		if err != nil {
			log.Fatalf("%q: %+v", llPath, err)
		}
		_ = m
		//pretty.Println(m)
		fmt.Printf("total time for file %q: %v\n", llPath, time.Since(fileStart))
	}
}
