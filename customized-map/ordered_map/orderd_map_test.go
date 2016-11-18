package ordered_map

import (
    "testing"
    "reflect"
    "runtime/debug"
    "github.com/hq-cml/go-case/random"
)

func tmplTestOrderedMap(t *testing.T, omap OrderedMapIntfs, genKey func() interface{}, genElem func() interface{}, valKind reflect.Kind) {
    defer func() {
        if err := recover(); err != nil {
            debug.PrintStack()
            t.Errorf("Fatal Error: OrderedMap(type=%s): %s\n", valKind, err)
        }
    }()
    t.Logf("Starting TestOrderedMap<elemType=%s>...", valKind)

    // Basic
    //omap := newOrderedMap()

    expectedLen := 5
    testMap := make(map[interface{}]interface{}, expectedLen)
    var invalidKey interface{}
    for i := 0; i < expectedLen; i++ {
        key := genKey()
        testMap[key] = genElem()
        if invalidKey == nil {
            invalidKey = key
        }
    }
    for key, elem := range testMap {
        oldElem, ok := omap.Put(key, elem)
        if !ok {
            t.Errorf("ERROR: Put (%v, %v) to OrderedMap(elemType=%s) value %d is failing!\n", key, elem, valKind, omap)
            t.FailNow()
        }
        if oldElem != nil {
            t.Errorf("ERROR: Already had a (%v, %v) in OrderedMap(elemType=%s) value %d!\n", key, elem, valKind, omap)
            t.FailNow()
        }
        t.Logf("Put (%v, %v) to the OrderedMap(elemType=%s) value %v.", key, elem, valKind, omap)
    }
    if omap.Len() != expectedLen {
        t.Errorf("ERROR: The length of OrderedMap(elemType=%s) value %d is not %d!\n", valKind, omap.Len(), expectedLen)
        t.FailNow()
    }

    for key, elem := range testMap {
        contains := omap.Contains(key)
        if !contains {
            t.Errorf("ERROR: The OrderedMap(elemType=%s) value %v do not contains %v!", valKind, omap, key)
            t.FailNow()
        }
        actualElem := omap.Get(key)
        if actualElem == nil {
            t.Errorf("ERROR: The OrderedMap(elemType=%s) value %v do not contains %v!", valKind, omap, key)
            t.FailNow()
        }
        t.Logf("The OrderedMap(elemType=%s) value %v contains key %v.", valKind, omap, key)
        if actualElem != elem {
            t.Errorf("ERROR: The element of OrderedMap(elemType=%s) value %v with key %v do not equals %v!\n", valKind, actualElem, key, elem)
            t.FailNow()
        }
        t.Logf("The element of OrderedMap(elemType=%s) value %v to key %v is %v.", valKind, omap, key, actualElem)
    }
    oldElem := omap.Remove(invalidKey)
    if oldElem == nil {
        t.Errorf("ERROR: Remove %v from OrderedMap(elemType=%s) value %d is failing!\n",
            invalidKey, valKind, omap)
        t.FailNow()
    }
    t.Logf("Removed (%v, %v) from the OrderedMap(elemType=%s) value %v.", invalidKey, oldElem, valKind, omap)
    delete(testMap, invalidKey)

    // Type
    actualElemType := omap.ValType()
    if actualElemType == nil {
        t.Errorf("ERROR: The element type of OrderedMap(elemType=%s) value is nil!\n", valKind)
        t.FailNow()
    }
    actualElemKind := actualElemType.Kind()
    if actualElemKind != valKind {
        t.Errorf("ERROR: The element type of OrderedMap(elemType=%s) value %s is not %s!\n", valKind, actualElemKind, valKind)
        t.FailNow()
    }
    t.Logf("The element type of OrderedMap(elemType=%s) value %v is %s.", valKind, omap, actualElemKind)
    actualKeyKind := omap.KeyType().Kind()
    keyKind := reflect.TypeOf(genKey()).Kind()
    if actualKeyKind != valKind {
        t.Errorf("ERROR: The key type of OrderedMap(elemType=%s) value %s is not %s!\n", keyKind, actualKeyKind, keyKind)
        t.FailNow()
    }
    t.Logf("The key type of OrderedMap(elemType=%s) value %v is %s.",  keyKind, omap, actualKeyKind)

    // Export
    keys := omap.Keys()
    vals := omap.Vals()
    pairs := omap.ToMap()
    for key, val := range testMap {
        var hasKey bool
        for _, k := range keys {
            if k == key {
                hasKey = true
            }
        }
        if !hasKey {
            t.Errorf("ERROR: The keys of OrderedMap(elemType=%s) value %v do not contains %v!\n", valKind, omap, key)
            t.FailNow()
        }
        var hasVal bool
        for _, e := range vals {
            if e == val {
                hasVal = true
            }
        }
        if !hasVal {
            t.Errorf("ERROR: The elems of OrderedMap(elemType=%s) value %v do not contains %v!\n", valKind, omap, val)
            t.FailNow()
        }
        var hasPair bool
        for k, e := range pairs {
            if k == key && e == val {
                hasPair = true
            }
        }
        if !hasPair {
            t.Errorf("ERROR: The elems of OrderedMap(elemType=%s) value %v do not contains (%v, %v)!\n",
                valKind, omap, key, val)
            t.FailNow()
        }
    }

    // Advance
    fKey := omap.FirstKey()
    if fKey != keys[0] {
        t.Errorf("ERROR: The first key of OrderedMap(elemType=%s) value %v is not equals %v!\n", valKind, fKey, keys[0])
        t.FailNow()
    }
    t.Logf("The first key of OrderedMap(elemType=%s) value %v is %s.", valKind, omap, fKey)
    lKey := omap.LastKey()
    if lKey != keys[len(keys)-1] {
        t.Errorf("ERROR: The last key of OrderedMap(elemType=%s) value %v is not equals %v!\n", valKind, lKey, keys[len(keys)-1])
        t.FailNow()
    }
    t.Logf("The last key of OrderedMap(elemType=%s) value %v is %s.", valKind, omap, lKey)
    endIndex := len(keys)/2 + 1
    toKey := keys[endIndex]
    headMap := omap.HeadMap(toKey)
    headKeys := headMap.Keys()
    for i := 0; i < endIndex; i++ {
        hKey := headKeys[i]
        tempKey := keys[i]
        if hKey != tempKey {
            t.Errorf("ERROR: The key of OrderedMap(elemType=%s) value %v with index %d is not equals %v!\n", valKind, tempKey, i, hKey)
            t.FailNow()
        }
    }
    beginIndex := len(keys)/2 - 1
    endIndex = len(keys) - 1
    fromKey := keys[beginIndex]
    tailMap := omap.TailMap(fromKey)
    tailKeys := tailMap.Keys()
    for i := beginIndex; i < endIndex; i++ {
        tKey := tailKeys[i-beginIndex]
        tempKey := keys[i]
        if tKey != tempKey {
            t.Errorf("ERROR: The key of OrderedMap(elemType=%s) value %v with index %d is not equals %v!\n", valKind, tempKey, i, tKey)
            t.FailNow()
        }
    }
    beginIndex = len(keys)/2 - 1
    endIndex = len(keys)/2 + 1
    fromKey = keys[beginIndex]
    toKey = keys[endIndex]
    subMap := omap.SubMap(fromKey, toKey)
    subKeys := subMap.Keys()
    for i := beginIndex; i < endIndex; i++ {
        sKey := subKeys[i-beginIndex]
        tempKey := keys[i]
        if sKey != tempKey {
            t.Errorf("ERROR: The key of OrderedMap(elemType=%s) value %v with index %d is not equals %v!\n", valKind, tempKey, i, sKey)
            t.FailNow()
        }
    }

    // Clear
    omap.Clear()
    if omap.Len() != 0 {
        t.Errorf("ERROR: Clear OrderedMap(elemType=%s) value %d is failing!\n", valKind, omap)
        t.FailNow()
    }
    t.Logf("The OrderedMap(elemType=%s) value %v has been cleared.", valKind, omap)
}

//测试按Key排序的map
func TestInt64OrderedKeyMap(t *testing.T) {
    keys := NewKeys(compareInt64Key, reflect.TypeOf(int64(1)), nil)

    //按key排序的map
    omap := NewOrderedMap(keys, reflect.TypeOf(int64(1)))

    tmplTestOrderedMap(
        t,
        omap,
        func() interface{} { return random.GenRandInt(1000) },
        func() interface{} { return random.GenRandInt(1000) },
        reflect.Int64)
}

//测试按value排序的map
func TestInt64OrderedMap(t *testing.T) {
    keys := NewKeys(compareInt64Val, reflect.TypeOf(int64(1)), nil)

    //按val排序的map
    omap := NewOrderedMap(keys, reflect.TypeOf(int64(1)))

    //！！！关键点：将omap的m放入keys，否则无法实现按值排序
    okeys := keys.(*orderedKeys)
    okeys.baseMap = omap.(*orderedMap).m

    tmplTestOrderedMap(
        t,
        omap,
        func() interface{} { return random.GenRandInt(1000) },
        func() interface{} { return random.GenRandInt(1000) },
        reflect.Int64)
}
