package bpool

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	pool sync.Pool
	once sync.Once
}

func (p *BufferPool) init() {
	p.once.Do(func() {
		p.pool = sync.Pool{New: func() interface{} {
			return bytes.NewBuffer(nil)
		}}
	})
}

func (p *BufferPool) Get() *bytes.Buffer {
	p.init()
	return p.pool.Get().(*bytes.Buffer)
}

func (p *BufferPool) Put(b *bytes.Buffer) {
	p.init()
	b.Reset()
	p.pool.Put(b)
}

var bufferpool = BufferPool{}

func GetBuffer() *bytes.Buffer {
	return bufferpool.Get()
}

func PutBuffer(b *bytes.Buffer) {
	bufferpool.Put(b)
}