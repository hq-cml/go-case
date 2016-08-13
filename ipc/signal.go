package main
/*
 * 信号的处理
 */
import (
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"
	"os/signal"
)

func main() {
	sigHandleDemo()
}

func sigHandleDemo() {
	//创建信号接收channel1
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("Set notification for %s... [sigRecv1]\n", sigs1)
	signal.Notify(sigRecv1, sigs1...)

	//创建信号接收channel2
	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("Set notification for %s... [sigRecv2]\n", sigs2)
	signal.Notify(sigRecv2, sigs2...)

	//wg用于同步，Add(2)，然后每次Done会减1，知道为0的时候，Wait退出
	var wg sync.WaitGroup
	wg.Add(2)

	//并发协程用于捕获接收的信号
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("Received a signal from sigRecv1: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv1]\n")
		wg.Done()
	}()
	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal from sigRecv2: %s\n", sig)
		}
		fmt.Printf("End. [sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("Wait for 20 seconds... ")
	time.Sleep(20 * time.Second)
	fmt.Printf("Stop notification...")

	//取消sigRecv1的作用
	signal.Stop(sigRecv1)
	//关闭sigRecv1，协程一方可退出(close之后，for才会自动停止)
	close(sigRecv1)
	fmt.Printf("done. [sigRecv1]\n")

	//等待子协程退出
	wg.Wait()
}
