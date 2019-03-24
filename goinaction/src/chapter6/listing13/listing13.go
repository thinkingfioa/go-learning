package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	count int64
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)

	go addCount(10000)
	go addCount(10000)

	wg.Wait()

	fmt.Printf("count %d", count)
}

func addCount(loopNum int) {
	defer wg.Done()

	for i := 0; i < loopNum; i++ {
		//count++
		atomic.AddInt64(&count, 1)
	}
}
