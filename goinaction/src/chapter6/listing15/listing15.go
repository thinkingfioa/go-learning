package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(10 * time.Second)
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
	fmt.Println("Finish")
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("output %s\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("%s shutdown.\n", name)
			break
		}
	}
}
