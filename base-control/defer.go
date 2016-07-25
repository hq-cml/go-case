package main

import "fmt"

//循环defer不推荐，会埋下坑
//参考下面是一些实例

//在定义的时候，传递给defer语句在参数会被保存
//4 3 2 1 0
func printNum1(){
    for i:=0; i<5; i++{
        defer fmt.Printf("%d ", i)
    }
}

//defer + 匿名函数
//注意，如果不是以参数形式传递给defer的匿名函数，则不会被保存
//5 5 5 5 5
func printNum2(){
    for i:=0; i<5; i++{
        defer func(){
            fmt.Printf("%d ", i) //非参数，不会保存
        }()
    }
}

//defer + 匿名函数
//传递给defer语句在参数会被保存
//4 3 2 1 0
func printNum3(){
    for i:=0; i<5; i++{
        defer func(j int){
            fmt.Printf("%d ", j)
        }(i)  //参数，会保存
    }
}

//defer是可以改变外部命名结果值的~
func modify(n int) (num int){
    defer func(){
        num += n
    }()
    num++
    fmt.Println("Num is ", num)
    return
}

func main() {
    printNum1()
    fmt.Println()
    printNum2()
    fmt.Println()
    printNum3()
    fmt.Println()
    num := modify(5)
    fmt.Println("Num is ", num)
}
