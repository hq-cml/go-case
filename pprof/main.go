package main

import (
	"fmt"
	"flag"
	"runtime/pprof"
)

var typ *string = flag.String("type", "cpu", "Test type: cpu/mem/block")

func main() {
	flag.Parse()
	if typ == nil {
		fmt.Println("Error argument")
	}

	if *typ == "cpu" {
		cpu("profile", "cpu.profile")
	} else if *typ == "mem" {
		mem("profile", "mem.profile")
	}
}

func cpu(dir, file string) {
	f, err := CreateFile(dir, file)
	if err != nil {
		fmt.Printf("%s/%s profile creation error: %v\n", dir, file, err)
		return
	}
	defer f.Close()

	//启动CPU分析
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Printf("%s/%s profile start error: %v\n", dir, file, err)
		return
	}
	defer pprof.StopCPUProfile()

	//模拟CPU起飞
	if err = Execute(CPUProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
}

func mem(dir, file string) {
	f, err := CreateFile(dir, file)
	if err != nil {
		fmt.Printf("%s/%s profile creation error: %v\n", dir, file, err)
		return
	}
	defer f.Close()

	//可以设定，也可以不设定，默认512KB
	//runtime.MemProfileRate = 256

	//模拟Mem起飞
	if err = Execute(MemProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
}
