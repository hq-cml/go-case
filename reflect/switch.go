package main
/*
 * 类型判断
 * 1. switch+type
 */

import "fmt"

type TEST struct {
    s string
}

type MYINT int
type MYINT_AAA int

//swtich用于类型判断：
//注意，v必须是某种接口类型
func checkType(v interface{}){
    switch v.(type) {
    case string: //可以写多个选项~
        fmt.Printf("The value %s's type is string\n", v)
    case int, uint:
        fmt.Printf("The value %d's type is int\n", v)
    case TEST:
        fmt.Printf("The value %+v's type is TEST\n", v)
    case MYINT:
        fmt.Printf("The value %v's type is MYINT\n", v)
    default:
        fmt.Printf("Unsupported type.(type=%T)\n", v)
    }
}

func main(){
    var v interface{} = "BBBB"
    checkType(v)
    checkType("AAA")
    checkType(13)
    checkType(13.0)

    v1 := TEST{

    }
    checkType(v1)

    var v2 MYINT = 10
    checkType(v2)

    var v3 MYINT_AAA = 10
    checkType(v3)
}