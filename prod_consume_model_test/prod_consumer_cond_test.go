package prod_consume_model

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 通过条件变量实现生产者消费者模型
var (
	cond    sync.Cond
	wg1     sync.WaitGroup
	reMutex sync.RWMutex
)

// 生产者
func producer2(in chan<- int, idx int) {
	for {
		cond.L.Lock()      // 条件变量对应互斥锁加锁
		for len(in) == 3 { // 产品区满，等待消费者消费
			/*
				采用for循环而不是if判断
				假设当前go程 len(in) == 2，往后面执行，写入数据
			*/
			cond.Wait() // 挂起当前协程，等待条件变量满足，被消费者唤醒
		}
		num := rand.Intn(1000) // 产生一个随机数
		in <- num              // 写入到channel中(生产)
		fmt.Printf("%dth 生产者，产生数据 %3d, 公共区剩余%d个数据\n", idx, num, len(in))
		cond.L.Unlock() // 生产结束，解锁互斥锁
		cond.Signal()   // 唤醒阻塞的消费者
		time.Sleep(time.Second)
	}
}

// 消费者
func consumer2(out <-chan int, idx int) {
	for {
		cond.L.Lock()      // 条件变量对应互斥锁加锁(与生产者是同一个)
		if len(out) == 0 { // 产品区为空，等待生产者生产
			cond.Wait() // 挂起当前协程，等待条件变量满足，被生产者唤醒
		}
	}
}
