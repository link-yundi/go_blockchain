package channel

import "testing"

func TestChannelWithBuffer(t *testing.T) {
	ch := make(chan int, 4)
	t.Log(len(ch), cap(ch))

	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
			t.Log("子go程写:", i)
			t.Log(len(ch), cap(ch))
		}
	}()
	for i := 0; i < 8; i++ {
		num := <-ch
		t.Log("主go程读:", num)
		t.Log(len(ch), cap(ch))
	}
	// 打印结果会有错乱，因为打印是io操作，需要访问硬件，具有严重的耗时，和goroutine的操作不能做到同步，具有延迟
}
