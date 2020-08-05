package main

import "fmt"

func main() {
	var a int
	var b int
	a = 1
	b = 2
	ret := add_a_and_b(a, b)
	fmt.Println(ret)
}

func add_a_and_b(a, b int) int {
	ret := a + b
	return ret
}
