package main

/*
 * 关于golang中定时器的用法，大致有三种，1. sleep+channel 2. NewTimer 3.AfterFunc
 * 三种用法的对比：
 * 1. sleep+channel，理解和使用都很简单，缺点是要启动一个独立的goroutine，因为要防止阻塞住主goroutine
 * 2. AfterFunc是最完美的，即便协程阻塞了，到期后函数的触发也可以完成，并且这个函数本身也是在一个goroutine中执行的（见源码及其注释）
 * 3. NewTimer到期后的触发，也是异步的，但是由于NewTimer返回的channel要被等待，所以还是要启动一个独立goroutine
 *
 * 综上，尽量使用AfterFunc。不推荐NewTimer，容易挖坑。
 */

import (
    "fmt"
    "time"
)

func testChan() {
    fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Begin")
    c := make(chan byte)

    go func() {
        time.Sleep(2*time.Second)
        fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Time's up")
        c <- 1
    }()

    go func() {
        <-c
        fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Got it")
    }()

    //主goroutine阻塞
    fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Main goroutine block!")
    time.Sleep(time.Second*10)

    fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "End")
}

//可以看出，即便主goroutine阻塞，AfterFunc的到期也不受影响。完全异步的。并且更难得的是到期触发的那个匿名函数，也是放在另一个独立goroutine中的
//这可以从golang源码及注释中看到。
func testAfterFunc() {
    fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Begin")
    c := make(chan byte)

    time.AfterFunc(time.Second*2, func(){
        fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Time's up")
        c <- 1
    })

    go func() {
        <-c
        fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Got it")
    }()

    //主goroutine阻塞
    fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Main goroutine block!")
    time.Sleep(time.Second*10)

    fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "End")
}

//可以看出，即便主goroutine阻塞，到期也不受影响。完全异步的。但是由于NewTimer的channel要被等待，所以还是要启动一个独立goroutine
func testNewTimer() {
    fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Begin")
    c := make(chan byte)

    timer := time.NewTimer(2*time.Second)
    go func() {
        <- timer.C
        fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")),"Time's up")
        c <- 1
    }()

    go func() {
        <-c
        fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Got it")
    }()

    //如果不搭配goroutine使用，则要非常小心，这段代码如果放在上面，就会造成死锁，而且这段代码效果多少会影响到程序本意。会使得主goroutine多出两秒
    //timer := time.NewTimer(2*time.Second)
    //<-timer.C
    //fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Time's up")
    //c <- 1

    fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "Main goroutine block!")
    time.Sleep(time.Second*10)
    fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")), "HAHA")
}

func main(){
    //testAfterFunc()
    //testChan()
    testNewTimer()
}

