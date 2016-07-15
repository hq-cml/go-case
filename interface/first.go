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

    speaker = SimpleSpeaker{"World"}
    //speaker.Speak()
}