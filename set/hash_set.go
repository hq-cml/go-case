package main

import "fmt"

/*
 * 利用内置的hash实现set功能，用key存储set的值
 * key用interface{}类型，表示set的值是任意类型
 * val用bool类型，因为1，省空间 2，用起来方便，可以表示是否存在等
 */
type HashSet struct{
    m map[interface{}]bool
}

//“构造”函数
func NewHashSet() *HashSet{
    return &HashSet{m: make(map[interface{}]bool)}
}

//添加元素
func (set *HashSet) Add(e interface{}) bool{
    if _, ok := set.m[e]; !ok{
        set.m[e] = true
        return true
    }
    return false
}

//删除元素
func (set *HashSet) Remove(e interface{}) {
    delete(set.m, e)
}

//清空set，有两个点说明~
//1. 直接重新赋值set.m，旧的m交由垃圾回收器去回收
//2. 注意接收者是*hashSet，否则无法达到
func (set *HashSet) Clear() {
    set.m = make(map[interface{}]bool)
}

//判断某个元素是否存在
func (set *HashSet) Contains(e interface{}) {
    if _, ok := set.m[e]; ok{
        return true
    } else {
        return false
    }
}


func main(){
    m := make(map[string]string)
    a := m["A"]

    if v,ok := m["A"]; ok{
        fmt.Println("Yes", v)
    } else {
        fmt.Println("No", v)
    }
    fmt.Println(a)
    m["A"] = "abc"
    a = m["A"]
    if v,ok := m["A"]; ok{
        fmt.Println("Yes", v)
    } else {
        fmt.Println("No", v)
    }
    fmt.Println(a)
}