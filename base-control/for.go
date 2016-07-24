package main

import(
    "fmt"
)

//最基本用法，和C类似
func demo1(){
    for i:=0; i<3; i++ {
        fmt.Printf("%d  ", i)
    }
    fmt.Println()
}

//golang风格
func demo2(){
    var num int
    for num < 3 {
        num ++
        fmt.Printf("%d  ", num)
    }
    fmt.Println()
}

//golang风格死循环
func demo3(){
    var num int
    for {
        num ++
        fmt.Printf("%d  ", num)
        if num > 3{
            break
        }
    }
    fmt.Println()
}

//go + range
func demo4(){
    arr := []int{1,2,3,4}
    for k,v := range arr{
        fmt.Printf("%d => %d ,", k, v)
    }
    fmt.Println()
    //注意，如果只有一个接收值，那么这个值是key，而不是value！！
    //如果只想要value，那么应该 for _,v := range arr{ ...
    for v := range arr{
        fmt.Printf("%d ,", v)
    }
    fmt.Println()
}

//go break，普通的break只能跳出一层，但是呆了标记的break，可以跳出指定层数字
func demo5(){
    arr1 := []int{1,2,3,4}
    arr2 := []int{1,2,3,4}

    L1: //直接跳到外层的L之后，注意，是不再执行而非重新执行
    for _,v1 := range arr1{
        //L2:
        for _,v2 := range arr2{
            if(v1 == 2){
                //尝试放开注释，看看结果~
                //break
                break L1
                //break L2
            }
            fmt.Println(v1, v2)
        }
    }
}

//go continue
func demo6(){
    arr1 := []int{1,2,3,4}
    arr2 := []int{1,2,3,4}

    L1: //直接跳到外层的L之后，注意，是继续执行而非重新执行
    for _,v1 := range arr1{
        //L2:
        for _,v2 := range arr2{
            fmt.Printf("Num:")
            if(v1 == 2){
                //尝试放开注释，看看结果~
                fmt.Println()
                //continue
                continue L1
                //continue L2
            }
            fmt.Printf("%d, %d", v1, v2)
            fmt.Println()
        }
    }
}

//go goto
func demo7(){
    arr1 := []int{1,2,3,4}
    arr2 := []int{1,2,3,4}

    L1: //直接跳到外层的L之后，注意，重新执行，这会造成死循环 ！！
    for _,v1 := range arr1{
        //L2:
        for _,v2 := range arr2{
            fmt.Printf("Num:")
            if(v1 == 2){
                //尝试放开注释，看看结果~
                fmt.Println()
                //continue
                goto L1
                //continue L2
            }
            fmt.Printf("%d, %d", v1, v2)
            fmt.Println()
        }
    }
}

func main(){
    demo1()
    demo2()
    demo3()
    demo4()

    //体会break、continue、goto的不同
    demo5()
    demo6()
    //demo7() //死循环


}
