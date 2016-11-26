package main

/*
 * 这个例子说明了：
 * goroutine没有父子关系不明显，貌似更像是主goroutine和从goroutine的关系。
 * 只要主goroutine不结束，那么所有的goroutine都能正常结束，哪怕goroutine的父亲结束了，
 * 只要主goroutine还在，子goroutine也不会受到影响
 */

import (
    "fmt"
    "time"
    //"sync"
)

func main(){
    //wg := sync.WaitGroup{}
    //wg.Add(3)
    go func() {
        go func() {
            fmt.Println("sunzi 1 start")
            time.Sleep(1 * time.Second)
            fmt.Println("sunzi 1 stop")
            //wg.Done()
        }()

        go func() {
            fmt.Println("sunzi 1 start")
            time.Sleep(1 * time.Second)
            fmt.Println("sunzi 2 stop")
            //wg.Done()
        }()
        fmt.Println("son stop")
        //wg.Done()
    }()
    //wg.Wait()
    time.Sleep(2*time.Second)
}
