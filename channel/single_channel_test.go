package channel

import (
	"testing"
)

// 单向channel 默认的channel 都是双向的
// 单向写channel var sendCh chan <- int 表示只能往里面写入int类型数据 make(chan <- int)
// 单向读channel var recvCh <- chan int 表示只能从里面读取int类型数据 make(<- chan int)
// 转换: 双向channel 可以隐式转换为任意一种channel    sendCh = ch
// 		单向channel 不能转换为双向channel  ch = sendCh/recvCh error!!!!
// 转换的作用：传参

//func TestSingleChannel(t *testing.T) {
//	ch := make(chan int)
//	// var sendCh chan <- int = ch
//	//sendCh <- 789		// 这种写法是可以的
//	//num := <- sendCh	报错
//	var recvCh <-chan int = ch
//	num := <-recvCh
//	t.Log("num =", num)
//}

func TestSendCh(t *testing.T) {
	// 创建一个双向channel
	ch := make(chan int)
	var sendCh chan<- int = ch
	go func() {
		ch <- 1
		sendCh <- 2
	}()
	chVal := <-ch
	//sendChVal := <- sendCh  报错 不能从单向写channle 中读取数据
	t.Log(chVal)
}

func TestRecvCh(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	var recvCh <-chan int = ch
	val := <-recvCh
	t.Log(val)
}
