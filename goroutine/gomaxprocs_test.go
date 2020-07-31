package goroutine

import (
	"runtime"
	"testing"
)

func TestGOMAXPROCS(t *testing.T) {
	//n := runtime.GOMAXPROCS(1)
	//fmt.Println(n)
	//for {
	//	go fmt.Print(0)
	//	fmt.Print(1)
	numOfCPU := runtime.NumCPU()
	t.Log(numOfCPU)
}
