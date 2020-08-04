package _select

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				quit <- true
				//break
				runtime.Goexit()
			}
		}
	}()
	for i := 0; i < 2; i++ {
		c <- i
		time.Sleep(4 * time.Second)
	}
	<-quit //主go程 阻塞等待 子go程写入数据
	fmt.Println("finish!")
}
