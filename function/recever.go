package main
/*
 * 这个例子说明一个问题：
 * 无论函数的接收者是否是一个指针，都存在两种调用方式，普通类型调用和指针调用。
 *
 * 同时，到底能否改变参数值，需要看接收者是否为指针，而不看调用者是否为指针
 */
import (
    "fmt"
)

type Integer int

func (a Integer)Add_1(b Integer){
    a += b
}

func (a *Integer)Add_2(b Integer){
    *a += b
}

func main(){
    var a Integer = 1
    c := &a

    //Add_1的接收者是Integer，但是却支持a和c两种调用方式
    //而且，即便c为指针，结果都改变不了a
    a.Add_1(10)
    fmt.Println(a)    //1
    c.Add_1(10)
    fmt.Println(a)    //1

    //Add_2的接收者是Integer指针，但是却支持a和c两种调用方式
    //而且，即便a不是指针，结果仍然会发生改变
    a.Add_2(10)
    fmt.Println(a)    //11
    c.Add_2(10)
    fmt.Println(a)    //21
}
