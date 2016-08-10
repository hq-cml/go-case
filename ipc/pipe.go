package main

import (
    "os/exec"
    "fmt"
)

func main() {
    //创建一个exec.Cmd类型变量
    cmd0 := exec.Command("echo", "-n", "Go cmd")

    //创建管道捕获命令输出
    stdout0, err := cmd0.StdoutPipe()
    if err != nil{
        fmt.Println("Error: Create Pipe failed. " + err.Error())
    }

    //启动执行cmd0
    if err := cmd0.Start(); err != nil{
        fmt.Println("Error: cmd0 failed. " + err.Error())
    }

    //从管道获取命令输出
    output0 := make([]byte, 30)
    n, err := stdout0.Read(output0)
    if err != nil{
        fmt.Println("Error: Read from Pipe failed. " + err.Error())
    }
    str := string(output0[0:n])
    fmt.Println(str)
}
