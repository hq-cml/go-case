package main

import (
    "fmt"
    "io"
    "reflect"
)

type T struct{
    A int
    B string
}

type MyReader struct{
    Name string
}

func (r MyReader)Read(p []byte)(n int, err error){
    fmt.Println("hello")
    return
}

func main(){
    var reader io.Reader
    reader = &MyReader{"a"} //是否加&都能编译通过
    fmt.Println(reader)

    var x float64 = 3.4
    tof := reflect.TypeOf(x)
    vof := reflect.ValueOf(x)
    fmt.Println("type:", tof)
    fmt.Println("type:", vof.Type())
    fmt.Println("kind is float64? :", vof.Kind() == reflect.Float64)
    fmt.Println("value:", vof.Float())
    fmt.Println("can set:", vof.CanSet())

    //利用指针的反射，修改值本身
    pvof := reflect.ValueOf(&x)
    fmt.Println("type:", pvof.Type())
    fmt.Println("can set:", pvof.CanSet())
    v := pvof.Elem()
    fmt.Println("Elem can set:", v.CanSet())
    v.SetFloat(7.1)
    fmt.Println(v.Interface())
    fmt.Println("x is :", x)

    //利用反射获取一个结构中的所有成员的值
    t := T{203, "haha"}
    s := reflect.ValueOf(&t).Elem()
    typeOFT := s.Type()
    for i:= 0; i<s.NumField(); i++{
        f := s.Field(i)
        fmt.Printf("%d:%s %s = %v\n", i, typeOFT.Field(i).Name, f.Type(), f.Interface())
    }
}
