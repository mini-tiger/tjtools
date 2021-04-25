package main

import (
	"fmt"
	"sync"
)

// 内置 并发安全map
func rangefn(key, value interface{}) bool {
	fmt.Printf("key:%v,value:%v\n", key, value)
	defer sw.Done()
	return true
}

var sw sync.WaitGroup

func main() {

	m := new(sync.Map)

	go func() {
		m.Range(rangefn)
	}()

	for i := 0; i < 100; i++ {
		m.Store(i, i)
		sw.Add(1)
	}

	//time.Sleep(5*time.Second)
	sw.Wait()
}
