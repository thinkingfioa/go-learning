package main

import (
	"fmt"
	"sync"
)

var (
	count int64
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {

	wg.Add(2)

	go addCount(10000)
	go addCount(10000)

	wg.Wait()
	fmt.Printf("count %d\n", count)
}

func addCount(loopNum int) {
	defer wg.Done()

	for i := 0; i < loopNum; i++ {
		mutex.Lock()
		count++
		mutex.Unlock()
	}
}
