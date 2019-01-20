package main

import (
	"fmt"
	"flag"
	"runtime/pprof"
	"runtime"
	"time"
)

var typ *string = flag.String("type", "cpu", "Test type: cpu/mem/block/lookup")
var sub_typ *string = flag.String("sub", "goroutine", "Sub type: goroutine/heap/allocs/threadcreate/block/mutex")

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
	} else if *typ == "lookup" {
		lookup("profile", sub_typ)
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

	//平均每分配多少个字节，就对堆内存的使用情况进行一次采样。
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

//lookup函数，会调用pprof.Lookup函数的6种可支持的参数，生成对应文件
func lookup(dir string, sub *string) {
	//平均每分配8个字节，就对堆内存的使用情况进行一次采样
	//一个阻塞事件的持续时间达到了2个纳秒，就对其进行采样
	runtime.MemProfileRate = 8
	runtime.SetBlockProfileRate(2)

	if sub == nil {
		fmt.Println("Sub type is nil")
		for name, _ := range lookupOps {
			for _, debug := range debugOpts {
				err := doLookup(dir, name, debug)
				if err != nil {
					return
				}
				time.Sleep(time.Millisecond)
			}
		}
	} else {
		fmt.Println("Sub type is ", *sub)
		for _, debug := range debugOpts {
			err := doLookup(dir, *sub, debug)
			if err != nil {
				return
			}
			time.Sleep(time.Millisecond)
		}
	}

}

func doLookup(dir, name string, debug int) error {
	fmt.Printf("Generate %s profile (debug: %d) ...\n", name, debug)
	fileName := fmt.Sprintf("%s_%d.out", name, debug)
	f, err := CreateFile(dir, fileName)
	if err != nil {
		fmt.Printf("create error: %v (%s)\n", err, fileName)
		return err
	}
	defer f.Close()
	if err = Execute(lookupOps[name], 10); err != nil {
		fmt.Printf("execute error: %v (%s)\n", err, fileName)
		return err
	}

	err = pprof.Lookup(name).WriteTo(f, debug)
	if err != nil {
		fmt.Printf("write error: %v (%s)\n", err, fileName)
		return err
	}
	return nil
}

// Lookup一共支持6种参数， 每种参数给一个负载函数
var lookupOps = map[string]OpFunc {
	"goroutine": 	BlockProfile,
	"heap": 		MemProfile,
	"allocs": 		MemProfile,
	"threadcreate": BlockProfile,
	"block": 		BlockProfile,
	"mutex": 		BlockProfile,
}

// debugOpts 代表debug参数的可选值列表。
var debugOpts = []int {
	0,
	1,
	2,
}