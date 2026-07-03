package mutexbench

import "sync"

type CounterMutex struct {
	mu      sync.Mutex
	counter int
}

func (c *CounterMutex) Init() {
	c.counter = 0
}

func (c *CounterMutex) Increment() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *CounterMutex) Get() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return int64(c.counter)
}
