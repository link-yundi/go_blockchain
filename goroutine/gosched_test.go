package goroutine

import (
	"runtime"
	"testing"
)

func TestGosched(t *testing.T) {
	go func() {
		for {
			t.Log("this is goroutine test")
			//time.Sleep(100 * time.Millisecond)
		}
	}()
	for {
		runtime.Gosched() // 出让当前CPU时间片
		t.Log("this is main test...")
		//time.Sleep(100 * time.Millisecond)
	}
}
