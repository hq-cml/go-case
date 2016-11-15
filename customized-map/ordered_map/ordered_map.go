package ordered_map

import (
    cmap "github.com/hq-cml/go-case/customized-map"
    "reflect"
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

/********orderdMap类型实现OrderedMapIntfs**********/
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



Keys() []interface{}                                       // 获取所有key所组成的切片值。
Vals() []interface{}                                       // 获取所有val所组成的切片值。
ToMap() map[interface{}]interface{}                        // 获取已包含的键值对所组成的字典值。
KeyType() reflect.Type                                     // 获取键的类型。
ValType() reflect.Type                                     // 获取元素的类型。