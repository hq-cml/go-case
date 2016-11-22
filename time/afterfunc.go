package main

// AfterFunc的注释：
// AfterFunc waits for the duration to elapse and then calls func in its own goroutine.
// 也就是说AfterFunc如果想要放在完美实现准时异步通知，需要放在独立的goroutine中
import (
    "fmt"
    "time"
)

func main(){
    c := make(chan byte)

    time.AfterFunc(time.Second*2, func(){
        fmt.Println("Time's up")
        c <- 1
    })

    //放开注释，查看效果对比
    //time.Sleep(time.Second*5)
    <-c
    fmt.Println("HAHA")

}

