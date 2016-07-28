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

//获取Set长度
func (set *HashSet) Len() {
    return len(set.m)
}

//判断HashSet是否相等（拥有相同的元素集合）
func (set *HashSet) Same(other *HashSet) bool {
    if other == nil {
        return false
    }

    if set.Len() != other.Len() {
        return false
    }

    for key := range set.m{ //range hash，只有一个接收值，得到的是key
        if !other.Contains(key){
            return false
        }
    }
    return true
}

//生成Hashset的一个快照，用slice存储，用于顺序遍历等场景
//Hashset.m并非线程安全，所以在快照生成过程中可能发生变化
//所以设计了actualLen，确保slice是完满的
func (set *HashSet) Elements() []interface{}{
    initialLen := len(set.m)
    snapshot := make([]interface{}, initialLen)
    actualLen := 0
    for key := range set.m{
        if actualLen >= initialLen{
            snapshot = append(snapshot, key)
        }else{
            snapshot[actualLen] = key
        }
        actualLen++
    }

    if actualLen < initialLen{
        snapshot = snapshot[:actualLen]  //二次切片，去除之前多申请的一部分
    }

    return snapshot
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