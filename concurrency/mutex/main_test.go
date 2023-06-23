package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestSetCount(t *testing.T) {
	s := &state{}
	for i := 0; i < 10; i++ {
		s.setCount()
	} // runs synchronously
	fmt.Println(s.count)
}

func Test1SetCount(t *testing.T) {
	s := &state{}
	for i := 0; i < 10; i++ {
		go func() { // passing i into the go routine block
			s.setCount()
		}()
	}
	fmt.Println(s.count)
	// This test causes race condition error, which can be checked via go test ./... -v -race
	// to use uncached version of the test file, use -count=1 flag
}

// sync.WaitGroup
// this is like a counter that gets incremented whenever sync(obj).Add(1) is invoked and then when sync(obj).Done() is invoked the count gets decremented
func Test2SetCount(t *testing.T) {
	s := &state{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() { // passing i into the go routine block
			s.setCount()
			wg.Done()
		}()
	}
	wg.Wait() // waits for all the thread to be done
	fmt.Println(s.count)
}

// Mutex
func Test3SetCountMutex(t *testing.T) {
	s := &state{}
	for i := 0; i < 10; i++ {
		go func() {
			s.setCount()
		}()
	}
	fmt.Println(s.count)
}

func Test4SetCountAtomic(t *testing.T) {
	s := &state{}
	for i := 0; i < 10; i++ {
		go func() {
			s.setCount()
		}()
	}
	fmt.Println(s.count)
}
