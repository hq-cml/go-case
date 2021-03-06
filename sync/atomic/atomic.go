package atomic

import (
    "os"
    "sync"
    "errors"
    "io"
    mysync "github.com/hq-cml/go-case/sync"
    "sync/atomic"
)
/*
 * 比cond.go更加改进，去掉了互斥锁，只使用 读写锁 + 条件变量
 */

//基于锁的DataFile实现
type condDataFile struct {
    f           *os.File     //文件句柄
    f_rwmutex   sync.RWMutex //读写锁，用于保护文件f本身的操作
    w_offset    int64        //写操作的偏移量
    r_offset    int64        //读操作偏移量
    data_len    uint32       //数据块长度
    r_cond      *sync.Cond   //和f读操作互斥量绑定的条件变量
}

//*lockDataFile 实现DataFileIntfs
//读取一个数据块，返回rsn表示读取到的数据块的编号
func (df *condDataFile)Read() (rsn int64, d mysync.Data, err error) {
    //获取读取偏移量，利用原子操作替代锁
    var offset int64
    for {
        offset = atomic.LoadInt64(&df.r_offset) //原子读取
        //原子交换试探（乐观锁）
        if atomic.CompareAndSwapInt64(&df.r_offset, offset, (offset + int64(df.data_len))) {
            break
        }
    }

    //读取文件数据块
    rsn = offset / int64(df.data_len)
    bytes := make([]byte, df.data_len)

    //锁与解锁变得很清晰，改进了lock.go的版本
    df.f_rwmutex.RLock()
    defer df.f_rwmutex.RUnlock()
    for {
        //ReadAt reads len(b) bytes from the File starting at byte offset off.
        _, err = df.f.ReadAt(bytes, offset)
        if err != nil {
            if err == io.EOF {
                df.r_cond.Wait() //原子得陷入阻塞等待
                continue
            }
            return
        }
        d = bytes
        return
    }
}

//写入一个数据块
func (df *condDataFile)Write(d mysync.Data) (wsn int64, err error) {
    //获取写偏移量，利用原子操作替代锁
    var offset int64
    for {
        offset = atomic.LoadInt64(&df.w_offset) //原子读取
        //原子交换试探（乐观锁）
        if atomic.CompareAndSwapInt64(&df.w_offset, offset, (offset + int64(df.data_len))) {
            break
        }
    }

    //写入一个数据块
    wsn = offset / int64(df.data_len)
    var bytes []byte
    if len(d) > int(df.data_len) {
        //如果数据超长，需要截断
        bytes = d[0:df.data_len]
    } else {
        bytes = d
    }
    df.f_rwmutex.Lock()
    defer df.f_rwmutex.Unlock()
    _, err = df.Write(bytes)
    df.r_cond.Signal()
    return
}

//获取最后读取的数据快序列号
func (df *condDataFile)Rsn() int64 {
    offset := atomic.LoadInt64(&df.r_offset)
    return offset / int64(df.data_len)
}

//获取最后写入的数据快序列号
func (df *condDataFile)Wsn() int64 {
    offset := atomic.LoadInt64(&df.w_offset)
    return offset / int64(df.data_len)
}

//获取数据块的长度
func (df *condDataFile)DataLen() uint32 {
    return df.data_len
}

//惯例New，通常返回值是某种接口的实现 + error的实现
func NewLockDataFile(path string, data_len uint32) (mysync.DataFileIntfs, error) {
    if data_len == 0 {
        return nil, errors.New("Invalid data length!")
    }

    f, err := os.Create(path)
    if err != nil {
        return nil, err
    }

    df := &condDataFile{
        f        : f,
        data_len : data_len,
    }

    //与f的读操作的互斥量绑定！
    df.r_cond = sync.NewCond(df.f_rwmutex.RLocker())

    return df, nil
}


