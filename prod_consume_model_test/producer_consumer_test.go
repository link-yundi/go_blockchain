package prod_consume_model_test

import (
	"fmt"
	"testing"
)

// 生产者消费者模型
// 生产者（发送数据端） -> 缓冲区 -> 消费者（接受数据端）
// 引入缓冲区可以  1.解耦
//				2.提高并发能力（生产者和消费者数量不对等时，能保持正常通信）
//				3.缓存（生产者和消费者处理速度不一致时，暂存数据）
// 无缓冲channel: 同步通信
// 有缓冲channel: 异步通信，并发高 两者的选择看具体的业务需求

func Producer(in chan<- int) {
	// 消息生产者，写入数据
	for i := 0; i < 10; i++ {
		in <- i * i
	}
	close(in)
}

func Consumer(out <-chan int) {
	// 消费者，读数据
	for num := range out {
		fmt.Println("消费者拿到:", num)
	}
}

func TestProdConsumer(t *testing.T) {
	ch := make(chan int)
	go Producer(ch)
	Consumer(ch)
}
