package main

/*
 * 反射三原则：
 * 1.Reflection goes from interface value to reflecton object.
 * 2.Reflection goes from reflection object to interface value.
 * 3.To modify a reflection object, the value must be settable.
 */
import (
    "reflect"
    "fmt"
    "runtime"
)

type T struct {
    A int
    B string
}

type MyFloat float64

func main() {
    fmt.Println(runtime.Version())

    //反射第一定律:从接口值到反射对象
    var x1 float64 = 3.4
    var x2 MyFloat = 3.4

    fmt.Println("TypeOf:", reflect.TypeOf(x1), reflect.TypeOf(x2))
    fmt.Println("Valueof:", reflect.ValueOf(x1), reflect.ValueOf(x2))

    v1 := reflect.ValueOf(x1)
    v2 := reflect.ValueOf(x2)
    fmt.Println("Type:", v1.Type(), v2.Type())
    fmt.Println("Kind:", v1.Kind(), v2.Kind())
    fmt.Println("Float:", v1.Float(), v2.Float())
    //fmt.Println("Int:", v.Int())    //runtime错误,

    //反射第二定律:从反射对象到接口值
    i := v1.Interface().(float64)
    fmt.Println("interface:", i)
    fmt.Println("interface:", i)

    //反射第三定律:为了修改一个反射对象，值必须是settable的
    //v1.SetFloat(7.1) //panic, Settability是false
    fmt.Println("settability of v:", v1.CanSet())

    p := reflect.ValueOf(&x1)
    fmt.Println("type of P:", p.Type())
    fmt.Println("Settablility of p:", p.CanSet())
    //修改需要用的Elem方法
    v := p.Elem()
    fmt.Println("Settablility of v:", v.CanSet())

    v.SetFloat(7.1)
    fmt.Println(v.Interface())
    fmt.Println(v)

    //利用反射遍历一个结构体
    t := T{23, "skidoo"}
    s := reflect.ValueOf(&t).Elem()
    typeOfT := s.Type()
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%d: %s(%s) %s = %v\n",
            i, typeOfT.Field(i).Name, typeOfT.Field(i).Type, f.Type(), f.Interface())
    }
    //修改结构体
    s.Field(0).SetInt(77)
    s.Field(1).SetString("Sunset Strip")
    fmt.Println("t is now", t)
}