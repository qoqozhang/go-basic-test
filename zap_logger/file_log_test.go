package zap_logger

import (
	"testing"
)

var file = RotateFileLog{
	FilePrefix: "user",
	FilePath:   "",
	MaxAge:     3,
}

func TestRotateFileLog_Write(t *testing.T) {
	test := []byte("hello aaaaaaa\n")
	n := 0
	var err error
	for i := 0; i < 100; i++ {
		n, err = file.Write(test)

	}
	if err != nil {
		t.Error("test write fail", err)
	}
	t.Log("write string length:", n)

}

func TestRotateFileLog_Read(t *testing.T) {
	var p = make([]byte, 1024)
	n, err := file.Read(p)
	if err != nil {
		t.Error("test read fail:", err)
		return
	}
	t.Log("read string length:", n)
	t.Log(string(p[:n]))
}
func TestRotateFileLog_ReadLines(t *testing.T) {
	p := file.ReadLines(4, 1000)
	for _, line := range p {
		t.Log(line)
	}
}
