package set

/*
 * 利用内置的hash实现set功能，用key存储set的值
 * key用interface{}类型，表示set的值是任意类型
 * val用bool类型，因为1，省空间 2，用起来方便，可以表示是否存在等
 */
type HashSet struct{
    m map[interface{}]bool
}

//“构造”函数
func NewHashSet() *HashSet{
    return &HashSet{m: make(map[interface{}]bool)}
}

