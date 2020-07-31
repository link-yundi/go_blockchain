package string

import (
	"strings"
	"testing"
)

var str string = "aa.txt bb.pdf"

// split
func TestSplit(t *testing.T) {
	ret := strings.Split(str, " ")
	t.Log(ret)
}

// fields，只能按照空格进行拆分
func TestFields(t *testing.T) {
	ret := strings.Fields(str)
	t.Log(ret)
}

// HasSuffix,判断后缀

func TestSuffix(t *testing.T) {
	ret := strings.Fields(str)
	t.Log(strings.HasSuffix(ret[0], "txt"), strings.HasSuffix(ret[1], "txt"))
	t.Log(strings.HasSuffix(ret[0], "pdf"), strings.HasSuffix(ret[1], "pdf"))
}

// HasPrefix, 判断前缀
func TestPrefix(t *testing.T) {
	ret := strings.Fields(str)
	t.Log(strings.HasPrefix(ret[0], "aa"), strings.HasPrefix(ret[1], "aa"))
	t.Log(strings.HasPrefix(ret[0], "bb"), strings.HasPrefix(ret[1], "bb"))
}
