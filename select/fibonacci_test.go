package _select

import (
	"fmt"
	"runtime"
	"testing"
)

func fibonacci(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Print(num, " ")
		case <-quit:
			runtime.Goexit()
		}
	}
}

func TestFibonacci(t *testing.T) {
	ch := make(chan int)
	quit := make(chan bool)
	go fibonacci(ch, quit) // 子go程，打印fibonacci数列
	x, y := 1, 1
	for i := 0; i < 20; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}
