package main

import (
    "fmt"
    "errors"
    "os"
    "bufio"
    "bytes"
    "io"
)

//标准的error生成方法
func err1(){
    var err error = errors.New("A normal err")
    fmt.Println(err)
}

//标准的error处理
//1. 通常的函数返回的都是error接口类型，所以处理的时候通过类型断言，逐层判断error类型
//2. 通过与预定义错误变量进行比较
func handle_err(){
    file, err := os.Open("/tmp/aaa")
    defer file.Close()
    //逐层递进判断错误
    if err != nil{
        if pathErr, ok := err.(*os.PathError); ok{
            //os.PathError类型错误
            fmt.Printf("Path Error:%s, (op:%s, path=%s)\n", pathErr.Err, pathErr.Op, pathErr.Path)
        } else {
            fmt.Println("Unknow Error :", err)
        }
        return
    }

    r := bufio.NewReader(file)
    var buf bytes.Buffer
    for{
        byteArray, _, err := r.ReadLine()
        if err != nil {
            //直接和预定义的错误变量进行比较
            if err == io.EOF{
                break
            } else {
                fmt.Println("Read error:", err)
            }
        }else{
            buf.Write(byteArray)
        }
    }
    fmt.Println("The file:", string(buf.Bytes()))
}
func main(){
    err1()
    handle_err()
}
