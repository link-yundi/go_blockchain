package file

import (
	"bufio"
	"github.com/pkg/errors"
	"io"
	"os"
	"testing"
)

var fileName string = "test_file.xyz"

func TestRead(t *testing.T) {
	//file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND｜os.CREATE, 6)
	file, errOpen := os.OpenFile(fileName, os.O_RDWR, 6)
	if errOpen != nil {
		errOpen = errors.Wrap(errOpen, "File Open Failed")
		t.Fatal("Open file error:", errOpen)
	}
	defer file.Close()
	n, errWrite := file.WriteString("HelloWorld!\n")
	if errWrite != nil {
		t.Fatal("写入异常:", errWrite)
	}
	t.Log(n)
	off, _ := file.Seek(-5, io.SeekEnd)
	t.Log("off:", off)
	//t.Log("write end!")
	n, _ = file.WriteAt([]byte("1111"), off)
	t.Log("WriteAt:", n)
}

func TestReadRow(t *testing.T) {
	file, errOpen := os.OpenFile(fileName, os.O_RDONLY, 4)
	t.Log(fileName)
	if errOpen != nil {
		errOpen = errors.Wrap(errOpen, "File Open Failed")
		t.Fatal("Open file error:", errOpen)
	}
	defer file.Close()
	// 按行读取文件
	// 创建一个带有缓冲区的Reader
	reader := bufio.NewReader(file)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Cause(err) == io.EOF {
				t.Log("文件读取完毕！")
				return
			} else {
				t.Error("ReadBytes Error.")
			}
		}
		t.Log(string(buf))
	}
}
