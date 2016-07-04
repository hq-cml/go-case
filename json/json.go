package main

import "fmt"
import "encoding/json"

type Book struct{
    Title string
    Authors []string
    IsPublished bool
    Price float64
    TestTag int       `json:"haha"`
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

//知道json格式的情况下
func json_decode1(){
    b := []byte(`{"Title":"Go 学习","Authors":["hanxiao","hq"],"IsPublished":true,"Price":9.9,"haha":12345}`)
    var book Book
    err := json.Unmarshal(b, &book)
    if err != nil{
        fmt.Println("Json解码出错:", err.Error())
    }else{
        fmt.Printf("%v",book)
    }
}
func main(){
    json_encode()
    json_decode1()
}