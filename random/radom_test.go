package random

import "testing"

func TestGenRandInt(t *testing.T) {
    var i int64
    i = GenRandInt(1)
    t.Log(i)
    i = GenRandInt(100)
    t.Log(i)
    i = GenRandInt(1000)
    t.Log(i)
    i = GenRandInt(100000)
    t.Log(i)
    i = GenRandInt(2)
    t.Log(i)
    i = GenRandInt(10000)
    t.Log(i)
}

func TestGenRandIntMinMax(t *testing.T) {
    var i int64
    i = GenRandIntMinMax(0, 100)
    t.Log(i)
    i = GenRandIntMinMax(0, 100)
    t.Log(i)
    i = GenRandIntMinMax(0, 100)
    t.Log(i)
    i = GenRandIntMinMax(99, 100)
    t.Log(i)
    i = GenRandIntMinMax(9900, 10000)
    t.Log(i)
}

func TestGenRandAscII(t *testing.T) {
    var c byte
    c = GenRandAscII()
    t.Logf("%c", c)
    c = GenRandAscII()
    t.Logf("%c", c)
    c = GenRandAscII()
    t.Logf("%c", c)
    c = GenRandAscII()
    t.Logf("%c", c)
    c = GenRandAscII()
    t.Logf("%c", c)
    c = GenRandAscII()
    t.Logf("%c", c)
}

//func TestGenRandString(t *testing.T) {
//    var s string
//    s = GenRandString()
//    t.Logf("%s", s)
//    s = GenRandString()
//    t.Logf("%s", s)
//    s = GenRandString()
//    t.Logf("%s", s)
//    s = GenRandString()
//    t.Logf("%s", s)
//    s = GenRandString()
//    t.Logf("%s", s)
//    s = GenRandString()
//    t.Logf("%s", s)
//    s = GenRandString()
//    t.Logf("%s", s)
//
//}