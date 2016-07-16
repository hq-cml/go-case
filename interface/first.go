package main

import "fmt"

type ISpeaker interface{
    Speak()
}

type SimpleSpeaker struct{
    Message string
}

func (speaker *SimpleSpeaker) Speak(){
    fmt.Println(speaker.Message)
}

func main(){
    var speaker ISpeaker
    speaker = &SimpleSpeaker{"Hello"}
    speaker.Speak()

    //如下写法会报错，因为SimpleSpeaker并未实现ISpeaker，而是*SimpleSpeaker实现了
    //speaker = SimpleSpeaker{"World"}
    //speaker.Speak()
}