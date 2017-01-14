package redis

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	//"fmt"
	//"time"
	"fmt"
	"sync"
)

//连接池句柄，p.Cmd是简化先发，包含了p.Get和p.Put，具体可以参看radix.V2的源码
var p *pool.Pool
var addr string

//用connVersion和m实现一个简单的乐观锁，防止并发情况下出现网络异常的时候，多次重连
var connVersion int64
var m sync.Mutex

//check Redis操作中的错误，如果是网络错误，则会进行重连
func handleError(resp *redis.Resp){
	if resp.IsType(redis.IOErr) {
		fmt.Println("Aoh, Network error:", resp.Err.Error())
		//进行重连
		err := InitRedisPool(addr)
		if err != nil {
			fmt.Println("Reconnect error:", err.Error())
		}
	}
}

//连接池
func InitRedisPool(address string) error{
	var err error
	addr = address //将地址存储，便于后续重连

	var local_ver = connVersion
	m.Lock()
	defer m.Unlock()
	//如果local_ver != connVersion说明其他协程已经实现了重连
	if local_ver == connVersion {
		fmt.Println("Connect Redis. connVer:", connVersion)
		p, err = pool.New("tcp", address, 5)
		if err != nil {
			return err
		}
		connVersion ++
	}


	return nil
}

//KV基本操作
func Set(key, val string) error{
	resp :=  p.Cmd("SET", key, val)
	if resp.Err != nil {
		return resp.Err
	}else{
		return nil
	}
}

func ExpireAt(key string, at int64) error{
	resp :=  p.Cmd("EXPIREAT", key, at)
	if resp.Err != nil {
		return resp.Err
	}else{
		return nil
	}
}

func SetEx(key, val string, timeout int) error{
	resp :=  p.Cmd("SET", key, val, "EX", timeout)
	if resp.Err != nil {
		return resp.Err
	}else{
		return nil
	}
}

func Get(key string) (string,error){
	resp :=  p.Cmd("GET", key)
	if resp.Err != nil {
		handleError(resp)
		return "", resp.Err
	}else{
		val, err := resp.Str()
		if err != nil {
			return "", err
		}
		return val, nil
	}
}

//List基本操作
func Lpush(key, val string) error{
	resp :=  p.Cmd("LPUSH", key, val)
	if resp.Err != nil {
		return resp.Err
	}else{
		return nil
	}
}

func Rpush(key, val string) error{
	resp :=  p.Cmd("RPUSH", key, val)
	if resp.Err != nil {
		return resp.Err
	}else{
		return nil
	}
}

func Rpop(key string) (string,error) {
	resp :=  p.Cmd("RPOP", key)
	if resp.Err != nil {
		return "", resp.Err
	}else{
		val, err := resp.Str()
		if err != nil {
			return "", err
		}
		return val, nil
	}
}

func Lpop(key string) (string,error) {
	resp :=  p.Cmd("LPOP", key)
	if resp.Err != nil {
		return "", resp.Err
	}else{
		val, err := resp.Str()
		if err != nil {
			return "", err
		}
		return val, nil
	}
}

func Lrange(key string, start, length int) ([]string,error) {
	resp :=  p.Cmd("LRANGE", key, start, length)
	if resp.Err != nil {
		return nil, resp.Err
	}else{
		lst, err := resp.List()
		if err != nil {
			return nil, err
		}
		return lst, nil
	}
}

func Llen(key string) (int64,error) {
	resp :=  p.Cmd("LLEN", key)
	if resp.Err != nil {
		return 0, resp.Err
	}else{
		val, err := resp.Int64()
		if err != nil {
			return 0, err
		}
		return val, nil
	}
}