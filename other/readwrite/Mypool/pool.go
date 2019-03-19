package Mypool

import (
	"sync"
	"io"
	"errors"
)

type pool struct {
	lock    sync.Mutex
	source  chan io.Closer
	factory func() io.Closer
	isClose bool
}

//正数用uint
func NewPool(size uint, f func() io.Closer) *pool {
	return &pool{
		source:  make(chan io.Closer, size),
		factory: f,
		isClose: false,
	}
}

var errPoolClosed = errors.New("池关闭")

func (p *pool) Acquire() (io.Closer, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.isClose {
		return nil, errPoolClosed
	}
	select {
	case s := <-p.source:
		return s, nil
	default:
		return p.factory(), nil
	}

}
func (p *pool) Release(s io.Closer) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.isClose {
		return errPoolClosed
	}
	select {
	case p.source <- s:
		return nil
	default:
		s.Close()
		return nil
	}

}
func (p *pool) close() {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.isClose {
		return
	} //通道重复关闭
	p.isClose = true
	close(p.source)
	for r := range p.source {
		r.Close()
	}
}
