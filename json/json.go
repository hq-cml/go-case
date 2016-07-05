package main

import "fmt"
import "encoding/json"

type Book struct{
    Title string
    Authors []string
    IsPublished bool
    Price float64
    TestTag int       `json:"haha"` //Tag, 在编码和解码中都会生效~
}

func json_encode(){
    book := Book{
        "Go 学习",
        []string{"hanxiao", "hq"},
        true,
        9.9,
        1,
    }
    b, err := json.Marshal(book)
    //fmt.Println(book)
    if err != nil{
        fmt.Println("Json编码出错")
    }else{
        fmt.Println(string(b))
    }
}

//json结构已知的情况下
func json_decode1(){
    b := []byte(`{"Title":"Go 学习","Authors":["hanxiao","hq"],"IsPublished":true,"Price":9.9,"haha":12345}`)
    var book Book
    err := json.Unmarshal(b, &book)
    if err != nil{
        fmt.Println("Json解码出错:", err.Error())
    }else{
        fmt.Printf("%v\n",book)
    }
}

//json结构未知的情况
func json_decode2(){
    b := []byte(`{"Str":"Go 学习","Arr":["hanxiao","hq"],"Bool":true,"Float":9.9}`)
    var r interface{} //结构未知，则解码接受变量用一个空接口变量
    err := json.Unmarshal(b, &r)
    if err != nil{
        fmt.Println("Json解码出错:", err.Error())
    }else{
        fmt.Printf("%v\n",r)
    }

    
}

func main(){
    json_encode()
    json_decode1()
    json_decode2()
}