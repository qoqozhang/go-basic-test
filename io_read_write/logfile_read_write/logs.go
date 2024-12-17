package logfile_read_write

import (
	"bufio"
	"io"
	"os"
)

type LogFileReadWriter interface {
	io.Closer
	io.Writer
	io.Reader
	ReadLines(int, int) ([]string, error)
	Sync() error
}
type LogFileReadWrite struct {
	src     string
	logFile *os.File
}

// NewLogFileReadWrite 创建一个可读写的日志文件对象
func NewLogFileReadWrite(src string) LogFileReadWriter {
	file, _ := os.OpenFile(src, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	return &LogFileReadWrite{src: src, logFile: file}
}

func (l *LogFileReadWrite) Close() error {
	return l.logFile.Close()
}

func (l *LogFileReadWrite) Read(p []byte) (n int, err error) {
	return l.logFile.Read(p)
}

func (l *LogFileReadWrite) Write(p []byte) (n int, err error) {
	return l.logFile.Write(p)
}

// ReadLines 从offset 指定位置读取 size 行
func (l *LogFileReadWrite) ReadLines(offset, size int) ([]string, error) {
	scanner := bufio.NewScanner(l.logFile)
	currentLine := 1
	lastLine := offset + size
	readLines := make([]string, 0)
	for scanner.Scan() {
		if currentLine >= offset && currentLine < lastLine {
			readLines = append(readLines, scanner.Text())
		}
		currentLine++
	}
	return readLines, nil
}

// Truncate 清空文件
func (l *LogFileReadWrite) Truncate(size int64) error {
	return l.logFile.Truncate(size)
}

// Sync 从内存写入文件
func (l *LogFileReadWrite) Sync() error {
	return l.logFile.Sync()
}
