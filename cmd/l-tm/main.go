package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/llir/llvm/asm"
)

func main() {
	var (
		cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
		memprofile = flag.String("memprofile", "", "write mem profile to file")
	)
	flag.Parse()

	if *cpuprofile != "" {
		fd, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		pprof.StartCPUProfile(fd)
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		runtime.MemProfileRate = 1
	}

	for _, llPath := range flag.Args() {
		fmt.Printf("=== [ %v ] =======================\n", llPath)
		fmt.Println()
		fileStart := time.Now()
		m, err := asm.ParseFile(llPath)
		if err != nil {
			log.Fatalf("%q: %+v", llPath, err)
		}
		fmt.Printf("total time for file %q: %v\n", llPath, time.Since(fileStart))
		_ = m
		//pretty.Println(m)
	}

	if *memprofile != "" {
		fd, err := os.Create(*memprofile)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		runtime.GC()
		err = pprof.WriteHeapProfile(fd)
		if err != nil {
			log.Fatalf("WriteHeapProfile: %v", err)
		}
	}
}
