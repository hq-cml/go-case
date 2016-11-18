package ordered_map

import(
    "fmt"
    "sort"
    "reflect"
    "bytes"
)

//函数类型声明
//compareFunction的结果值：
//小于0: 第一个参数小于第二个参数
//等于0: 第一个参数等于第二个参数
//大于0: 第一个参数大于第二个参数
//第1和2个参数是key，第三个参数是map本身.用户自定制compare功能，选择比较key本身或者是比较的是map的值
type CompareFunction func(interface{}, interface{}, map[interface{}]interface{}) int8

/*
 * 定义接口，作为ordered-map的keys成员的变量
 */
type KeysIntfs interface {
    sort.Interface                 //嵌入sort.Interface接口，意味着Keys的实现类型是可以排序的~
    Add(k interface{}) bool        //增加一个key
    Remove(k interface{}) bool     //去除一个key
    Clear()                        //清空
    Get(index int) interface{}     //按照index获取一个key
    GetAll() []interface{}         //全部获取，存于一个切片中
    ElemType() reflect.Type        //获取key运行时的类型
    CompareFunc() CompareFunction  //获取运行时用于比较key大小的具体方法
    Search(k interface{}) (index int, contains bool)  //查找key
}

//自定义类型myKeys
type orderedKeys struct {
    container    []interface{}    //keys的实际容器，keys元素可以是任意类型
    compareFunc  CompareFunction  //函数也是一种类型，compareFunc负责比较元素的大小，具体实现交给上层开发者
    elemType     reflect.Type     //存储keys元素的实际类型，（运行时确定）
    omap         OrderedMapIntfs  //myKeys所归属的ordered_map
}

//让类型*myKeys实现KeysIntfs接口:
//首先，实现嵌入接口sort.Interface：Len，Less，Swap
func (keys *orderedKeys) Len() int{
    return len(keys.container)
}
func (keys *orderedKeys) Less(i, j int) bool{
    if keys.omap == nil {
        return keys.compareFunc(keys.container[i], keys.container[j], nil) < 0
    }else{
        omap := keys.omap.(*orderedMap)//暂时只支持orderedMap一种
        return keys.compareFunc(keys.container[i], keys.container[j], omap.m) < 0
    }
}
func (keys *orderedKeys) Swap(i, j int){
    keys.container[i], keys.container[j] = keys.container[j], keys.container[i]
}

//接着，实现接口Keys的其他方法
//判断k是否是可以存入myKeys.container的合法值
func (keys *orderedKeys) isAcceptableElem(k interface{}) bool {
    if k == nil {
        return false
    }
    //获取k的实际类型，与elemType进行比较
    if reflect.TypeOf(k) != keys.elemType {
        return false
    }
    return true
}
//Add方法
func (keys *orderedKeys) Add(k interface{}) bool {
    ok := keys.isAcceptableElem(k)
    if !ok {
        return false
    }
    keys.container = append(keys.container, k)
    sort.Sort(keys) //新元素加入进来之后，应该立刻进行一次排序！（因为*myKeys实现了sort.Interface接口，所以可以作为sort.Sort参数）
    return true
}
//Search方法，返回值已命名；利用了sort.Search方法
func (keys *orderedKeys) Search(k interface{}) (index int, contains bool) {
    ok := keys.isAcceptableElem(k)
    if !ok {
        return
    }
    //sort.Serach的第二个参数是匿名函数，功能是判断i对应的元素，是否>=要寻找的k值
    //仔细看sort.Search的源码发现，返回值index其实是k对应的索引id(存在)，或者是大于k的最小的索引id(不存在)
    if keys.omap == nil {
        index = sort.Search(keys.Len(), func(i int) bool { return keys.compareFunc(keys.container[i], k, nil) >= 0 })
    } else {
        omap := keys.omap.(*orderedMap) //暂时只支持orderedMap一种
        index = sort.Search(keys.Len(), func(i int) bool { return keys.compareFunc(keys.container[i], k, omap.m) >= 0 })
    }

    //由于index并非一定是找到了的索引id，所以要在此确认一下
    if index < keys.Len() && keys.container[index] == k {
        contains = true //给命名返回值赋值，golang特色
    }
    return
}
//Remove方法
func (keys *orderedKeys) Remove(k interface{}) bool {
    index, contains := keys.Search(k)
    if !contains {
        return false
    }
    //利用再切片的方式，实现切片删除指定元素，注意后面要有三个点
    keys.container = append(keys.container[0:index], keys.container[index+1:]...)
    return true
}
//Clear方法
func (keys *orderedKeys) Clear() {
    keys.container = make([]interface{}, 0)
}
//Get方法
func (keys *orderedKeys) Get(index int) interface{} {
    if index >= keys.Len() {
        return nil
    }
    return keys.container[index]
}
//GetAll，获得全部keys，放在一个slice中作为快照返回
func (keys *orderedKeys) GetAll() []interface{} {
    initialLen := len(keys.container)
    snapshot := make([]interface{}, initialLen)
    actualLen := 0
    for _, key := range keys.container{
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
//ElemType方法，获取key运行时的类型
func (keys *orderedKeys) ElemType() reflect.Type {
    return keys.elemType
}
//CompareFunc方法，获取运行时用于比较key大小的具体方法
func (keys *orderedKeys) CompareFunc() CompareFunction {
    return keys.compareFunc
}
//String方法，golang惯例，提供给fmt包
func (keys *orderedKeys) String() string {
    var buf bytes.Buffer
    buf.WriteString("Keys<")
    buf.WriteString(keys.elemType.Kind().String())
    buf.WriteString(">{")
    first := true
    buf.WriteString("[")
    for _, key := range keys.container {
        if first {
            first = false
        } else {
            buf.WriteString(" ")
        }
        buf.WriteString(fmt.Sprintf("%v", key))
    }
    buf.WriteString("]")
    buf.WriteString("}")
    return buf.String()
}

//golang惯例，“构造”函数
//返回值是KeysIntfs实现，所以是myKeys的指针
func NewKeys(compareFunc CompareFunction, elemType reflect.Type, omap OrderedMapIntfs) KeysIntfs {
    return &orderedKeys{
        container:    make([]interface{}, 0),
        compareFunc:  compareFunc,
        elemType:     elemType,
        omap:         omap,
    }
}