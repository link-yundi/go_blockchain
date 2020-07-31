package Map

import (
	"fmt"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	//var m1 map[string]int // 只是声明，还没有在内存中开辟空间，直接使用会报错
	m1 := make(map[string]int) // 要估算好map的容量，避免在程序运行期间再动态扩容
	//m2 := make()
	m1["理想"] = 18
	m1["jiwuming"] = 35
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	v, ok := m1["娜扎"]
	fmt.Println(v, ok)
}

func TestMapLiter(t *testing.T) {
	m := make(map[string]int)
	m["ZhangYundi"] = 30
	m["yinjie"] = 25
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k1 := range m {
		fmt.Println(k1)
	}
	// 只遍历value
	for _, v1 := range m {
		fmt.Println(v1)
	}
	// 删除
	delete(m, "yinjie")
	fmt.Println(m)
	// 删除不存在的key
	delete(m, "娜扎")
	fmt.Println(m)
}

// 封装一个wcFunc(),接受一段英文字符串str,返回一个map，统计每个词出现的次数
func wcFunc(input string) map[string]int {
	// 1. string 切割
	//stringSlice := strings.Split(input, " ")
	stringSlice := strings.Fields(input)
	//fmt.Println(strings.Fields(input))
	ret := make(map[string]int)
	for _, elem := range stringSlice {
		if _, ok := ret[elem]; !ok {
			ret[elem] = 1
		} else {
			ret[elem] += 1
		}
	}
	return ret
}

func Test(t *testing.T) {
	str := "I love my work and I love my family too"
	ret := wcFunc(str)
	t.Log(ret)
}
