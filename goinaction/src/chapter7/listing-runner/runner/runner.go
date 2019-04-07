package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// 接收操作系统发送的中断信号
	interrupt chan os.Signal

	complete chan error

	// 接收超时事件
	timeout <-chan time.Time

	// 执行的函数，必须是一个接收整数且什么都不返回的函数
	tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		// interrupt通道的缓冲区容量初始化为1，确保语言运行是发送这个事件的时候不会阻塞
		interrupt: make(chan os.Signal, 1),
		// 当任务完成或退出后，返回一个error或者nil值，将等待main函数区接收这个值
		complete: make(chan error),
		// 在指定duration时间后，向通道发送一个time.Time的值
		timeout: time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	//
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		// 当发生中断事件信号时，停止接收后续信号
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
