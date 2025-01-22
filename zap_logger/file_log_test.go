package zap_logger

import (
	"testing"
)

func TestRotateFileLog_Write(t *testing.T) {
	file := RotateFileLog{
		FilePrefix: "user",
		FilePath:   "",
		MaxAge:     3,
	}
	test := []byte("hello world\n")
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
