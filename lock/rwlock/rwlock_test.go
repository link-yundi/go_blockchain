package rwlock

import (
	"sync"
	"testing"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	rwLock sync.RWMutex
	lock   sync.Mutex
)

func write() {
	rwLock.Lock() // 上写锁
	//lock.Lock()
	x += 1
	time.Sleep(10 * time.Millisecond)
	//lock.Unlock()
	rwLock.Unlock()
}

func read() {
	rwLock.RLock() // 上写锁
	//x += 1
	//lock.Lock()
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
}

func Test(t *testing.T) {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
		wg.Done()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
		wg.Done()
	}
	wg.Wait()
	t.Log(time.Now().Sub(start))
}
