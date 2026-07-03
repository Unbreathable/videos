package mutexbench

import (
	"sync"
	"time"
)

type SleeperMutex struct {
	mu    sync.Mutex
	count int64
}

func (s *SleeperMutex) Init() {}

func (s *SleeperMutex) Run(d time.Duration) {
	s.mu.Lock()
	time.Sleep(d)
	s.count++
	s.mu.Unlock()
}

func (s *SleeperMutex) Get() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

type SleeperWorker struct {
	count int64
	ch    chan time.Duration
}

func (s *SleeperWorker) Init() {
	s.ch = make(chan time.Duration)
	go func() {
		for d := range s.ch {
			time.Sleep(d)
			s.count++
		}
	}()
}

func (s *SleeperWorker) Run(d time.Duration) {
	s.ch <- d
}

func (s *SleeperWorker) Get() int64 {
	return s.count
}
