package set

/*
 * Set接口类型，规定了一个Set拥有的基本操作
 */
type SetIntfs interface {
    Add(e interface{}) bool
    Remove(e interface{})
    Clear()
    Contains(e interface{}) bool
    Len() int
    Same(other SetIntfs) bool
    Elements() []interface{}
    String() string
}

/*
 * Set接口类型的高级方法，作为公用函数，没必要每种实现各自实现
 */
// 判断集合 self 是否是集合 other 的超集
func IsSuperset(self SetIntfs, other SetIntfs) bool {
    if self == nil || other == nil {
        return false
    }
    selfLen := self.Len()
    otherLen := other.Len()

    if otherLen == 0 {
        return true
    }

    if selfLen < otherLen {
        return false
    }

    for _, v := range other.Elements() {
        if !self.Contains(v) {
            return false
        }
    }
    return true
}

// 生成集合 self 和集合 other 的并集
func Union(self SetIntfs, other SetIntfs) SetIntfs {
    if self == nil || other == nil {
        return nil
    }
    unionedSet := NewSimpleSet()
    for _, v := range self.Elements() {
        unionedSet.Add(v)
    }

    for _, v := range other.Elements() {
        unionedSet.Add(v) //Add可以保证相同的元素无法插入的
    }
    return unionedSet
}

// 生成集合 self 和集合 other 的交集
func Intersect(self SetIntfs, other SetIntfs) SetIntfs {
    if self == nil || other == nil {
        return nil
    }
    intersectedSet := NewSimpleSet()
    if self.Len() == 0 || other.Len() == 0 {
        return intersectedSet
    }
    if self.Len() < other.Len() {
        for _, v := range self.Elements() {
            if other.Contains(v) {
                intersectedSet.Add(v)
            }
        }
    } else {
        for _, v := range other.Elements() {
            if self.Contains(v) {
                intersectedSet.Add(v)
            }
        }
    }
    return intersectedSet
}

// 生成集合 self 对集合 other 的差集
func Difference(self SetIntfs, other SetIntfs) SetIntfs {
    if self == nil || other == nil {
        return nil
    }
    differencedSet := NewSimpleSet()

    //if other.Len() == 0 {
    //    for _, v := range self.Elements() {
    //        differencedSet.Add(v)
    //    }
    //    return differencedSet
    //}
    for _, v := range self.Elements() {
        if !other.Contains(v) {
            differencedSet.Add(v)
        }
    }
    return differencedSet
}

// 生成集合 self 和集合 other 的对称差集
func SymmetricDifference(self SetIntfs, other SetIntfs) SetIntfs {
    if self == nil || other == nil {
        return nil
    }
    diffA := Difference(self, other)
    if other.Len() == 0 {
        return diffA
    }
    diffB := Difference(other, self)
    return Union(diffA, diffB)
}

//TODO 这个地方应该做成可以生成多种Set版本的功能
func NewSimpleSet() SetIntfs {
    return NewHashSet()
}

//判断value是否实现了Set接口
func IsSet(value interface{}) bool {
    if _, ok := value.(SetIntfs); ok {
        return true
    }
    return false
}
