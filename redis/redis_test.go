package redis

import "testing"

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
}