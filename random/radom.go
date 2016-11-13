package random

import (
    "math/rand"
    "time"
    "fmt"
)

func init() {
    //设置种子放在init中，全局一次即可
    seed := time.Now().UnixNano()
    rand.Seed(seed)
    fmt.Println("The rand seed:", seed)
}

//生成64位非负随机整型
func GenRandInt(max int64) int64{
    //rand.Seed(time.Now().UnixNano()) //根据random的原理，种子是不能放在这里的
    return rand.Int63n(max)
}
