package main

import (
	"sync"
	"sync/atomic"
)

type state struct {
	mu    sync.Mutex
	count int64
}

func (s *state) setCount() {
	s.count += 1
}

func (s *state) setCountMutex() {
	s.mu.Lock()
	s.count += 1
	s.mu.Unlock()
	// or defer s.mu.Unlock(), this bascially waits until before the function returns
}

func (s *state) setCountAtomic() {
	atomic.AddInt64(&s.count, 1) // atomically adds 1 to the variable
}

func main() {

}
