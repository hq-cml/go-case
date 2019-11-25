package main

import (
	"context"
	"fmt"
	"time"
)

type DocNode struct {
	Weight uint16
	DocId  uint32
}

//父亲
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	//儿子
	go func(ctx context.Context) {

		//ctxSon, _ := context.WithCancel(ctx)
		ctxSon := context.WithValue(ctx, "A", 1)

		//孙子
		go func (ctx context.Context) {
			fmt.Println("K--------", ctx.Value("A"))
			select {
			case <- ctx.Done():
				fmt.Println("Son exit!")
			}
		}(ctxSon)

		select {
		case <- ctx.Done():
			fmt.Println("Father exit!")
		}

	}(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}