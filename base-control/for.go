package main

import(
    "fmt"
)

//最基本用法，和C类似
func demo1(){
    for i:=0; i<3; i++ {
        fmt.Printf("%d  ", i)
    }
    fmt.Println()
}

//golang风格
func demo2(){
    var num int
    for num < 3 {
        num ++
        fmt.Printf("%d  ", num)
    }
    fmt.Println()
}

//golang风格死循环
func demo3(){
    var num int
    for {
        num ++
        fmt.Printf("%d  ", num)
        if num > 3{
            break
        }
    }
    fmt.Println()
}

//go + range
func demo4(){
    arr := []int{1,2,3,4}
    for k,v := range arr{
        fmt.Printf("%d => %d ,", k, v)
    }
    fmt.Println()
    //注意，如果只有一个接收值，那么这个值是key，而不是value！！
    for v := range arr{
        fmt.Printf("%d ,", v)
    }
    fmt.Println()
}

func main(){
    demo1()
    demo2()
    demo3()
    demo4()
}
