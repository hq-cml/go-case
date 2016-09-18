package customized_map

import "reflect"

//泛化的定制Map的接口类型
type GenericMapIntfs interface {
    Get(key interface{}) interface{}                           // 获取给定键值对应的元素值。若没有对应元素值则返回nil。
    Put(key interface{}, val interface{}) (interface{}, bool)  // 添加键值对，并返回与给定键值对应的旧的元素值。若没有旧元素值则返回(nil, true)。
    Remove(key interface{}) interface{}                        // 删除与给定键值对应的键值对，并返回旧的元素值。若没有旧元素值则返回nil。
    Clear()                                                    // 清除所有的键值对。
    Len() int                                                  // 获取键值对的数量。
    Contains(key interface{}) bool                             // 判断是否包含给定的键值。
    Keys() []interface{}                                       // 获取所有key所组成的切片值。
    Vals() []interface{}                                       // 获取所有val所组成的切片值。
    ToMap() map[interface{}]interface{}                        // 获取已包含的键值对所组成的字典值。
    KeyType() reflect.Type                                     // 获取键的类型。
    ValType() reflect.Type                                     // 获取元素的类型。
}