package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go addCount()
	go addCount()

	wg.Wait()
	fmt.Printf("count %d", count)

}

func addCount() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		if i == 500 {
			// 当前的gorotine从线程退出，并放回本地运行队列
			runtime.Gosched()
		}
		count++
	}
}
