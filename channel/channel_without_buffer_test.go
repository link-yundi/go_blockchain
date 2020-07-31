package channel

import "testing"

func TestChannelWithoutBuffer(t *testing.T) {
	// 无缓冲channel是没有存储数据的功能的，数据不能再其中停留，只能作为传递中介
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			t.Log("子go程写:", i)
		}
	}()
	for i := 0; i < 5; i++ {
		num := <-ch
		t.Log("主go程读:", num)
	}

}
