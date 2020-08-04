### 锁

1. 死锁 deadlock

> 不是锁的一种，是一种错误使用锁导致的现象

```go
// 死锁1 单goroutine自己死锁
func main() {
  ch := make(chan int)
  ch <- 789 // 在同一个进程中读写 在写入的时候，通道阻塞，【写阻塞】等待读取，后面的代码不再执行
  num := <- ch
  fmt.Println("num =", num)
}
// channel 应该至少在2个以上的goroutine中进行通信，否则死锁！
```

```go
// 死锁2 goroutine间channel访问顺序导致死锁
func main() {
  ch := make(chan int)
  num := <- ch // 在这里读取数据的时候阻塞了【读阻塞】，后面的代码不再执行
  fmt.Println("num =", num)
  go func() {
    ch <- 789
  }()
}
// 使用channel一端读（写）,要保证另一端写（读）操作，同时有机会执行。
```

```go
// 死锁3 
func main() {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go func() {
    select {
      case num := <- ch1:
      ch2 <- num
    }
  }()
  for {
    select {
      case num := <- ch2:
      ch1 <- num
    }
  }
}
// 多goroutine，多channel交叉死锁
// 在上述的例子中，ch1需要ch2读到数据才能往里面写数据，ch2需要ch1读到数据才能往里面写数据，但是一开始的时候没有goroutine往ch1和ch2写数据，会阻塞，引发死锁
// A goroutine掌握M的同时，尝试拿N；B goroutine掌握N的同时尝试拿M
// 子goroutine拿到了ch1的读，尝试拿ch2的写，主goroutine拿到ch2的读，同时想拿ch1的写
```

在go中，尽量不要将mutex、rwMutex与channel混用，容易导致死锁（隐性死锁）单独使用channel或者mutex、rwMutex已经足够完成数据同步，根据推荐做法：通过通信来共享内存，而不是通过共享内存来通信。也就是说，建议使用channel

2. 互斥锁 mutex

> 使用channel也可以实现同步

```go
func printer(str string) {
  for _, ch := range str {
    fmt.Printf("%c", ch)
    time.Sleep(300 * time.Millisecond)
  }
}

var syncCh = make(chan int)
var quit = make(chan bool)	// 用于通知主go程退出

func person1() {
  printer("hello")
  syncCh <- 1
}

func person2() {
  <- syncCh
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
```

> 通过互斥锁来完成同步 读 -> 锁 -> 共享数据，写 -> 锁 -> 共享数据，也就是说，对于共享数据的操作，都得拿到锁，拿不到锁的goroutine会阻塞在锁这里

```go
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
	wg.Done()	// done不是在printer完成后加，是为了保证先执行执行person1再执行person2的同时查看锁的效果
	printer("hello")
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
```

3. 读写锁

> 读时共享，写时独占。写锁优先级高于读锁