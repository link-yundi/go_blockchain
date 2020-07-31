package mutex

import (
	"sync"
	"testing"
)

// 锁的学习
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func TestAdd(t *testing.T) {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	t.Log(x)
}
