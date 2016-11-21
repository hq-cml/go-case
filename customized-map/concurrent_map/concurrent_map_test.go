package concurrent_map

import (
    "testing"
    "reflect"
    "fmt"
    "runtime/debug"
    "math/rand"
    random "github.com/hq-cml/go-case/random"
)

/*************** 功能测试 *************/
//测试Int64型的cmap
func TestInt64Cmap(t *testing.T) {
    newFunc := func() ConcurrentMapIntfs {
        keyType := reflect.TypeOf(int64(2))
        valType := keyType
        return NewConcurrentMap(keyType, valType)
    }

    genFunc := func() interface{} { return rand.Int63n(1000) }

    test(t, newFunc, genFunc, genFunc, reflect.Int64,reflect.Int64)
}
//测试Float64型的cmap
func TestFloat64Cmap(t *testing.T) {
    newFunc := func() ConcurrentMapIntfs {
        keyType := reflect.TypeOf(float64(2))
        valType := keyType
        return NewConcurrentMap(keyType, valType)
    }
    genFunc := func() interface{} { return rand.Float64() }

    test(t, newFunc, genFunc, genFunc, reflect.Float64, reflect.Float64)
}
//测试string型cmap
func TestStringCmap(t *testing.T) {
    newCmap := func() ConcurrentMapIntfs {
        keyType := reflect.TypeOf(string(2))
        valType := keyType
        return NewConcurrentMap(keyType, valType)
    }
    genFunc := func() interface{} { return random.GenRandString(10) }
    test(t, newCmap, genFunc, genFunc, reflect.String, reflect.String)
}

func test(t *testing.T, newConcurrentMap func() ConcurrentMapIntfs, genKey func() interface{}, genVal func() interface{}, keyKind reflect.Kind, valKind reflect.Kind) {
    mapType := fmt.Sprintf("ConcurrentMap<keyType=%s, elemType=%s>", keyKind, valKind)
    defer func() {
        if err := recover(); err != nil {
            debug.PrintStack()
            t.Errorf("Fatal Error: %s: %s\n", mapType, err)
        }
    }()
    t.Logf("Starting Test%s...", mapType)

    // Basic
    cmap := newConcurrentMap()
    expectedLen := 0
    if cmap.Len() != expectedLen {
        t.Errorf("ERROR: The length of %s value %d is not %d!\n", mapType, cmap.Len(), expectedLen)
        t.FailNow()
    }
    expectedLen = 5
    testMap := make(map[interface{}]interface{}, expectedLen)
    var invalidKey interface{}
    for i := 0; i < expectedLen; i++ {
        key := genKey()
        testMap[key] = genVal()
        if invalidKey == nil {
            invalidKey = key
        }
    }
    for key, val := range testMap {
        oldVal, ok := cmap.Put(key, val)
        if !ok {
            t.Errorf("ERROR: Put (%v, %v) to %s value %d is failing!\n", key, val, mapType, cmap)
            t.FailNow()
        }
        if oldVal != nil {
            t.Errorf("ERROR: Already had a (%v, %v) in %s value %d!\n", key, val, mapType, cmap)
            t.FailNow()
        }
        t.Logf("Put (%v, %v) to the %s value %v.", key, val, mapType, cmap)
    }
    if cmap.Len() != expectedLen {
        t.Errorf("ERROR: The length of %s value %d is not %d!\n", mapType, cmap.Len(), expectedLen)
        t.FailNow()
    }
    for key, val := range testMap {
        contains := cmap.Contains(key)
        if !contains {
            t.Errorf("ERROR: The %s value %v do not contains %v!", mapType, cmap, key)
            t.FailNow()
        }
        actualVal := cmap.Get(key)
        if actualVal == nil {
            t.Errorf("ERROR: The %s value %v do not contains %v!", mapType, cmap, key)
            t.FailNow()
        }
        t.Logf("The %s value %v contains key %v.", mapType, cmap, key)
        if actualVal != val {
            t.Errorf("ERROR: The element of %s value %v with key %v do not equals %v!\n", mapType, actualVal, key, val)
            t.FailNow()
        }
        t.Logf("The element of %s value %v to key %v is %v.", mapType, cmap, key, actualVal)
    }
    oldVal := cmap.Remove(invalidKey)
    if oldVal == nil {
        t.Errorf("ERROR: Remove %v from %s value %d is failing!\n", invalidKey, mapType, cmap)
        t.FailNow()
    }
    t.Logf("Removed (%v, %v) from the %s value %v.", invalidKey, oldVal, mapType, cmap)
    delete(testMap, invalidKey)

    // Type
    actualValType := cmap.ValType()
    if actualValType == nil {
        t.Errorf("ERROR: The element type of %s value is nil!\n", mapType)
        t.FailNow()
    }
    actualValKind := actualValType.Kind()
    if actualValKind != valKind {
        t.Errorf("ERROR: The element type of %s value %s is not %s!\n", mapType, actualValKind, valKind)
        t.FailNow()
    }
    t.Logf("The element type of %s value %v is %s.", mapType, cmap, actualValKind)
    actualKeyKind := cmap.KeyType().Kind()
    if actualKeyKind != keyKind {
        t.Errorf("ERROR: The key type of %s value %s is not %s!\n", mapType, actualKeyKind, keyKind)
        t.FailNow()
    }
    t.Logf("The key type of %s value %v is %s.", mapType, cmap, actualKeyKind)

    // Export
    keys := cmap.Keys()
    vals := cmap.Vals()
    pairs := cmap.ToMap()
    for key, elem := range testMap {
        var hasKey bool
        for _, k := range keys {
            if k == key {
                hasKey = true
            }
        }
        if !hasKey {
            t.Errorf("ERROR: The keys of %s value %v do not contains %v!\n", mapType, cmap, key)
            t.FailNow()
        }
        var hasVal bool
        for _, e := range vals {
            if e == elem {
                hasVal = true
            }
        }
        if !hasVal {
            t.Errorf("ERROR: The elems of %s value %v do not contains %v!\n", mapType, cmap, elem)
            t.FailNow()
        }
        var hasPair bool
        for k, e := range pairs {
            if k == key && e == elem {
                hasPair = true
            }
        }
        if !hasPair {
            t.Errorf("ERROR: The elems of %s value %v do not contains (%v, %v)!\n",
                mapType, cmap, key, elem)
            t.FailNow()
        }
    }

    // Clear
    cmap.Clear()
    if cmap.Len() != 0 {
        t.Errorf("ERROR: Clear %s value %d is failing!\n", mapType, cmap)
        t.FailNow()
    }
    t.Logf("The %s value %v has been cleared.", mapType, cmap)
}

