package customized_map
/*
 * 定制化map之：并发安全的map实现: ConcurrentMap
 */
import (
    "bytes"
    "fmt"
    "reflect"
    "sync"
)

//并发map接口，接口的嵌套（可以理解为继承关系）
type ConcurrentMapIntfs interface {
    GenericMapIntfs
}

//*concurrentMap实现ConcurrentMapIntfs
type concurrentMap struct {
    m         map[interface{}]interface{}  //m是实际map的句柄
    keyType   reflect.Type
    valType   reflect.Type
    rwmutext  sync.RWMutex                 //配一把读写锁，保证并发安全
}

//校验k，v对儿是否为空即动态类型是否符合要求
func (cmap *concurrentMap) checkPair(k, v interface{}) bool {
    if k == nil || reflect.TypeOf(k) != cmap.keyType {
        return false
    }

    if v == nil || reflect.TypeOf(v) != cmap.valType {
        return false
    }
    return true
}

/************** ConcurrentMap实现ConcurrentMapIntfs **********************/
// 获取给定键值对应的元素值。若没有对应元素值则返回nil。
func (cmap *concurrentMap) Get(key interface{}) interface{} {
    //读锁
    cmap.rwmutext.RLock()
    defer cmap.rwmutext.RUnlock()
    return cmap.m[key]
}

//添加键值对，并返回与给定键值对应的旧的元素值。若没有旧元素值则返回(nil, true)。
func (cmap *concurrentMap) Put(key interface{}, val interface{}) (interface{}, bool) {
    if cmap.checkPair(key, val) {
        return nil, false
    }
    //写锁
    cmap.rwmutext.Lock()
    defer cmap.rwmutext.Unlock()

    old_val := cmap.m[key]
    cmap.m[key] = val
    return old_val, true
}

// 删除与给定键值对应的键值对，并返回旧的元素值。若没有旧元素值则返回nil
func (cmap *concurrentMap) Remove(key interface{}) interface{} {
    //写锁
    cmap.rwmutext.Lock()
    defer cmap.rwmutext.Unlock()
    old_val := cmap.m[key]
    delete(cmap.m, key)
    return old_val
}

// 清除所有的键值对。
func (cmap *concurrentMap) Clear() {
    //写锁
    cmap.rwmutext.Lock()
    defer cmap.rwmutext.Unlock()
    cmap.m = make(map[interface{}]interface{})
}

// 获取键值对的数量。
func (cmap *concurrentMap) Len() int {
    //读锁
    cmap.rwmutext.RLock()
    defer cmap.rwmutext.RUnlock()
    return len(cmap.m)
}

// 判断是否包含给定的键值。
func (cmap *concurrentMap) Contains(key interface{}) bool {
    //读锁
    cmap.rwmutext.RLock()
    defer cmap.rwmutext.RUnlock()
    _, ok := cmap.m[key] //判断字典是否存在的方法
    return ok
}

// 获取已排序的key所组成的切片值。
func (cmap *concurrentMap) Keys() []interface{} {
    //读锁
    cmap.rwmutext.RLock()
    defer cmap.rwmutext.RUnlock()
    keys := make([]interface{}, len(cmap.m))
    idx := 0
    for k,_ := range cmap.m {
        keys[idx] = k
        idx++
    }
    return keys
}

// 获取已排序的元素值所组成的切片值。
func (cmap *concurrentMap) Vals() []interface{} {
    //读锁
    cmap.rwmutext.RLock()
    defer cmap.rwmutext.RUnlock()
    vals := make([]interface{}, len(cmap.m))
    idx := 0
    for _,v := range cmap.m {
        vals[idx] = v
        idx++
    }
    return vals
}

// 获取已包含的键值对所组成的字典值。
func (cmap *concurrentMap) ToMap() map[interface{}]interface{} {
    //读锁
    cmap.rwmutext.RLock()
    defer cmap.rwmutext.RUnlock()
    replica := make(map[interface{}]interface{})
    for k, v := range cmap.m {
        replica[k] = v
    }
    return replica
}

// 获取键的类型。
func (cmap *concurrentMap) KeyType() reflect.Type {
    return cmap.keyType
}

// 获取元素的类型。
func (cmap *concurrentMap) ValType() reflect.Type {
    return cmap.valType
}

//惯例New函数
func NewConcurrentMap(keyType, valType reflect.Type) ConcurrentMapIntfs {
    return &concurrentMap{
        keyType: keyType,
        valType: valType,
        m      : make(map[interface{}]interface{}),
    }
}

//String方法
func (cmap *concurrentMap) String() string {
    var buf bytes.Buffer
    buf.WriteString("ConcurrentMap<")
    buf.WriteString(cmap.keyType.Kind().String())
    buf.WriteString(",")
    buf.WriteString(cmap.valType.Kind().String())
    buf.WriteString(">{")
    first := true
    //读锁
    cmap.rwmutext.RLock()
    defer cmap.rwmutext.RUnlock()
    for k, v := range cmap.m {
        if first {
            first = false
        } else {
            buf.WriteString(" ")
        }
        buf.WriteString(fmt.Sprintf("%v", k))
        buf.WriteString(":")
        buf.WriteString(fmt.Sprintf("%v", v))
    }
    buf.WriteString("}")
    return buf.String()
}