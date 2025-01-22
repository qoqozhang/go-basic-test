package zap_logger

import (
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
)

type RotateFileLog struct {
	FilePrefix string
	FilePath   string
	MaxAge     int
	file       *os.File
	mu         sync.Mutex
	currentDay string
}

func (log *RotateFileLog) Rotate() error {
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.rotate()
}

func (log *RotateFileLog) Close() error {
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.file.Close()
}

func (log *RotateFileLog) Write(p []byte) (n int, err error) {
	//获取锁操作
	log.mu.Lock()
	defer log.mu.Unlock()

	// 如果是首次使用，则需要打开日志文件操作符
	if log.file == nil {
		if err := log.openExistingOrNew(); err != nil {
			return 0, err
		}
	}

	// 判断当前打开的日志文件是否是今天的日志文件，不是的话则轮转日志文件
	currentDay := time.Now().Format("2006-01-02")
	if log.currentDay != currentDay {
		err = log.rotate()
	}
	if err != nil {
		return 0, err
	}

	// 写入日志到文件
	n, err = log.file.Write(p)
	return n, err
}

// deleteMaxAgeBeforeLog 删除保留保留日志之前的日志
func (log *RotateFileLog) deleteMaxAgeBeforeLog() error {
	filePath, err := log.getAbsoluteFilePath()
	if err != nil {
		return err
	}
	// 遍历日志目录下面的所有文件，并判断文件名称格式是不是 xxxxx_2016-01-02的格式，如果是的话则删除
	err = filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		IsLogFile, _ := regexp.MatchString(log.FilePrefix+"_"+`\d{4}-\d{2}-\d{2}`, info.Name())
		if IsLogFile {
			modifiedTime := info.ModTime()
			diff := time.Now().Sub(modifiedTime)
			if diff > time.Duration(log.MaxAge*24)*time.Hour {
				err := os.Remove(path)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// rotate 关闭之前打开的文件操作符，并创建一个新的文件操作符
func (log *RotateFileLog) rotate() error {
	if err := log.file.Close(); err != nil {
		return err
	}
	logfile, err := log.GetCurrentLogName()
	if err != nil {
		return err
	}
	return log.openNew(logfile)
}

// openNew 打开一个新的文件
func (log *RotateFileLog) openNew(filename string) error {
	f, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	log.file = f
	currentDay := time.Now().Format("2006-01-02")
	log.currentDay = currentDay
	return nil
}

// openExistingOrNew 尝试打开文件，不存在则创建一个新的文件
func (log *RotateFileLog) openExistingOrNew() error {
	logfile, err := log.GetCurrentLogName()
	if err != nil {
		return nil
	}
	_, err = os.Stat(logfile)
	if os.IsNotExist(err) {
		return log.openNew(logfile)
	}
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	log.file = file
	currentDay := time.Now().Format("2006-01-02")
	log.currentDay = currentDay
	return nil
}

// GetAbsoluteFilePath  获取文件的绝对路径
// 1. 判断日志文件夹是 ""，则返回执行程序的当前路径下的 logs/路径,不存在则创建
// 2. 如果日志文件夹路径不为空，则判断是否存在，不存在则创建，创建失败则返回error
func (log *RotateFileLog) getAbsoluteFilePath() (fileFullName string, err error) {
	var filePath string
	filePath, err = os.Getwd()
	if err != nil {
		return "", err
	}
	if log.FilePath == "" {
		fileFullName = filePath + "/logs/"
	}
	// 判断是否存在，不存在则尝试创建
	if _, err := os.Stat(fileFullName); os.IsNotExist(err) {
		err := os.Mkdir(fileFullName, 0777)
		if err != nil {
			return "", err
		}
	}
	return fileFullName, nil
}

// GetCurrentLogName 获取当前的日志文件绝对路径
func (log *RotateFileLog) GetCurrentLogName() (logfile string, err error) {
	currentDay := time.Now().Format("2006-01-02")
	filepath, err := log.getAbsoluteFilePath()
	if err != nil {
		return "", err
	}
	logfile = filepath + "/" + log.FilePrefix + "_" + currentDay + ".log"
	return logfile, nil
}
