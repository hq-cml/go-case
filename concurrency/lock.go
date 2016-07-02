package main
//golang更推荐用channel做同步，不推荐用这种方式
//但也保留了这种用法
import "fmt"
import "sync"
import "runtime"

var counter int = 0 //全局变量

func Count(id int, lock *sync.Mutex){
    lock.Lock()
    counter ++
    fmt.Println("goroutine", id, "get lock. counter:", counter)
    lock.Unlock()
}

func main() {
    lock := &sync.Mutex{}

    for i:=0; i<10; i++ {
        go Count(i, lock)
    }

    for {
        lock.Lock()
        c := counter
        lock.Unlock()
        //fmt.Println("A", c)
        //让出时间分片,其他goroutine将得到执行机会。如果将上下面的Print打开，会发现Gosched作用明显。
        //有点类似C里面的sleep让出CPU，但是比sleep高级，因为sleep会导致线程睡指定时间，有点浪费
        runtime.Gosched()
        //fmt.Println("B", c)
        if c >= 10 {
            break
        }
    }
}
