package PointerMakeNew

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	//n := 18
	//p := &n
	//fmt.Println(p)
	//fmt.Printf("%T\n", p)
	//// * 根据地址取值
	//m := *p
	//fmt.Println(m)
	//fmt.Printf("%T\n", m)
	a := new(int)
	fmt.Println(a)
	*a = 100
	fmt.Println(*a)
}

func TestPointer2(t *testing.T) {
	var a int = 10
	var p *int = &a	// 指针变量: 存储地址的变量
	a = 100
	t.Log("a =", a)
	// * 解引用、间接引用
	*p = 250
	t.Log("a =", a)
	t.Log("*p =", *p)
}

func swap1(a, b int) {
	a, b = b, a
	fmt.Println("swap1 a:", a, "b:", b)
}
func swap2(a, b *int) {
	*a, *b = *b, *a
	fmt.Println("swap2 a:", *a, "b:", *b)
}
func main(){
	a, b := 10, 20
	swap1(a, b)
	fmt.Println("main-swap1 a:", a, "b:", b)
	swap2(&a, &b)
	fmt.Println("main-swap2 a:", a, "b:", b)
}

func Test1(t *testing.T) {
	main()
}






