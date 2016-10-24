利用golang内置的map实现一个set功能。

set.go定义了一个Set的通用行为，即一个Set接口。

HashSet类型实现了Set接口，一个利用hash实现的Set。用key存储set的值key用interface{}类型，表示set的值是任意类型
val用bool类型（因为1，省空间 2，用起来方便，可以表示是否存在等）。

此外，set.go中还定义了一系列公用函数，他们作为高级功能，没必要每种实现各自重复实现，而是应该基础方法组装而成。

