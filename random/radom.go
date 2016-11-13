package random

import (
    "math/rand"
    "time"
    "fmt"
    //"bytes"
)

//！！种子很重要！！
func init() {
    //设置种子放在init中，全局一次即可
    s := time.Now().UnixNano()
    rand.Seed(s)
    fmt.Println("The rand seed:", s)
}

//生成64位非负随机整型: [0, max)
func GenRandInt(max int64) int64{
    //rand.Seed(time.Now().UnixNano()) //根据random的原理，种子是不能放在这里的
    return rand.Int63n(max)
}

//生成64位非负随机整型: [min, max)
func GenRandIntMinMax(min, max int64) int64{
    if min >= max {
        return -1
    }
    return min + rand.Int63n(max-min)
}

//生成随机字母 a-z, A-Z
func GenRandAscII() byte {
    min_big := 65 // A
    max_big := 90 // Z

    min_small := 97 // a
    max_small := 122 // z

    var c int
    switch rand.Intn(100) % 2 {
    case 0:
        c = min_big + rand.Intn(max_big-min_big)
    case 1:
        c = min_small + rand.Intn(max_small-min_small)
    }

    return byte(c)
}
