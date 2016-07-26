package main

import (
    "fmt"
    "errors"
    //"os"
)

//标准的error生成方法
func err1(){
    var err error = errors.New("A normal err")
    fmt.Println(err)
}

func main(){
    err1()
}
