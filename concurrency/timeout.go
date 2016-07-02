package main
//利用select+time+专用channel的方式，解决普通channel的超时问题
import (
    "time"
    "fmt"
)

func main(){
    timeout := make(chan bool, 1)
    ch := make(chan bool, 1)
    go func(){
        time.Sleep(1e9)
        timeout <- true
    }()

    select {
    case <- ch:
        fmt.Println("waiting")
    case <- timeout:
        fmt.Println("time out")
        break
    }
}
