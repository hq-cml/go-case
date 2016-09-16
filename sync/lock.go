package main

import (
    "os"
    "sync"
    "errors"
)
/*
 * 锁的使用：互斥锁 & 读写锁
 */

//基于锁的DataFile实现
type lockDataFile struct {
    f           *os.File     //文件句柄
    mutex_rw    sync.RWMutex //读写锁，用于文件读写
    woffset     int64        //写操作的偏移量
    roffset     int64        //读操作偏移量
    wmutex      sync.Mutex   //互斥锁，用于写操作
    rmutex      sync.Mutex   //互斥锁，用于读操作
    data_len    uint32       //数据块长度
}

//*lockDataFile 实现DataFileIntfs
//读取一个数据块
func (df *lockDataFile)Read() (rsn int64, d Data, err error) {
    
}

//写入一个数据块
func (df *lockDataFile)Write(d Data) (wsn int64, err error) {

}

//获取最后读取的数据快序列号
func (df *lockDataFile)Rsn() int64 {

}

//获取最后写入的数据快序列号
func (df *lockDataFile)Wsn() int64 {

}

//获取数据块的长度
func (df *lockDataFile)DataLen() uint32 {

}

//惯例New，通常返回值是某种接口的实现 + error的实现
func NewLockDataFile(path string, data_len uint32) (DataFileIntfs, error) {
    if data_len == 0 {
        return nil, errors.New("Invalid data length!")
    }

    f, err := os.Create(path)
    if err != nil {
        return nil, err
    }

    df := &lockDataFile{
        f        : f,
        data_len : data_len,
    }

    return df, nil
}


