package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mewmew/l-tm/asm"
	"github.com/mewmew/l-tm/asm/ll/ast"
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
		fmt.Println("module:", module.Text())
		for _, entity := range module.TopLevelEntities() {
			fmt.Printf("entity %T: %v\n", entity, entity.Text())
			switch entity := entity.(type) {
			case *ast.SourceFilename:
				fmt.Println("   name:", entity.Name().Text())
			case *ast.TargetDataLayout:
				fmt.Println("   datalayout:", entity.DataLayout().Text())
			case *ast.TargetTriple:
				fmt.Println("   target triple:", entity.TargetTriple().Text())
			case *ast.ModuleAsm:
				fmt.Println("   module asm:", entity.Asm().Text())
			}
		}
		fmt.Println()
	}
}
