package redis

import (
	"testing"
	"time"
	"runtime"
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
	t.Log("Begin")
	err := InitRedisPool("127.0.0.1:6379")
	if err != nil {
		t.Log("Init redis pool error:", err.Error())
	}

	v,err := Get("aaa");
	if err !=nil {
		t.Log("Error:", err.Error())
	}
	t.Log("Get aaa:", v)

	//redis-server 的超时设置成了10s
	time.Sleep(12 * time.Second)

	//这时候读取会失败
	v,err = Get("aaa");
	if err !=nil {
		t.Log("Get aaa Error:", err.Error())

		v,err = Get("aaa")
		if err !=nil {
			t.Log("Try again. Get aaa error:", err.Error())
		}
		t.Log("Try again. Get aaa :", v)
	} else {
		t.Log("Get aaa:", v)
	}

	t.Log("Get aaa:", v)
	time.Sleep(5 * time.Second)
}

//测试并发协程下的重连
func TestMultiGoroutineConnTimeout(t *testing.T) {
	//这句话必须要加的，否则没法测出并发的效果
	runtime.GOMAXPROCS(4)
	t.Log("Begin")
	err := InitRedisPool("127.0.0.1:6379")
	if err != nil {
		t.Log("Init redis pool error:", err.Error())
	}

	go func() {
		v,err := Get("bbb");
		if err !=nil {
			t.Log("Error:", err.Error())
		}
		t.Log("Get bbb:", v)

		//redis-server 的超时设置成了10s
		time.Sleep(12 * time.Second)

		//这时候读取会失败
		v,err = Get("bbb");
		if err !=nil {
			t.Log("Get bbb Error:", err.Error())

			v,err = Get("bbb")
			if err !=nil {
				t.Log("Try again. Get bbb error:", err.Error())
			}
			t.Log("Try again. Get bbb :", v)
		} else {
			t.Log("Get bbb:", v)
		}

		t.Log("Get bbb:", v)
		time.Sleep(5 * time.Second)
	}()

	v,err := Get("ccc");
	if err !=nil {
		t.Log("Error:", err.Error())
	}
	t.Log("Get ccc:", v)

	//redis-server 的超时设置成了10s
	time.Sleep(12 * time.Second)

	//这时候读取会失败
	v,err = Get("ccc");
	if err !=nil {
		t.Log("Get aaa Error:", err.Error())

		v,err = Get("aaa")
		if err !=nil {
			t.Log("Try again. Get ccc error:", err.Error())
		}
		t.Log("Try again. Get ccc :", v)
	} else {
		t.Log("Get ccc:", v)
	}

	t.Log("Get ccc:", v)
	time.Sleep(5 * time.Second)
}