package logfile_read_write

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var wg = new(sync.WaitGroup)

func BenchmarkNewLogFileReadWrite(b *testing.B) {
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		logfile := NewLogFileReadWrite("./test.txt")
		logfile.ReadLines(1000, 3)
		go func() {
			logfile.Close()
			defer wg.Done()
		}()
	}
	wg.Wait()
	runtime.ReadMemStats(&m2)
	fmt.Println("total: ", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Println("mallocs: ", m2.Mallocs-m1.Mallocs)
}

func BenchmarkLogFileReadWrite_ReadLines(b *testing.B) {
	logfile := NewLogFileReadWrite("./test.txt")
	defer logfile.Close()
	var file []byte = make([]byte, 1024)
	n, _ := logfile.Read(file)
	for i := 0; i < b.N; i++ {
		if _, err := logfile.Write(file[:n]); err != nil {
			b.Error(err)
		}
	}
	logfile.Sync()

}
