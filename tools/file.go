package tools

import (
	"fmt"
	"io"
	"log"
	"os"
)

// IsDir 判断是否是目录
func IsDir(filename string) bool {
	return isFileOrDir(filename, true)
}

// isFileOrDir 判断是文件还是目录，根据decideDir为true表示判断是否为目录；否则判断是否为文件
func isFileOrDir(filename string, decideDir bool) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	isDir := fileInfo.IsDir()
	if decideDir {
		return isDir
	}
	return !isDir
}

// CreateDir 创建目录
func CreateDir(path string) bool {
	if IsDirOrFileExist(path) == false {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

// 创建目录
func CreateDirs(path string) string {
	if IsDirOrFileExist(path) == false {
		b := CreateDir(path)
		if !b {
			log.Fatalf("Directory created failed>>%s\n", path)
			return ""
		}
		fmt.Printf("Directory created success:%s\n", path)
	}
	return path
}

// IsDirOrFileExist 判断文件 或 目录是否存在
func IsDirOrFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// WriteFileAppend 追加写文件
func WriteFileAppend(filename string, data string) (count int, err error) {
	var f *os.File
	if IsDirOrFileExist(filename) == false {
		f, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
	}
	defer f.Close()
	count, err = io.WriteString(f, data)
	if err != nil {
		return
	}
	return
}

// WriteAppendFile 追加写文件
func WriteAppendFile(path, data string) (err error) {
	if _, err := WriteFileAppend(path, data); err == nil {
		fmt.Printf("Generate success:%s\n", path)
		return nil
	} else {
		return err
	}
}
