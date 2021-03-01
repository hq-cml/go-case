package main

import (
    "fmt"
    "runtime"
)

//标准的panic处理几个原则：
// 1. 尽早捕获
// 2. 应该是分级处理，处理不了，再继续上报或者返回错误

type MyErrorIntfs interface {
    error                  //接口嵌入
    ErrorPrint()
}

//自定义类型实现接口
type MyError struct {
    msg string
}

func (e MyError)Error() string{
    return e.msg
}

func (e MyError)ErrorPrint(){
    fmt.Println(e.msg)
}

func throwPanic(s int){
    //尽早捕获，如果没法处理，再上报
    defer func(){
        if r:= recover(); r!=nil{
            if e, ok := r.(MyErrorIntfs); ok {
                fmt.Println("Recover(MyErrorIntfs).", e)
                e.ErrorPrint()
            } else {
                fmt.Println("Can't handle. Throw")
                panic(r)  //上报~
            }
        }
    }()

    if s == 1{
        var p = MyError{"test"}
        panic(p)                     //引发自定义错误
    }else{
        var arr = [3]int{1,2,3}
        idx := 3
        arr[idx] = 10                  //运行时错误
    }

}

func main(){
    //尽早捕获
    defer func(){
        if r:= recover(); r!=nil{
            if e, ok := r.(runtime.Error); ok {
                //调用栈
                buf := make([]byte, 1<<16)
                len := runtime.Stack(buf, true)
                fmt.Println("Recover(Runtime).", e)
                fmt.Println("Call stack: ", string(buf[0: len]))
            } else {
                fmt.Println("Recover(Unknown).", e)  //兜底
            }
        }
    }()

    fmt.Println("Main Begin!")

    throwPanic(1)
    fmt.Println("Main Middle!")
    throwPanic(0)

    fmt.Println("Main End!")  //无法运行到此处
}


