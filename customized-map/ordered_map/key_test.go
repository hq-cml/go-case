package ordered_map

import (
    "sort"
    "testing"
    "reflect"
    "runtime/debug"
)

//测试Keys的模板函数
func tmplTestKeys(t *testing.T, newKeys func() KeysIntfs, genKey func() interface{}, elemKind reflect.Kind) {
    defer func() {
        if err := recover(); err != nil {
            debug.PrintStack()
            t.Errorf("Fatal Error: Keys(%s): %s\n", elemKind, err)
        }
    }()
    t.Logf("Starting TestKeys<%s>...", elemKind)
    keys := newKeys()

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
