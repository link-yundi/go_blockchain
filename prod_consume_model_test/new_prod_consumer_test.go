package prod_consume_model

// 1对多 基于 RWMutex的实现
import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	shareData int // 共享数据
	wg        sync.WaitGroup
	lock      sync.RWMutex
)

//func readData() int {
//	// 读取共享数据
//	lock.RLock()
//	defer lock.RUnlock()
//	ret := shareData
//	return ret
//}
//
//func writeData(data int) {
//	// 写入共享数据
//	lock.Lock()
//	defer lock.Unlock()
//	shareData = data
//}

func newProducer() {
	for i := 0; i < 10; i++ {
		lock.Lock()
		fmt.Printf("生产者，生产 %d \n", i)
		shareData = i
		lock.Unlock()
		time.Sleep(300 * time.Millisecond)
	}
	wg.Done()
}

func newConsumer(idx int) {
	for {
		lock.RLock()
		num := shareData
		fmt.Printf("%dth 消费者，消费 %d \n", idx, num)
		lock.RUnlock()
	}
	wg.Done()
}

func TestModel(t *testing.T) {
	wg.Add(1)
	go newProducer()
	// 开3个消费者
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go newConsumer(i + 1)
	}
	wg.Wait()
}
