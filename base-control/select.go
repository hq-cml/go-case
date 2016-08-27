package main
//for + select的惯用用法
//利用time+专用channel，解决普通channel的超时问题

import (
    "time"
    "fmt"
)

//超时发生函数
func genTimeout() chan bool {
    timeout := make(chan bool, 1)

    //匿名goroutine，在t秒之后，向timeout中产出一个bool值
    go func() {
        time.Sleep(5*time.Second)
        timeout <- false
    }()

    return timeout
}

func main(){
    ch := make(chan int, 10)
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
        case ok = <-genTimeout(): //每次select的超时时间是相同的5s
            fmt.Println("Time out. End")
            break //跳出select
        }

        if !ok {
            break //跳出 for
        }

    }
}
