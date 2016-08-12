package main

import (
    "os/exec"
    "fmt"
    "bufio"
    "io"
    "bytes"
)

//最简单实例：执行echo -n "go cmd"
func pipeDemo1() {
    //创建一个exec.Cmd类型变量
    cmd0 := exec.Command("echo", "-n", "Go cmd")

    //创建管道用于捕获命令输出
    pipeOut0, err := cmd0.StdoutPipe()
    if err != nil{
        fmt.Println("Error: Create Pipe failed. " + err.Error())
        return
    }

    //启动执行cmd0
    if err := cmd0.Start(); err != nil{
        fmt.Println("Error: cmd0 failed. " + err.Error())
        return
    }

    //法一：从管道获取命令输出
    //output0 := make([]byte, 30)
    //n, err := stdout0.Read(output0)
    //if err != nil{
    //    fmt.Println("Error: Read from Pipe failed. " + err.Error())
    //}
    //str := string(output0[0:n])

    //法二：创建一个带缓冲的reader，获取命令输出
    outputBuf0 := bufio.NewReader(pipeOut0)
    for {
        output0, isPrefix, err := outputBuf0.ReadLine()

        if err != nil {
            if err == io.EOF {
                break // 结束
            }else{
                fmt.Println("Read somthing wrong!")
                return
            }

        }

        //当前行的长度超出缓冲区长度
        if isPrefix {
            fmt.Println("A too long line, seems unexpected.")
        }else{
            fmt.Println(string(output0))
        }
    }
}

//执行命令类似： ps aux | grep http
func pipeDemo2(){
    //创建cmd变量
    cmd1 := exec.Command("ps", "aux")
    cmd2 := exec.Command("grep", "http")

    //创建输出管道用于捕获命令输出
    pipeOut, err := cmd1.StdoutPipe()
    if err != nil{
        fmt.Println("Error: Create Pipe failed. " + err.Error())
    }

    //启动执行cmd1
    if err := cmd1.Start(); err != nil{
        fmt.Println("Error: cmd1 failed. " + err.Error())
    }

    //为cmd2创建输入管道，并将cmd1的输出管道的内容写入前者
    outputBuf1 := bufio.NewReader(pipeOut)

    pipeIn, err := cmd2.StdinPipe()

    if err != nil {
        fmt.Println("Error in create pipe in: "+err.Error())
    }
    outputBuf1.WriteTo(pipeIn)
    //err = pipeOut.Close()   //经过测试，这个close可有可无
    //创建cmd2的输出缓冲区，这样当cmd2执行完毕之后，其输出将保存在此缓冲
    var outputBuf2 bytes.Buffer
    cmd2.Stdout = &outputBuf2

    //执行cmd2,并关闭管道
    if err := cmd2.Start(); err != nil{
        fmt.Println("Error: cmd2 failed. " + err.Error())
        return
    }
    err = pipeIn.Close() //这个关闭时必须的，否则下面的Wait会一直等待
    if err != nil{
        fmt.Println("pipeIn close failed. " + err.Error())
        return
    }

    //等待cmd2执行完毕，如果不等待，主线程会提前结束
    if err := cmd2.Wait(); err != nil {
        fmt.Println("cmd wait failed. " + err.Error())
        return
    }

    //输出结果
    fmt.Println(outputBuf2.String())
}

func main() {
    pipeDemo1()
    pipeDemo2()
}
