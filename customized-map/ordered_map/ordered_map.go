package ordered_map

import cmap "github.com/hq-cml/go-case/customized-map"

// 有序的Map的接口类型。
type OrderedMap interface {
    cmap.GenericMapIntfs  // 泛化的Map接口
    // 获取第一个键值。若无任何键值对则返回nil。
    FirstKey() interface{}
    // 获取最后一个键值。若无任何键值对则返回nil。
    LastKey() interface{}
    // 获取由小于键值toKey的键值所对应的键值对组成的OrderedMap类型值。
    HeadMap(toKey interface{}) OrderedMap
    // 获取由小于键值toKey且大于等于键值fromKey的键值所对应的键值对组成的OrderedMap类型值。
    SubMap(fromKey interface{}, toKey interface{}) OrderedMap
    // 获取由大于等于键值fromKey的键值所对应的键值对组成的OrderedMap类型值。
    TailMap(fromKey interface{}) OrderedMap
}