package random

import "testing"

func TestGenRandInt(t *testing.T) {
    var i int64
    i = GenRandInt(10000)
    t.Log(i)
    i = GenRandInt(10000)
    t.Log(i)
    i = GenRandInt(10000)
    t.Log(i)
    i = GenRandInt(10000)
    t.Log(i)
    i = GenRandInt(10000)
    t.Log(i)
    i = GenRandInt(10000)
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