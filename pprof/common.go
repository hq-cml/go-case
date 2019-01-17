package main

import (
	"os"
	"path/filepath"
	"errors"
	"fmt"
	"time"
)

// CreateFile 用于在当前目录下创建一个指定名称的文件。
// 若同名文件已存在，则清空并复用。
func CreateFile(dir, name string) (*os.File, error) {
	if dir == "" {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}
	path := filepath.Join(dir, name)
	return os.Create(path)
}

// OpFunc 代表包含高负载操作的函数。
type OpFunc func() error

// Execute 用于执行指定次数的高负载的函数。
func Execute(op OpFunc, times int) (err error) {
	if op == nil {
		return errors.New("nil operation function")
	}
	if times <= 0 {
		return fmt.Errorf("invalid times: %d", times)
	}
	var t1 time.Time
	defer func() {
		diff := time.Now().Sub(t1)
		fmt.Printf("(elapsed time: %s)\n", diff)
		if p := recover(); p != nil {
			err = fmt.Errorf("fatal error: %v", p)
		}
	}()
	t1 = time.Now()
	for i := 0; i < times; i++ {
		if err = op(); err != nil {
			return
		}
		time.Sleep(time.Microsecond)
	}
	return
}