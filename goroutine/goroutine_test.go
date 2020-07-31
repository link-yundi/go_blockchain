package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func sing() {
	for i := 0; i < 5; i++ {
		// wg.Add(1)
		fmt.Println("-------正在唱歌：隔壁泰山--------")
		time.Sleep(100 * time.Millisecond)
		//wg.Done()
	}
}

func dance() {
	for i := 0; i < 5; i++ {
		// wg.Add(1)
		fmt.Println("-------正在跳舞：赵四街舞--------")
		time.Sleep(100 * time.Millisecond)
		//wg.Done()
	}
}

func TestFuncs(t *testing.T) {
	sing()
	dance()
}

func TestGoroutine(t *testing.T) {
	wg.Add(10)
	go sing()
	go dance()
	wg.Wait()
}

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func TestTask(t *testing.T) {
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
