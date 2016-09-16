package main

/*
 * golang的同步机制：
 * 一个支持并发的文件，用于存放数据。
 * 每次写操作会追加固定长度的数据到文件中，并发goroutine之间写的内容不能出现穿插，即写时原子的。
 * 每次读取操作必须是固定长度完整的数据块，并且并发goroutine读取的数据应该不能重复，顺序进行。
 */

//数据类型
type Data []byte

//文件接口
type DataFileIntfs interface {
    Read() (rsn int64, d Data, err error)   //读取一个数据块
    Write(d Data) (wsn int64, err error)    //写入一个数据块
    Rsn() int64                             //获取最后读取的数据快序列号
    Wsn() int64                             //获取最后写入的数据快序列号
    DataLen() uint32                        //获取数据块的长度
}