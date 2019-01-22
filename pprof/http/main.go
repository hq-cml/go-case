package main

import (
	"net/http"
	_ "net/http/pprof"
	"log"
	"flag"
	"github.com/hq-cml/go-case/pprof/common"
	"fmt"
	"time"
)

var typ *string = flag.String("type", "cpu", "Test type: cpu/mem/block/lookup")

func main() {
	flag.Parse()
	if typ == nil {
		fmt.Println("Error argument")
	}

	go func() {
		log.Println(http.ListenAndServe(":8080", nil))
	}()

	for i:=20; i>0; i-- {
		fmt.Println("The", *typ, "test will begin in", i)
		time.Sleep(1*time.Second)
	}

	if *typ == "cpu" {
		fmt.Println("Cpu Test")
		//模拟CPU起飞
		if err := common.Execute(common.CPUProfile, 100); err != nil {
			fmt.Printf("execute error: %v\n", err)
			return
		}
	} else if *typ == "mem" {
		fmt.Println("Mem Test")
		//模拟Mem起飞
		if err := common.Execute(common.MemProfile, 100); err != nil {
			fmt.Printf("execute error: %v\n", err)
			return
		}
	} else if *typ == "block" {
		fmt.Println("Block Test")
		//模拟Block
		if err := common.Execute(common.BlockProfile, 100); err != nil {
			fmt.Printf("execute error: %v\n", err)
			return
		}
	} else {
		fmt.Println("Error")
	}

	time.Sleep(1000*time.Second)
}

