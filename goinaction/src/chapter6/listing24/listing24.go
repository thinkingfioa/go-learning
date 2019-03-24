package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberWorker = 4
	taskNum      = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	wg.Add(numberWorker)

	tasks := make(chan string, taskNum)

	// 启动${numberWorker}个gorotine
	for gr := 1; gr <= numberWorker; gr++ {
		go worker(tasks, gr)
	}
	// 提交任务
	for post := 1; post <= taskNum; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}
	// 关闭通道
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, work int) {
	defer wg.Done()

	for {
		// 获取任务
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", work)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", work, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : Completed %s\n", work, task)
	}
}
