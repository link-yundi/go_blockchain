package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 需要记住：channel 一定要和go程一起出现，一个go程负责读，一个负责写，两端同时操作，不然会报错

var wg sync.WaitGroup

func printer(str string) {
	for _, s := range str {
		//fmt.Print(1)
		fmt.Printf("%c", s)
		time.Sleep(300 * time.Millisecond)
	}
}

func person1() {
	printer("hello")
	wg.Done()
}

func person2() {
	printer("world")
	wg.Done()
}

func TestChannel(t *testing.T) {
	wg.Add(2)
	go person1()
	go person2()
	wg.Wait()
}
