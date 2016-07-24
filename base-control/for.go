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

func main(){
    demo1()
    demo2()
}
