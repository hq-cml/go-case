package main

import (
    "crypto/md5"
    "crypto/sha1"
    "fmt"
)

func main() {
    TestString := "hello world!"

    //md5
    Md5Inst := md5.New()
    Md5Inst.Write([]byte(TestString))
    Result := Md5Inst.Sum([]byte(""))
    fmt.Printf("%x\n\n", Result)

    //sha-1
    Sha1Inst := sha1.New()
    Sha1Inst.Write([]byte(TestString))
    Result = Sha1Inst.Sum([]byte(""))
    fmt.Printf("%x\n\n", Result)
}

