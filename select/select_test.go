package _select

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// select作用：监听channel上的数据流动
// 每个case语句里面必须是一个io操作：case <- chan1...
// case <- chan1 : 如果chan1成功读到数据，则进行该case处理语句
// case <- chan2 <- 1: 如果成功向chan2写入数据，则进行该case处理语句
// 注意：如果多个case同时满足，那么就会从那些可以执行的语句中任意选择一条来使用
// 如果给出了default语句，那么在不满足所有case的情况下执行default，同时程序的执行会从select语句后的语句中恢复。可能会导致忙轮询，系统消耗资源比较大
// 如果没有default，那么select语句将会被阻塞，直到至少有一个通信可以进行下去，让出时间片，提高程序效率

func TestSelect(t *testing.T) {
	ch := make(chan int)    // 用来进行数据通信的channel
	quit := make(chan bool) // 用来判断是否退出的channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true
		runtime.Goexit()
	}()
	// 主go读数据
	for {
		select {
		case num, ok := <-ch:
			if ok {
				fmt.Println("读到:", num)
			}
		case <-quit:
			//break // break跳出select 没有跳出for loop
			return
		}
		fmt.Println("==========================")
	}
}

/*
注意事项:
1. 监听的case中，没有满足监听条件，阻塞
2. 监听的case中，有多个满足监听条件，任选一个执行
3. 可以使用default来处理所有case都不满足监听条件的状况。通常不用，因为会产生忙轮询
4. select 自身不带有循环机制，需要借助外层for来循环监听
5. break只能跳出select,类似switch中的用法
*/
