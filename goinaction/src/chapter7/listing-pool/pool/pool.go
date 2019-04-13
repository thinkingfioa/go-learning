package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m sync.Mutex

	//有缓冲的通道资源池
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value is wrong")
	}

	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Printf("Acquire:", "New Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Printf("Release:", "In Queue")
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// 在清空通道里的资源之前，将通道关闭，如果不这样做，会发生死锁
	close(p.resources)

	for r := range p.resources {
		r.Close()
	}

}
