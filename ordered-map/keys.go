package ordered_map

import(
    "fmt"
    "sort"
    "reflect"
)

//函数类型声明
type CompareFunction func(interface{}, interface{}) int8

/*
 * 定义接口一枚，
 *
 */
type Keys interface {
    sort.Interface             //嵌入sort.Interface接口，意味着Keys的实现类型是可以排序的~
    Add(k interface{}) bool    //增加一个key
    Remove(k interface{}) bool //去除一个key
    Clear()                    //清空
    Get(index int) interface{} //按照index获取一个key
    GetAll() []interface{}     //全部获取
    Search(k interface{}) (index int, contains boo)  //查找key
    ElemType() reflect.Type    //获取key运行时的类型
    CompareFunc() CompareFunction   //返回一个函数类型
}