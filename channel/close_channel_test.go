package channel

import (
	"testing"
	"time"
)

// 关闭channel

func TestCloseChannel(t *testing.T) {
	// 关闭语法：close(channel)
	// 数据没有读取完，不应该关闭channel
	// 关闭的channel 写入数据会引发错误，但是可以读取数据，读取的零值（默认值），也就是说，close(channel)关闭的是写端
	// value, ok := <- channel
	// ok 为false channel已经关闭，true 则channel没有关闭
	ch := make(chan int)
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}
		close(ch)
	}()
	for {
		if num, ok := <-ch; ok {
			t.Log("读到数据,", num)
		} else {
			break
		}
	}
	// 读取无缓冲channel: 读到0
	// 读取有缓冲channel: 如果缓冲区内有数据，先读数据，读完数据后，可以继续读，读到0
}

func TestCloseChannelWithBuffer(t *testing.T) {

	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		t.Log("通道关闭")
	}()
	time.Sleep(5 * time.Second)
	for {
		if num, ok := <-ch; ok {
			t.Log("读到数据,", num)
		} else {
			t.Log("关闭后:", num)
			break
		}
		// 可以看出，读端是通过读取零值来判断channel是否关闭的
	}
}
