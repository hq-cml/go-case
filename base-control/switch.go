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

//swtich用于类型判断：
//注意，v必须是某种接口类型
func checkType(v interface{}){
    //s := "SSSS"
    //v := interface{}(s) //v必须是某种接口类型，否则编译报错
    switch v.(type) {
    case string: //可以写多个选项~
        fmt.Printf("The string is %s\n", v)
    case int, uint:
        fmt.Printf("The int is %d\n", v)
    default:
        fmt.Printf("Unsupported type.(type=%T)\n", v)
    }
}

func main(){
    checkLan("Golang")
    checkLan("PHP")
    checkLan("go")

    var v interface{} = "BBBB"
    checkType(v)
    checkType("AAA")
    checkType(13)
    checkType(13.0)
}