package mutex

import (
	"sync"
	"testing"
	"time"
)

var mutex sync.Mutex // 创建一个互斥锁，新建的互斥锁状态为0，未加锁。锁只有一把。
var get []int32

func printer(str string) {
	mutex.Lock() // 访问共享数据之前，加锁
	for _, ch := range str {
		get = append(get, ch)
		time.Sleep(300 * time.Millisecond)
	}
	mutex.Unlock() // 共享数据访问结束，解锁
}

func person1() {
	wg.Done()
	printer("hello")
	//wg.Done()
}

func person2() {
	printer("world")
	wg.Done()
}

func TestMutex(t *testing.T) {
	want := "helloworld"
	wg.Add(1)
	go person1()
	wg.Wait()
	wg.Add(1)
	go person2()
	wg.Wait()
	if want != string(get) {
		t.Errorf("expect %s, get %s", want, string(get))
	}
}
