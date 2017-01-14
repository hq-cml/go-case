package redis

import (
	"testing"
	"time"
	"fmt"
)

//测试基本的key、val方法
func TestSetGet(t *testing.T) {
	err := InitRedisPool("127.0.0.1:6379")
	if err != nil {
		t.Fatal("Init redis pool error:", err.Error())
	}
	err = SetEx("hq1", "AAA",  3600)
	if err != nil {
		t.Fatal("SetEx hq1 error:", err.Error())
	}
	err = Set("hq2", "BBB")
	if err != nil {
		t.Fatal("Set hq2 error:", err.Error())
	}
	a,err := Get("hq1")
	if err != nil {
		t.Fatal("Get hq1 error:", err.Error())
	}
	b,err := Get("hq2")
	if err != nil {
		t.Fatal("Get hq2 error:", err.Error())
	}

	t.Log(a)
	t.Log(b)
}

//测试List的方法
func TestList(t *testing.T) {
	err := InitRedisPool("127.0.0.1:6379")
	if err != nil {
		t.Fatal("Init redis pool error:", err.Error())
	}
	err = Lpush("hq_test", "aaa1")
	if err != nil {
		t.Fatal("Lpush hq_test error:", err.Error())
	}
	err = Lpush("hq_test", "aaa2")
	if err != nil {
		t.Fatal("Lpush hq_test error:", err.Error())
	}
	err = Rpush("hq_test", "aaa3")
	if err != nil {
		t.Fatal("Lpush hq_test error:", err.Error())
	}
	lst,err := Lrange("hq_test", 0, -1)
	if err != nil {
		t.Fatal("Rpop hq_test error:", err.Error())
	}
	t.Log("The list is:", lst)
	l,err := Llen("hq_test")
	if err != nil {
		t.Fatal("Llen hq_test error:", err.Error())
	}
	t.Log("The len is:", l)
	val, err := Rpop("hq_test")
	if err != nil {
		t.Fatal("Rpop hq_test error:", err.Error())
	}
	t.Log("Rpop list is:", val)
	val, err = Lpop("hq_test")
	if err != nil {
		t.Fatal("Rpop hq_test error:", err.Error())
	}
	t.Log("Lpop list is:", val)
	lst,err = Lrange("hq_test", 0, -1)
	if err != nil {
		t.Fatal("Rpop hq_test error:", err.Error())
	}
	t.Log("The list is:", lst)

	l,err = Llen("hq_test")
	if err != nil {
		t.Fatal("Llen hq_test error:", err.Error())
	}
	t.Log("The len is:", l)
}

//测试基本的key、val方法
func TestExpire(t *testing.T) {
	err := InitRedisPool("10.94.112.246:6379")
	if err != nil {
		t.Fatal("Init redis pool error:", err.Error())
	}
	err = Set("hq1", "AAA")
	if err != nil {
		t.Fatal("Set hq1 error:", err.Error())
	}

	time := time.Now().Unix()
	t.Log("Current time:", time)

	err = ExpireAt("hq1", time + 1800)
	if err != nil {
		t.Fatal("Expire error:", err.Error())
	}
}

func TestConnTimeout(t *testing.T) {
	fmt.Println("AAAAAAAAA")
	err := InitRedisPool("127.0.0.1:6379")
	if err != nil {
		t.Fatal("Init redis pool error:", err.Error())
	}

	v,err := Get("aaa");
	if err !=nil {
		t.Fatal("Error:", err.Error())
	}

	t.Log("V is ", v)

	//redis-server 的超时设置成了10s
	time.Sleep(12 * time.Second)

	v,err = Get("a");
	if err !=nil {
		t.Log("Error:", err.Error())
	}

	t.Log("V is ", v)
	time.Sleep(5 * time.Second)
}
