package main

import (
    "fmt"
    "sort"
)

//在sort包中，接口Interface定义如下，有三个方法
//type Interface interface {
//    Len() int
//    Less(i, j int) bool
//    Swap(i, j int)
//}

//接口嵌套，Sortable已经拥有了sort.Interface的方法
type Sortable interface{
    sort.Interface
    Sort()
}

//自定义类型
type SortableArr [3]string

//func1
func (self SortableArr) Len() int{
    return len(self)
}

//func2
func (self SortableArr) Less(i,j int) bool{
    return self[i] < self[j]
}

//func3
func (self SortableArr) Swap(i,j int) {
    self[i], self[j] = self[j], self[i]
}

//func4
func (self SortableArr) Sort() {
    sort.Sort(self)
}

func main(){
    //func1,2,3的实现，表明SortableArr实现了sort.Interface
    _,ok := interface{}(SortableArr{}).(sort.Interface)
    fmt.Println(ok)  //true

    //可以看到，指针版本也实现了sort.Interface，这说明，在接口判定的时候，非指针版本可以推导出指针版本
    _,ok = interface{}(&SortableArr{}).(sort.Interface)
    fmt.Println(ok) //true

    //如果没有func4的定义，则此处是false，这很好理解
    //将接func4收者改成指针，则此处是false，这说明，在接口判定的时候，指针版本不能推导出非指针版本
    _,ok = interface{}(SortableArr{}).(Sortable)
    fmt.Println(ok) //true

    //正常情况打印，是2,3,1，很好理解，因为func4的接收者是普通值
    //如果要改变值，则需要将func1,2,3,4全改成指针版本接收者即可，但此时，上面只有第二个是true了，其他两个是false
    ss := SortableArr{"2", "3", "1"}
    ss.Sort()
    fmt.Println(ss)
}