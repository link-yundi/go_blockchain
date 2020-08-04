package rwlock

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var (
	wg1     sync.WaitGroup
	rwmutex sync.RWMutex
	//ch      = make(chan int)
	shareData int
)

func writeGo(idx int) {
	// 生成随机数
	for {
		rwmutex.Lock()
		shareData = rand.Intn(1000)
		//ch <- num
		fmt.Printf("%dth 写go程，写入%d\n", idx, shareData)
		time.Sleep(300 * time.Millisecond)
		rwmutex.Unlock()
	}
	wg1.Done()
}

func readGo(idx int) {
	for {
		rwmutex.RLock()
		//num := <-ch
		fmt.Printf("%dth -------读go程，读出%d\n", idx, shareData)
		rwmutex.RUnlock()
	}
	wg1.Done()
}

func TestRWMutex(t *testing.T) {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		wg1.Add(1)
		go readGo(i)
	}
	for i := 0; i < 5; i++ {
		wg1.Add(1)
		go writeGo(i)
	}
	wg1.Wait()
}
