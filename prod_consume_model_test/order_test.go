package prod_consume_model

import (
	"fmt"
	"testing"
)

// 模拟订单系统
type OrderInfo struct {
	id int
}

func producer(in chan<- OrderInfo) {
	for i := 0; i < 10; i++ {
		order := OrderInfo{i + 1}
		in <- order
	}
	close(in)
}

func consumer(out <-chan OrderInfo) {
	for order := range out {
		fmt.Println("订单id为:", order.id)
	}
}

func TestOrder(t *testing.T) {
	buffer := make(chan OrderInfo)
	go producer(buffer)
	consumer(buffer)
}
