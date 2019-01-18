package main

import (
	"fmt"
	"flag"
	"runtime/pprof"
	"runtime"
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
	} else if *typ == "block" {
		block("profile", "block.profile")
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

	//最终报告写入文件
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

	//runtime.MemProfileRate = 256 //可以设定，也可以不设定，默认512KB

	//最终报告写入文件
	defer pprof.WriteHeapProfile(f)

	//模拟Mem起飞
	if err = Execute(MemProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
}

func block(dir, file string) {
	f, err := CreateFile(dir, file)
	if err != nil {
		fmt.Printf("%s/%s profile creation error: %v\n", dir, file, err)
		return
	}
	defer f.Close()

	//必须设定，因为源码中没有默认值
	//只要发现一个阻塞事件的持续时间达到了2个纳秒，就可以对其进行采样
	runtime.SetBlockProfileRate(2)

	//最终报告写入文件
	defer pprof.Lookup("block").WriteTo(f, 2)

	//模拟Block
	if err = Execute(BlockProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
}