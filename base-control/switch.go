package main

import(
    "fmt"
)

//最普通用法
func checkLan(lan string){
    switch lan {
    case "Python", "PHP": //可以写多个选项~
        fmt.Println(lan, "is a interpreted language")
        //fallthrough     //放开注释的话，fallthrough的功能是类似C的switch
    case "Golang", "C", "Java":
        fmt.Println(lan, "is a compiled language")
    default:
        fmt.Println(lan, "is a unknown language")
    }
}


func main(){
    checkLan("Golang")
    checkLan("PHP")
    checkLan("go")

}