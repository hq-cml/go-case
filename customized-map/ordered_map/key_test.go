package ordered_map

import (
    "sort"
    "testing"
    "reflect"
    "runtime/debug"
    "github.com/hq-cml/go-case/random"
)

//测试Keys的模板函数
func tmplTestKeys(t *testing.T, keys KeysIntfs, genKey func() interface{}, elemKind reflect.Kind) {
    defer func() {
        if err := recover(); err != nil {
            debug.PrintStack()
            t.Errorf("Fatal Error: Keys(%s): %s\n", elemKind, err)
        }
    }()
    t.Logf("Starting TestKeys<%s>...", elemKind)
    //keys := newKeys()

    //生成5个随机key
    expectedLen := 5
    tmpKeys := make([]interface{}, expectedLen)
    for i := 0; i < expectedLen; i++ {
        tmpKeys[i] = genKey()
    }

    //测试Add
    for _, key := range tmpKeys {
        result := keys.Add(key)
        if !result {
            t.Errorf("ERROR: Add %v to Keys(%s) value %d is failing!\n", key, elemKind, keys)
            t.FailNow()
        }
        t.Logf("Added %v to the Keys(%s) value %v.", key, elemKind, keys)
    }

    //测试Search和Get
    for _, key := range tmpKeys {
        index, contains := keys.Search(key)
        if !contains {
            t.Errorf("ERROR: The Keys(%s) value %v do not contains %v!",  elemKind, keys, key)
            t.FailNow()
        }
        t.Logf("The Keys(%s) value %v contains key %v.", elemKind, keys, key)
        actualElem := keys.Get(index)
        if actualElem != key {
            t.Errorf("ERROR: The element of Keys(%s) value %v with index %d do not equals %v!\n", elemKind, actualElem, index, key)
            t.FailNow()
        }
        t.Logf("The element of Keys(%s) value %v with index %d is %v.",  elemKind, keys, index, actualElem)
    }

    //测试revomve
    invalidElem := tmpKeys[len(tmpKeys)/2]
    result := keys.Remove(invalidElem)
    if !result {
        t.Errorf("ERROR: Remove %v from Keys(%s) value %d is failing!\n",
            invalidElem, elemKind, keys)
        t.FailNow()
    }
    t.Logf("Removed %v from the Keys(%s) value %v.", invalidElem, elemKind, keys)
    if !sort.IsSorted(keys) {
        t.Errorf("ERROR: The Keys(%s) value %v is not sorted yet?!\n", elemKind, keys)
        t.FailNow()
    }
    t.Logf("The Keys(%s) value %v is sorted.", elemKind, keys)

    //测试Kind和Type是否一致
    actualElemType := keys.ElemType()
    if actualElemType == nil {
        t.Errorf("ERROR: The element type of Keys(%s) value is nil!\n", elemKind)
        t.FailNow()
    }
    actualElemKind := actualElemType.Kind()
    if actualElemKind != elemKind {
        t.Errorf("ERROR: The element type of Keys(%s) value %s is not %s!\n", elemKind, actualElemKind, elemKind)
        t.FailNow()
    }
    t.Logf("The element type of Keys(%s) value %v is %s.", elemKind, keys, actualElemKind)

    currCompFunc := keys.CompareFunc()
    if currCompFunc == nil {
        t.Errorf("ERROR: The compare function of Keys(%s) value is nil!\n", elemKind)
        t.FailNow()
    }

    //测试Clear和Len
    keys.Clear()
    if keys.Len() != 0 {
        t.Errorf("ERROR: Clear Keys(%s) value %d is failing!\n", elemKind, keys)
        t.FailNow()
    }
    t.Logf("The Keys(%s) value %v have been cleared.", elemKind, keys)
}


//func NewKeys(compareFunc CompareFunction, elemType reflect.Type) KeysIntfs
//func tmplTestKeys(t *testing.T, newKeys func() KeysIntfs, genKey func() interface{}, elemKind reflect.Kind)

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

//Int64测试
func TestInt64Keys(t *testing.T) {
    keys := NewKeys(compareInt64Key, reflect.TypeOf(int64(1)));
    //调用测试模板
    tmplTestKeys(t,                                                          //参数1
        keys,                                                                //参数2
        func() interface{} { return random.GenRandInt(1000) },               //参数3
        reflect.Int64)                                                       //参数4
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

//String测试
func TestStringKeys(t *testing.T) {
    keys := NewKeys(compareStringKey, reflect.TypeOf(string(1)));
    //调用测试模板
    tmplTestKeys(t,
        keys,
        func() interface{} { return random.GenRandString(10) },
        reflect.String)
}