package redis

import (
	"github.com/mediocregopher/radix.v2/pool"
	"fmt"
	"time"
)

var p *pool.Pool

//连接池
func InitRedisPool(address string) error{
	var err error
	p, err = pool.New("tcp", address, 5)
	if err != nil {
		return err
	}

	//启动协程保证Redis连接不超时
	go func() {
		for {
			if err := p.Cmd("PING").Err; err !=nil {
				fmt.Println("PING is err:", err.Error())
			}else {
				fmt.Println("PING is OK")
			}
			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}

func ExpireAt(key string, at int64) error{
	//p.Cmd简化写法
	resp :=  p.Cmd("EXPIREAT", key, at)
	if resp.Err != nil {
		return resp.Err
	}else{
		return nil
	}
}
//KV基本操作
func Set(key, val string) error{
	conn, err := p.Get()
	if err != nil {
		return err
	}
	defer p.Put(conn)

	if conn.Cmd("SET", key, val).Err != nil {
		return err
	}
	return nil
}

func SetEx(key, val string, timeout int) error{
	conn, err := p.Get()
	if err != nil {
		return err
	}
	defer p.Put(conn)

	if conn.Cmd("SET", key, val, "EX", timeout).Err != nil {
		return err
	}
	return nil
}

func Get(key string) (string,error){
	conn, err := p.Get()
	if err != nil {
		return "", err
	}
	defer p.Put(conn)

	val,err := p.Cmd("GET", key).Str()
	if err != nil {
		return "", err
	}

	return val, nil
}

//List基本操作
func Lpush(key, val string) error{
	conn, err := p.Get()
	if err != nil {
		return err
	}
	defer p.Put(conn)

	if conn.Cmd("LPUSH", key, val).Err != nil {
		return err
	}
	return nil
}

func Rpush(key, val string) error{
	conn, err := p.Get()
	if err != nil {
		return err
	}
	defer p.Put(conn)

	if conn.Cmd("RPUSH", key, val).Err != nil {
		return err
	}
	return nil
}

func Rpop(key string) (string,error) {
	//p.Cmd简化写法
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
	//p.Cmd简化写法
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
	//p.Cmd简化写法
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
	//p.Cmd简化写法
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