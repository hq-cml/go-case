package main

import (
    "fmt"
    "errors"
    "os"
)

//标准的error生成方法
func err1(){
    var err error = errors.New("A normal err")
    fmt.Println(err)
}

//标准的error处理
//通常的函数返回的都是error接口类型，所以
//处理的时候通过类型断言，逐层判断error类型
func handle_err(){
    _, err := os.Open("/tmp/aaa")
    if err != nil{
        if pathErr, ok := err.(*os.PathError); ok{
            //os.PathError类型错误
            fmt.Printf("Path Error:%s, (op:%s, path=%s)\n", pathErr.Err, pathErr.Op, pathErr.Path)
        } else {
            fmt.Println("Unknow Error :", err)
        }
    }
}
func main(){
    err1()
    handle_err()
}
