package main
/*
 * 闭包：通过捕获“自由变量”的绑定对函数文本执行的“闭合”动作。
 * 前提：函数可以作为一个函数的返回值返回
 * 自由变量：从外层透传进来的外部的变量
 *
 * 比如，例子中的x就是自由变量，它在内部那个匿名函数生成的时候，与getAdd函数的
 * 参数x进行了绑定，从而使得内部那个匿名函数变得完整起来。可以这么描述，通过捕
 * 获自由变量x的绑定，使得getAdd函数返回的那个函数“闭合”了。
 */
import (
    "fmt"
)

//getAdd函数实现了一个闭包：
//每调用一次getAdd，参数x都会被作为内部那个匿名引用着。
//所以只要这个返回的函数仍然还可以被调用，这个参数x就会一直存在，不会被Go的垃圾回收器收回。
func getAdd(x int) func(y int) int {
    //返回一个 “func(y int)int” 类型的匿名的函数
    return func(y int) int {
        return x + y            //x是自由变量，从外部透传进来。
    }
}

func main(){
    my_add := getAdd(10)
    fmt.Println(my_add(20))  //30
}
