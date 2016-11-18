package ordered_map

//int64map的key比较函数
func compareInt64Key(e1 interface{}, e2 interface{}, m map[interface{}]interface{}) int8{
    k1 := e1.(int64)
    k2 := e2.(int64)
    if k1 < k2 {
        return -1
    } else if k1 > k2 {
        return 1
    } else {
        return 0
    }
}

//int64的val比较函数，比较map的val
func compareInt64Val(k1 interface{}, k2 interface{}, m map[interface{}]interface{}) int8{
    var v1 int64
    var v2 int64
    if v, ok := m[k1]; ok{
        v1 = v.(int64)
    } else {
        return -1
    }
    if v, ok := m[k2]; ok{
        v2 = v.(int64)
    } else {
        return -1
    }

    if v1 < v2 {
        return -1
    } else if v1 > v2 {
        return 1
    } else {
        return 0
    }
}

//string的key比较函数
func compareStringKey(e1 interface{}, e2 interface{}, m map[interface{}]interface{}) int8{
    k1 := e1.(string)
    k2 := e2.(string)
    if k1 < k2 {
        return -1
    } else if k1 > k2 {
        return 1
    } else {
        return 0
    }
}
