package main
//for + select的惯用用法
//利用time+专用channel，解决普通channel的超时问题

import (
    "time"
    "fmt"
)

//超时发生函数：Channel版本
func genTimeoutChannel() chan bool {
    fmt.Println("Geneate Timer")
    timeout := make(chan bool, 1)
    //timeout := make(chan bool) //这个地方也可以使用不带缓冲版本，效果相同

    //匿名goroutine，在t秒之后，向timeout中产出一个bool值
    go func() {
        time.Sleep(5*time.Second)
        timeout <- false
    }()

    return timeout
}

//超时发生函数：官方定时器版本1
func genTimeoutTimer() chan bool {
    fmt.Println("Geneate Timer")
    timeout := make(chan bool, 1)

    go func() {
        timer := time.NewTimer(5*time.Second)
        <- timer.C
        timeout <- false
    }()

    return timeout
}

//超时发生函数：官方定时器版本2
func genTimeoutTimerFunc() chan bool {
    fmt.Println("Geneate Timer")
    timeout := make(chan bool, 1)

    f := func() {
        timeout <- false
    }

    //异步的方式自动执行，和上个版本启动goroutinue效果相同
    time.AfterFunc(5*time.Second, f)

    return timeout
}

//超时发生函数
func genTimeout(t int) chan bool {
    if 1 == t {
        return genTimeoutTimer()
    }else if 2 == t {
        return genTimeoutTimerFunc()
    } else {
        return genTimeoutChannel()
    }
}

func main(){
    ch := make(chan int, 10)
    //e 和 ok都需要实现初始化完毕，因为：
    //下面的case中不能使用:=，只能使用=，使用了前者就会变成小作用域临时变量
    //哪怕当中有一个同名的变量也是如此，两个变量都会成为临时。那么下面的
    //跳出for的break将永远失去执行机会
    var e int
    ok := true

    go func() {
       for i:=0; i<5; i++ {
           time.Sleep(2*time.Second)
           ch <- i
       }
       //放开注释就无法观察超时效果了
       //close(ch)
    }()

    //通常select都会配套一个for，否则就成了一次性
    for {
        select {
        case e, ok = <-ch:
            if !ok {
                fmt.Println("Channel is closed. End")
                break //跳出select
            } else {
                fmt.Println("Got", e)
            }
        case ok = <-genTimeout(2): //每次select的超时时间是相同的5s
            fmt.Println("Time out. End")
            break //跳出select
        }

        if !ok {
            break //跳出 for
        }

    }
}
