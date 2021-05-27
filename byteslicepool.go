package bpool

import (
	"sync"
)

const DefaultByteSlicePoolCapacity = 8192

type ByteSlicePool struct {
	// Capacity if defined, new slices will be initialised with that capacity
	Capacity int

	pool sync.Pool
	once sync.Once
}

func (p *ByteSlicePool) init() {
	p.once.Do(func() {
		if p.Capacity == 0 {
			p.Capacity = DefaultByteSlicePoolCapacity
		}
		p.pool = sync.Pool{New: func() interface{} {
			return make([]byte, 0, p.Capacity)
		}}
	})
}

func (p *ByteSlicePool) Get() []byte {
	p.init()
	return p.pool.Get().([]byte)
}

func (p *ByteSlicePool) Put(b []byte) {
	p.init()
	p.pool.Put(b[:0])
}

var byteslicepool = ByteSlicePool{}

func GetByteSlice() []byte {
	return byteslicepool.Get()
}

func PutByteSlice(p []byte) {
	byteslicepool.Put(p)
}
