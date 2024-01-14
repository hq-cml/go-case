/*
 * 关于时间的用法，重点是时区
 * 补充：
 * 	 UTC -- 世界统一时间，约等于GMT（已废弃），简单理解就是0度经线时间
 *   CST -- 中国标准时间，东八区时间，北京时间，UTC + 8
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println(now, "创建的时间，是带了时区的，且是系统时区保持一致") // 2024-01-14 20:35:06.9360665 +0800 CST m=+0.009026401，默认的打印，带有时区信息
	fmt.Println("-----------------------------")

	// 格式化时间
	layout := "2006-01-02 15:04:05"
	fmt.Println(now.Format(layout))
	fmt.Println("-----------------------------")

	// 关于时区
	locUSA, err := time.LoadLocation("America/New_York") // 新建一个美国时区
	if err != nil {
		panic(err)
	}
	nowUSA := now.In(locUSA)
	fmt.Println(nowUSA, "此刻的美国时间（其实时间戳是相同的，同一时刻）") // 可以看到比北京时间晚了13小时
	fmt.Println(nowUSA.Format(layout))
	fmt.Println("-----------------------------")

	// 将字符串转换成时间对象
	str := "2024-01-10 18:00:00"
	t1, err := time.Parse(layout, str) // 注意：这里不带时区，转成的是UTC时间！
	if err != nil {
		panic(err)
	}
	fmt.Println(t1, "注意：time.Parse默认是转成UTC时间！")
	fmt.Println(t1.Format(layout), "格式化之后不显示时区了，更具有迷惑性")
	locBeijing, err := time.LoadLocation("Asia/Shanghai") // 创建北京时区
	if err != nil {
		panic(err)
	}
	fmt.Println(t1.In(locBeijing), "对应的北京时间（东八区，所以加8小时）")
	fmt.Println(t1.In(locBeijing).Format(layout))
	t1Usa := t1.In(locUSA) // 将UTC时间转成美国时间，美国慢了5小时
	fmt.Println(t1Usa, "此时美国时间，比UTC慢了5小时，比北京慢了13小时")
	// 时间戳是否一样？
	fmt.Println(t1.Unix(), ",", t1.In(locBeijing).Unix(), "可以看到，只要时间一样则时间戳一样，尽管他们根据时区格式化之后打印不一样")
	fmt.Println("-----------------------------")

	// 反向验证上面这个问题，同一时间戳，在不同时区的打印区别
	timestamp := 0 // 0时间戳，理论上表示：1970年1月1日 00:00:00 UTC，看看它格式化出来什么样子
	orgTime := time.Unix(int64(timestamp), 0)
	fmt.Println(orgTime, "可以看到0时间戳打印出来已经是8点了，原因是因为orgTime的时区是北京市区")
	fmt.Println(orgTime.In(time.UTC), "UTC时区是1970年1月1日 00:00:00")
	fmt.Println(orgTime.In(locUSA), "此时美国的时间，神奇到了1969年")
	fmt.Println("-----------------------------")

	// 因为time.Parse默认是UTC的时间，所以字符串解析的时候一定要小心，务必加上时区
	t2, err := time.ParseInLocation(layout, str, time.Local) // time.Local是使用了系统的时区，这里是东八区
	if err != nil {
		panic(err)
	}
	fmt.Println(t2, "这里终于变成了预期中的时间2024-01-10 18:00:00 CST（时间符合，时区也符合）")
	fmt.Println(t2.In(time.UTC), "此刻的UTC时间")
	fmt.Println(t2.In(locUSA), "美国时间")
	fmt.Println("-----------------------------")

	// 时间的运算
	fmt.Println(now.Add(1*time.Hour), "一小时后") //当前时间加1小时，如果减的话就用负号
	fmt.Println(now.AddDate(0, 1, 0), "一个月后") // 一个月后的时间，如果减就用负数
	delta := now.Sub(now.Add(1 * time.Hour))      // 时间相减，得到Duration
	fmt.Println(delta)                            // 打印
	fmt.Println(delta.String())
	fmt.Println(int64(delta))
	fmt.Println("-----------------------------")

	// 时间比较
	t3 := now.Add(1 * time.Hour)
	fmt.Println(now.Equal(t3))  // now == t3 ?
	fmt.Println(now.After(t3))  // now > t3 ?
	fmt.Println(now.Before(t3)) // now < t3 ?
}
