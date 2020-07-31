package _struct

import (
	"fmt"
	"testing"
)

type Person struct {
	Name   string
	Gender string
	Age    int
}

func test(man Person) {
	man.Name = "Golang"
	man.Age = 14
}

func TestStruct(t *testing.T) {
	p := Person{
		Name:   "ZhangYundi",
		Gender: "Male",
		Age:    30,
	}
	p1 := Person{
		Name:   "YinJie",
		Gender: "Female",
		Age:    25,
	}
	p = p1
	p1.Name = "YinJie"
	//fmt.Println(p == p1)
	t.Logf("%v %v\n", p, p1)
	test(p) // 值传递
	t.Log(p)
}

type Person3 struct {
	name     string
	age      int
	flag     bool
	interest []string
}

func initFunc(p *Person3) {
	p.name = "ZhangYundi"
	p.age = 30
	p.flag = true
	p.interest = []string{"Python", "Golang"}
}

func TestInit(t *testing.T) {
	var p1 Person3
	initFunc(&p1)
	fmt.Println(p1)
	p2 := new(Person3)
	initFunc(p2)
	fmt.Println(p2)
}
