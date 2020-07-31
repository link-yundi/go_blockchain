package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func test() {
	defer func() {
		fmt.Println("cccccccccccccccc")
		//wg.Done()
	}()
	runtime.Goexit()
	fmt.Println("dddddddddddddddd")
}

func Test(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("aaaaaaaaaaaaaaa")
		test()
		//wg.Done()
		fmt.Println("bbbbbbbbbbbbbbb")
	}()
	wg.Wait()
}
