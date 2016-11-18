package ordered_map

import (
    cmap "github.com/hq-cml/go-case/customized-map"
    "reflect"
    "bytes"
    "fmt"
)

//有序的Map的接口类型。
type OrderedMapIntfs interface {
    cmap.GenericMapIntfs                                            //嵌入泛化的Map接口，类似于继承

    FirstKey() interface{}                                          //获取第一个键值。若无任何键值对则返回nil。
    LastKey() interface{}                                           //获取最后一个键值。若无任何键值对则返回nil。
    HeadMap(toKey interface{}) OrderedMapIntfs                      //获取由小于键值toKey的键值所对应的键值对组成的OrderedMap类型值。
    SubMap(fromKey interface{}, toKey interface{}) OrderedMapIntfs  //获取由小于键值toKey且大于等于键值fromKey的键值所对应的键值对组成的OrderedMap类型值。
    TailMap(fromKey interface{}) OrderedMapIntfs                    //获取由大于等于键值fromKey的键值所对应的键值对组成的OrderedMap类型值。
}

type orderedMap struct {
    keys    KeysIntfs
    valType reflect.Type
    m       map[interface{}]interface{}
}

func (omap *orderedMap) isAcceptableVal(v interface{}) bool {
    if v == nil {
        return false
    }
    if reflect.TypeOf(v) != omap.valType {
        return false
    }
    return true
}

/******* //*orderdMap类型实现OrderedMapIntfs**********/
//获取给定键值对应的元素值。若没有对应元素值则返回nil。
func (omap *orderedMap) Get(key interface{}) interface{} {
    return omap.m[key]
}
//添加键值对，并返回与给定键值对应的旧的元素值。若没有旧元素值则返回(nil, true)
func (omap *orderedMap) Put(key interface{}, v interface{}) (interface{}, bool) {
    if !omap.isAcceptableVal(v) {
        return nil, false
    }
    oldElem, ok := omap.m[key]
    omap.m[key] = v
    if !ok {
        omap.keys.Add(key)
    }
    return oldElem, true
}
//删除与给定键值对应的键值对，并返回旧的元素值。若没有旧元素值则返回nil
func (omap *orderedMap) Remove(key interface{}) interface{} {
    oldElem, ok := omap.m[key] //通过ok判断是否存在
    delete(omap.m, key)
    if ok {
        omap.keys.Remove(key)
    }
    return oldElem
}
//清除所有的键值对
func (omap *orderedMap) Clear() {
    omap.m = make(map[interface{}]interface{})
    omap.keys.Clear()
}
//获取键值对的数量
func (omap *orderedMap) Len() int {
    return len(omap.m)
}
//判断是否包含给定的键值
func (omap *orderedMap) Contains(key interface{}) bool {
    _, ok := omap.m[key]
    return ok
}
//获取所有key所组成的切片值
func (omap *orderedMap) Keys() []interface{} {
    return omap.keys.GetAll()
}
//获取所有val所组成的切片值。
func (omap *orderedMap) Vals() []interface{} {
    initialLen := omap.Len()
    vals := make([]interface{}, initialLen)
    actualLen := 0
    for _, key := range omap.keys.GetAll() {
        val := omap.m[key]
        if actualLen < initialLen {
            vals[actualLen] = val
        } else {
            vals = append(vals, val)
        }
        actualLen++
    }
    if actualLen < initialLen {
        vals = vals[:actualLen]
    }
    return vals
}
//获取已包含的键值对所组成的字典值。类似于得到一个快照
func (omap *orderedMap) ToMap() map[interface{}]interface{} {
    replica := make(map[interface{}]interface{})
    for k, v := range omap.m {
        replica[k] = v
    }
    return replica
}
//获取键的类型
func (omap *orderedMap) KeyType() reflect.Type {
    return omap.keys.ElemType()
}
//获取元素的类型
func (omap *orderedMap) ValType() reflect.Type {
    return omap.valType
}
////////////////实现扩展的方法
//获取第一个键值。若无任何键值对则返回nil。
func (omap *orderedMap) FirstKey() interface{} {
    if omap.Len() == 0 {
        return nil
    }
    return omap.keys.Get(0)
}
//获取最后一个键值。若无任何键值对则返回nil。
func (omap *orderedMap) LastKey() interface{} {
    length := omap.Len()
    if length == 0 {
        return nil
    }
    return omap.keys.Get(length - 1)
}
//获取由小于键值toKey且大于等于键值fromKey的键值所对应的键值对组成的OrderedMap类型值。其实本质上就是将map切出一部分来
func (omap *orderedMap) SubMap(fromKey interface{}, toKey interface{}) OrderedMapIntfs {
    newOmap := &orderedMap{
        keys:     NewKeys(omap.keys.CompareFunc(), omap.keys.ElemType(), omap.m),
        valType:  omap.valType,
        m:        make(map[interface{}]interface{})}
    omapLen := omap.Len()
    if omapLen == 0 {
        return newOmap
    }
    beginIndex, contains := omap.keys.Search(fromKey)
    if !contains {
        beginIndex = 0
    }
    endIndex, contains := omap.keys.Search(toKey)
    if !contains {
        endIndex = omapLen
    }
    var key, elem interface{}
    for i := beginIndex; i < endIndex; i++ {
        key = omap.keys.Get(i)
        elem = omap.m[key]
        newOmap.Put(key, elem)
    }
    return newOmap
}
//获取由小于键值toKey的键值所对应的键值对组成的OrderedMap类型值。其实就是取出map的某个点的前半部分
func (omap *orderedMap) HeadMap(toKey interface{}) OrderedMapIntfs {
    return omap.SubMap(nil, toKey)
}
//获取由大于等于键值fromKey的键值所对应的键值对组成的OrderedMap类型值。其实就是取出map的某个点的后半部分
func (omap *orderedMap) TailMap(fromKey interface{}) OrderedMapIntfs {
    return omap.SubMap(fromKey, nil)
}
//String
func (omap *orderedMap) String() string {
    var buf bytes.Buffer
    buf.WriteString("OrderedMap<")
    buf.WriteString(omap.keys.ElemType().Kind().String())
    buf.WriteString(",")
    buf.WriteString(omap.valType.Kind().String())
    buf.WriteString(">{")
    first := true
    omapLen := omap.Len()
    for i := 0; i < omapLen; i++ {
        if first {
            first = false
        } else {
            buf.WriteString(" ")
        }
        key := omap.keys.Get(i)
        buf.WriteString(fmt.Sprintf("%v", key))
        buf.WriteString(":")
        buf.WriteString(fmt.Sprintf("%v", omap.m[key]))
    }
    buf.WriteString("}")
    return buf.String()
}
//惯例
func NewOrderedMap(keys KeysIntfs, valType reflect.Type) OrderedMapIntfs {
    return &orderedMap{
        keys:     keys,
        valType: valType,
        m:        make(map[interface{}]interface{})}
}