package file

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	file, err := os.Create("test_file.xyz")
	defer file.Close()
	if err != nil {
		t.Fatal("create error:", err)
	}
	t.Log("Successful!!")
}

func TestOpen(t *testing.T) {
	file, err := os.Open("test_file.xyz")
	defer file.Close()
	if err != nil {
		t.Fatal("create error:", err)
	}
	_, err = file.WriteString("##########")
	if err != nil {
		t.Fatal("create error:", err)
	}
	t.Log("Successful!!")
}

func TestOpenFile(t *testing.T) {
	file, err := os.OpenFile("test_file.xyz", os.O_RDWR, 6)
	defer file.Close()
	if err != nil {
		t.Fatal("create error:", err)
	}
	_, err = file.WriteString("##########")
	if err != nil {
		t.Fatal("create error:", err)
	}
	t.Log("Successful!!")
}

func Test(t *testing.T) {
	a, b, c, d, e := os.O_WRONLY, os.O_RDWR, os.O_CREATE, os.O_APPEND, os.O_RDONLY
	//t.Logf("%T \n", a)
	t.Log(a, b, c, d, e)
	t.Log(a | b | c)
}
