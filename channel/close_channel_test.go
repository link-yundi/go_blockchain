package channel

import "testing"

// 关闭channel

func TestCloseChannel(t *testing.T) {
	// 关闭语法：close(channel)
	// 数据没有读取完，不应该关闭channel
	// 关闭的channel 写入数据会引发错误，但是可以读取数据，读取的零值（默认值）
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
}
