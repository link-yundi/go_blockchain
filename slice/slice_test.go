package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", a)
}

func TestAppend(t *testing.T) {
	s := make([]int, 3)
	s = append(s, 1, 2, 3)
	t.Log(s)
}

func TestSlice1(t *testing.T) {
	// 去除空元素
	s := []string{"red", "", "black", "", "", "pink", "blue"}
	var sNew []string
	for _, elem := range s {
		if elem != "" {
			sNew = append(sNew, elem)
		}
	}
	t.Log(sNew)
}

func Duplication1(sliceRaw []string) []string {
	sliceNew := []string{sliceRaw[0]}
	skip := false // skip flag：一旦存在同样的元素，则设置为true，表示跳过该元素，
	for _, elemOut := range sliceRaw {
		skip = false
		for _, elemIn := range sliceNew {
			if elemOut == elemIn {
				skip = true
				break
			}
		}
		if !skip {
			sliceNew = append(sliceNew, elemOut)
		}
	}
	return sliceNew
}

func Duplication2(sliceRaw []string) []string {
	sliceNew := []string{sliceRaw[0]}
	//skip := false		// skip flag：一旦存在同样的元素，则设置为true，表示跳过该元素，
	for _, elemOut := range sliceRaw {
		//skip = false
		i := 0
		for ; i < len(sliceNew); i++ {
			if elemOut == sliceNew[i] {
				//skip = true
				break
			}
		}
		if i == len(sliceNew) {
			sliceNew = append(sliceNew, elemOut)
		}
	}
	return sliceNew
}

func TestDuplication(t *testing.T) {
	sliceRaw := []string{"red", "red", "black", "pink", "blue", "pink", "blue"}
	sliceNew1 := Duplication1(sliceRaw)
	sliceNew2 := Duplication2(sliceRaw)
	t.Log(sliceNew1)
	t.Log(sliceNew2)
}

func BenchmarkDuplication1(b *testing.B) {
	sliceRaw := []string{"red", "red", "black", "pink", "blue", "pink", "blue"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Duplication1(sliceRaw)
	}
	b.StopTimer()
}

func BenchmarkDuplication2(b *testing.B) {
	sliceRaw := []string{"red", "red", "black", "pink", "blue", "pink", "blue"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Duplication2(sliceRaw)
	}
	b.StopTimer()
}

func TestCopy(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[8:] // {8, 9}
	s2 := data[:5] // {0, 1, 2, 3, 4}
	t.Log(s2)
	copy(s2, s1) // {8, 9, 2, 3, 4}
	t.Log(s2)
}

// 练习: 删除slice中某个元素，并且保存原有的元素位置
// {5, 6, 7, 8, 9} -> {5, 6, 8, 9}
func removeSliceElem(data []int, idx int) []int {
	dst := data[idx:]   // {7, 8, 9}
	src := data[idx+1:] // {8, 9}
	copy(dst, src)      // {8, 9, 9}
	return data[:len(data)-1]
}

func Test(t *testing.T) {
	data := []int{5, 6, 7, 8, 9}
	out := removeSliceElem(data, 2)
	t.Log(out)
}
