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
