package customized_map
/*
 * 定制化map之：并发安全的map实现
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

//*ConcurrentMap实现ConcurrentMapIntfs


//惯例New函数
func NewConcurrentMap(keyType, valType reflect.Type) ConcurrentMapIntfs {
    return &concurrentMap{
        keyType: keyType,
        valType: valType,
        m      : make(map[interface{}]interface{}),
    }
}
