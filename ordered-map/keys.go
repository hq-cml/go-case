package ordered_map

import(
    "fmt"
    "sort"
    "reflect"
)

//函数类型声明
type CompareFunction func(interface{}, interface{}) int8

/*
 * 定义接口一枚，作为ordered-map的keys成员的变量
 * 需要实现这个接口的方法
 */
type Keys interface {
    sort.Interface                 //嵌入sort.Interface接口，意味着Keys的实现类型是可以排序的~
    Add(k interface{}) bool        //增加一个key
    Remove(k interface{}) bool     //去除一个key
    Clear()                        //清空
    Get(index int) interface{}     //按照index获取一个key
    GetAll() []interface{}         //全部获取，存于一个切片中
    ElemType() reflect.Type        //获取key运行时的类型
    CompareFunc() CompareFunction  //获取运行时用于比较key大小的具体方法
    Search(k interface{}) (index int, contains boo)  //查找key
}

//自定义类型myKeys
type myKeys struct {
    container    []interface{}    //keys的实际容器，keys元素可以是任意类型
    // compareFunc的结果值：
    //   小于0: 第一个参数小于第二个参数
    //   等于0: 第一个参数等于第二个参数
    //   大于0: 第一个参数大于第二个参数
    compareFunc  CompareFunction  //函数也是一种类型，compareFunc负责比较元素的大小，具体实现交给上层开发者
    elemType     reflect.Type     //存储keys元素的实际类型，（运行时确定）
}

//让类型*myKeys实现Keys接口:
//首先实现嵌入接口sort.Interface：Len，Less，Swap
func (keys *myKeys) Len() int{
    return len(keys.container)
}
func (keys *myKeys) Less(i, j int) bool{
    return keys.compareFunc(keys.container[i], keys.container[j]) < 0
}
func (keys *myKeys) Swap(i, j int){
    keys.container[i], keys.container[j] = keys.container[j], keys.container[i]
}
