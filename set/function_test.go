package set
/*
 * 功能测试
 */
import (
    "testing"
    "runtime/debug"
    //"fmt"
    //"strings"
    myrandom "github.com/hq-cml/go-case/random"
    "strings"
    "fmt"
)

/**************************** HashSet相关功能 **************************/
//必须大写Test开头，测试创建
func TestHashSetCreation(t *testing.T) {
    defer func() {
        if err := recover(); err != nil {
            debug.PrintStack()
            t.Errorf("Fatal Error: %s\n", err)
        }
    }()
    t.Log("Starting TestHashSetCreation...")
    hs := NewHashSet()
    t.Logf("Create a HashSet value: %v\n", hs)
    if hs == nil {
        t.Errorf("The result of func NewHashSet is nil!\n")
    }
    isSet := IsSet(hs)//判断是否实现了Set接口
    if !isSet {
        t.Errorf("The value of HashSet is not Set!\n")
    } else {
        t.Logf("The HashSet value is aSet.\n")
    }
}

//测试Set的操作
func TestSetOperation(t *testing.T) {
    defer func() {
        if err := recover(); err != nil {
            debug.PrintStack()
            t.Errorf("Fatal Error: %s\n", err)
        }
    }()

    t.Logf("Starting TestHashSetOp...")
    hs := NewHashSet()
    if hs.Len() != 0 {
        t.Errorf("ERROR: The length of original HashSet value is not 0!\n")
        t.FailNow()
    }
    var randElem interface{}
    expectedElemMap := make(map[interface{}]bool)

    //添加8个元素
    var result bool
    for i := 0; i < 8; i++ {
        randElem = genRandElement()
        t.Logf("Add %v to the HashSet value %v.\n", randElem, hs)
        result = hs.Add(randElem)
        if expectedElemMap[randElem] && result {
            t.Errorf("ERROR: The element adding (%v => %v) is successful but should be failing!\n",
                randElem, hs)
            t.FailNow()
        }
        if !expectedElemMap[randElem] && !result {
            t.Errorf("ERROR: The element adding (%v => %v) is failing!\n",
                randElem, hs)
            t.FailNow()
        }
        expectedElemMap[randElem] = true
    }
    expectedLen := len(expectedElemMap)
    if hs.Len() != expectedLen {
        t.Errorf("ERROR: The length of HashSet value %d is not %d!\n", hs.Len(), expectedLen)
        t.FailNow()
    }

    //挨个检查添加的元素是否存在
    for k, _ := range expectedElemMap {
        if !hs.Contains(k) {
            t.Errorf("ERROR: The HashSet value %v do not contains %v!", hs, k)
            t.FailNow()
        }
    }

    //测试删除
    number := 2
    for k, _ := range expectedElemMap {
        if number%2 == 0 {
            t.Logf("Remove %v from the HashSet value %v.\n", k, hs)
            hs.Remove(k)
            if hs.Contains(k) {
                t.Errorf("ERROR: The element adding (%v => %v) is failing!\n",
                    randElem, hs)
                t.FailNow()
            }
            delete(expectedElemMap, k)
        }
        number++
    }

    expectedLen = len(expectedElemMap)
    if hs.Len() != expectedLen {
        t.Errorf("ERROR: The length of HashSet value %d is not %d!\n", hs.Len(), expectedLen)
        t.FailNow()
    }
    t.Logf("After remove. HashSet value %v.\n", hs)

    //测试是否相同
    hs2 := NewHashSet()
    for k, _ := range expectedElemMap {
        hs2.Add(k)
    }
    if !hs.Same(hs2) {
        t.Errorf("ERROR: HashSet value %v do not same %v!\n", hs, hs2)
        t.FailNow()
    }

    //测试字符串化
    str := hs.String()
    t.Logf("The string of HashSet value %v is '%s'.\n", hs, str)
    for _, v := range hs.Elements() {
        if !strings.Contains(str, fmt.Sprintf("%v", v)) {
            t.Errorf("ERROR: '%s' do not contains '%v'!", str, v)
            t.FailNow()
        }
    }
}

// ----- 随机测试对象生成函数 -----
func genRandSet(newSet func() SetIntfs) (set SetIntfs, elemMap map[interface{}]bool) {
    set = newSet()
    elemMap = make(map[interface{}]bool)
    var enough bool
    for !enough {
        e := genRandElement()
        set.Add(e)
        elemMap[e] = true
        if len(elemMap) >= 3 {
            enough = true
        }
    }
    return
}

func genRandElement() interface{} {
    i := myrandom.GenRandIntMinMax(0, 4)
    switch i {
    case 0:
        return myrandom.GenRandInt(10000)
    case 1:
        return myrandom.GenRandStringMaxLen(15)
    case 2:
        return struct {
            num int64
            str string
        }{myrandom.GenRandInt(10000), myrandom.GenRandStringMaxLen(15)}
    default:
        const length = 2
        arr := new([length]interface{})
        for i := 0; i < length; i++ {
            if i%2 == 0 {
                arr[i] = myrandom.GenRandInt(10000)
            } else {
                arr[i] = myrandom.GenRandStringMaxLen(15)
            }
        }
        return *arr
    }
}