/***************性能测试***************/
func BenchmarkConcurrentMap(b *testing.B) {
    keyType := reflect.TypeOf(int32(2))
    valType := keyType
    cmap := NewConcurrentMap(keyType, valType)
    var key, val int32
    //fmt.Printf("N=%d.\n", b.N)
    //StartTimer在benchmark函数开始的时候，是自动执行的，所以这里reset
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        b.StopTimer()
        seed := int32(i)
        key = seed
        val = seed << 10
        b.StartTimer()
        cmap.Put(key, val)
        _ = cmap.Get(key)
        b.StopTimer()
        b.SetBytes(8)
        b.StartTimer()
    }
    ml := cmap.Len()
    b.StopTimer()
    mapType := fmt.Sprintf("N is %d, ConcurrentMap<%s, %s>", b.N,
        keyType.Kind().String(), valType.Kind().String())
    b.Logf("The length of %s value is %d.\n", mapType, ml)
    b.StartTimer()
}

func BenchmarkMap(b *testing.B) {
    keyType := reflect.TypeOf(int32(2))
    valType := keyType
    imap := make(map[interface{}]interface{})
    var key, elem int32
    //fmt.Printf("N=%d.\n", b.N)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        b.StopTimer()
        seed := int32(i)
        key = seed
        elem = seed << 10
        b.StartTimer()
        imap[key] = elem
        b.StopTimer()
        _ = imap[key]
        b.StopTimer()
        b.SetBytes(8)
        b.StartTimer()
    }
    ml := len(imap)
    b.StopTimer()
    mapType := fmt.Sprintf("N is %d, Map<%s, %s>", b.N,
        keyType.Kind().String(), valType.Kind().String())
    b.Logf("The length of %s value is %d.\n", mapType, ml)
    b.StartTimer()
}