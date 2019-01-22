package common

import (
	"bytes"
	"math/rand"
	"strconv"
	"encoding/json"
	"sync"
	"time"
)

/*****************  模拟CPU负载 ***************/
func CPUProfile() error {
	max := 10000000
	var buf bytes.Buffer
	for j := 0; j < max; j++ {
		num := rand.Int63n(int64(max))
		str := strconv.FormatInt(num, 10)
		buf.WriteString(str)
	}
	_ = buf.String()
	return nil
}


/****************** 模拟内存 ********************/
// box 代表数据盒子。
type box struct {
	Str   string
	Code  rune
	Bytes []byte
}

func MemProfile() error {
	max := 50000
	var buf bytes.Buffer
	for j := 0; j < max; j++ {
		seed := rand.Intn(95) + 32
		one := createBox(seed)
		b, err := genJSON(one)
		if err != nil {
			return err
		}
		buf.Write(b)
		buf.WriteByte('\t')
	}
	_ = buf.String()
	return nil
}

func createBox(seed int) box {
	if seed <= 0 {
		seed = 1
	}
	var array []byte
	size := seed * 8
	for i := 0; i < size; i++ {
		array = append(array, byte(seed))
	}
	return box{
		Str:   string(seed),
		Code:  rune(seed),
		Bytes: array,
	}
}

func genJSON(one box) ([]byte, error) {
	return json.Marshal(one)
}

/****************** 模拟阻塞 ******************/
//Send多余recv， 所以send会阻塞
func BlockProfile() error {
	max := 100
	senderNum := max / 2
	receiverNum := max / 4
	ch1 := make(chan int, max/4)

	var sendGroup sync.WaitGroup
	sendGroup.Add(senderNum)
	repeat := 50000
	for j := 0; j < senderNum; j++ {
		go send(ch1, &sendGroup, repeat)
	}

	go func() {
		sendGroup.Wait()
		close(ch1)
	}()

	var recvGroup sync.WaitGroup
	recvGroup.Add(receiverNum)
	for j := 0; j < receiverNum; j++ {
		go recv(ch1, &recvGroup)
	}
	recvGroup.Wait()
	return nil
}

func send(ch1 chan int, wg *sync.WaitGroup, repeat int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 10)
	for k := 0; k < repeat; k++ {
		elem := rand.Intn(repeat)
		ch1 <- elem
	}
}

func recv(ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for elem := range ch1 {
		_ = elem
	}
}
