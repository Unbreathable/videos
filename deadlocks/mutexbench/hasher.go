package mutexbench

import (
	"crypto/sha512"
	"fmt"
	"sync"
)

type HasherMutex struct {
	mu    sync.Mutex
	count int64
	N     int
}

func (s *HasherMutex) Init() {}

func (s *HasherMutex) Run() {
	s.mu.Lock()
	doWork(s.N)
	s.count++
	s.mu.Unlock()
}

func (s *HasherMutex) Get() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

type HasherWorker struct {
	count int64
	ch    chan struct{}
	N     int
}

func (s *HasherWorker) Init() {
	s.ch = make(chan struct{})
	go func() {
		for range s.ch {
			doWork(s.N)
			s.count++
		}
	}()
}

func (s *HasherWorker) Run() {
	s.ch <- struct{}{}
}

func (s *HasherWorker) Get() int64 {
	return s.count
}

// Calculate n sha512 hashes to simulate work
func doWork(n int) {
	for i := range n {
		sha512.Sum512(fmt.Appendf(nil, "work %d", i))
	}
}
