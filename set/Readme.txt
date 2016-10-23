利用golang内置的map实现一个set功能。

HashSet类型实现了Set接口，即利用hash实现的Set类型。
用key存储set的值key用interface{}类型，表示set的值是任意类型
val用bool类型，因为1，省空间 2，用起来方便，可以表示是否存在等

