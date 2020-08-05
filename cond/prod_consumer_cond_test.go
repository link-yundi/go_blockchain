package cond

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 通过条件变量实现生产者消费者模型
var (
	cond sync.Cond
	wg   sync.WaitGroup
)

// 生产者
func producer(in chan<- int, idx int) {
	for {
		cond.L.Lock()      // 条件变量对应互斥锁加锁
		for len(in) == 3 { // 产品区满，等待消费者消费
			/*
				采用for循环而不是if判断
				这是为了保险起见，如果一个goroutine因为受到通知被唤醒，但是却发现共享资源的状态依旧不符合他的要求
				那么就应该继续掉用wait方法，等待下一次的通知唤醒
				假设wait唤醒后，没有进行共享资源的判断（别的goroutine写满了channel），往下执行，到了写入channel
				的时候，会堵塞住。别的goroutine也会相继阻塞，进而造成所有的goroutine都阻塞 deadlock
			*/
			cond.Wait() // 挂起当前协程，等待条件变量满足，被消费者唤醒
		}
		num := rand.Intn(1000) // 产生一个随机数
		in <- num              // 写入到channel中(生产)
		fmt.Printf("%dth 生产者，产生数据 %3d, 公共区剩余%d个数据\n", idx, num, len(in))
		cond.L.Unlock()         // 生产结束，解锁互斥锁
		cond.Signal()           // 唤醒阻塞的消费者
		time.Sleep(time.Second) //生产完休息一会，给其他的go程执行机会
	}
	wg.Done()
}

// 消费者
func consumer(out <-chan int, idx int) {
	for {
		cond.L.Lock()       // 条件变量对应互斥锁加锁(与生产者是同一个)
		for len(out) == 0 { // 产品区为空，等待生产者生产
			cond.Wait() // 挂起当前协程，等待条件变量满足，被生产者唤醒
		}
		num := <-out
		fmt.Printf("------%dth消费者，消费数据 %3d，公共区剩余%d个数据\n", idx, num, len(out))
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(500 * time.Millisecond)
	}
	wg.Done()
}

func TestCond(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	product := make(chan int, 3)
	cond.L = new(sync.Mutex)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go producer(product, i+1)
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go consumer(product, i+1)
	}
	wg.Wait()
}
