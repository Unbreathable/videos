package mutexbench

import "sync/atomic"

type CounterAtomic struct {
	counter atomic.Int64
}

func (c *CounterAtomic) Init() {
	c.counter = atomic.Int64{}
	c.counter.Store(0)
}

func (c *CounterAtomic) Increment() {
	c.counter.Add(1)
}

func (c *CounterAtomic) Get() int64 {
	return c.counter.Load()
}
