package lock

import (
	"fmt"
	"testing"
	"time"
)

func printer(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}

var syncCh = make(chan int)
var quit = make(chan bool) // 用于通知主go程退出

func person1() {
	printer("hello")
	syncCh <- 1
}

func person2() {
	<-syncCh
	printer("world")
	quit <- true
}

func TestChannel(t *testing.T) {
	go person1()
	go person2()
	<-quit
	close(syncCh)
	close(quit)
} // 在不借助channel的情况下，打印会错乱，通过channel的阻塞可以保证数据的正确输出
