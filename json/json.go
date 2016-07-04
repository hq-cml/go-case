package main

import "fmt"
import "encoding/json"

type Book struct{
    Title string
    Authors []string
    IsPublished bool
    Price float64
}

func json_encode(){
    book := Book{
        "Go 学习",
        []string{"hanxiao", "hq"},
        true,
        9.9,
    }
    b, err := json.Marshal(book)
    //fmt.Println(book)
    if err != nil{
        fmt.Println("Json编码出错")
    }else{
        fmt.Println(string(b))
    }
}
func main(){
    json_encode()
}