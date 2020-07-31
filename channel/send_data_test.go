package channel

import "testing"

func TestSendData(t *testing.T) {
	ch := make(chan string)
	//t.Log("len(ch) =", len(ch))
	go func() {
		for i := 0; i < 2; i++ {
			t.Log("i =", i)
		}
		ch <- "子go程打印完毕"
	}()
	str := <- ch
	t.Log(str)
}

