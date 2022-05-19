package tools

import "os"

// 判断是否是目录
func IsDir(filename string) bool {
	return isFileOrDir(filename, true)
}

// 判断是文件还是目录，根据decideDir为true表示判断是否为目录；否则判断是否为文件
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

//创建目录
func CreateDir(path string) bool {
	if IsDirOrFileExist(path) == false {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

//判断文件 或 目录是否存在
func IsDirOrFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}