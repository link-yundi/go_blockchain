package timer

import (
	"fmt"
	"testing"
	"time"
)

// time.NewTimer 定时器，时间完成后会往其中的成员变量C写入系统的当前时间（对chan的写操作）
func TestTimer(t *testing.T) {
	fmt.Println("当前时间:", time.Now())
	// 定义一个2s的定时器，2s后系统会往其中写入最新的时间
	timer := time.NewTimer(2 * time.Second)
	nowTimer := <-timer.C
	fmt.Println("现在时间:", nowTimer)
}

func TestAfter(t *testing.T) {
	// time.After  return channel
	fmt.Println("当前时间:", time.Now())
	nowTime := <-time.After(2 * time.Second)
	fmt.Println("现在时间:", nowTime)
}

// 定时器的重置和停止
//var wg = sync.WaitGroup{}

func TestStop(t *testing.T) {
	timer := time.NewTimer(10 * time.Second)
	timer.Reset(1 * time.Second) // 重设定时器时间为1s
	//wg.Add(1)
	go func() {
		<-timer.C
		fmt.Println("子go程，定时完毕")
		//wg.Done()
	}()
	// timer.Stop() // 设置定时器停止
	//wg.Wait()
	for {
		foo()
	}
}

func foo() {

}

// 定时器周期定时

func TestTicker(t *testing.T) {
	quit := make(chan bool) // 创建一个判断是否终止的channel
	fmt.Println("当前时间:", time.Now())
	myTicker := time.NewTicker(time.Second)
	i := 0
	go func() {
		for {
			nowTime := <-myTicker.C
			i++
			fmt.Println("现在时间:", nowTime)
			if i == 8 {
				quit <- true // 接触主go程阻塞
			}
		}
	}()
	<-quit
}
