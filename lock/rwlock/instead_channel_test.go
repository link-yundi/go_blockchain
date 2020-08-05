package rwlock

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 通过channel来实现读写锁同样的效果
var (
	wg2 sync.WaitGroup
	ch  = make(chan int)
)

func TestChannelSync(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go writeGo2(i)
	}
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go readGo2(i)
	}
	wg2.Wait()
}

func readGo2(idx int) {
	for {
		num := <-ch
		fmt.Printf("-------%dth go程 读 %d\n", idx, num)
	}
	wg2.Done()
}

func writeGo2(idx int) {
	for {
		num := rand.Intn(1000)
		ch <- num
		fmt.Printf("%dth go程 写 %d\n", idx, num)
		time.Sleep(300 * time.Millisecond)
	}
	wg2.Done()
}
