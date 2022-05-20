package tools

import (
	"LyGen/constant"
	"os"
	"os/exec"
	"runtime"
)

// Gofmt FMT 格式代码
func Gofmt(path string) bool {
	if IsDirOrFileExist(path) {
		if !ExecCommand("goimports", "-l", "-w", path) {
			if !ExecCommand("gofmt", "-l", "-w", path) {
				return ExecCommand("go", "fmt", path)
			}
		}
		return true
	}
	return false
}

// ExecCommand 执行命令
func ExecCommand(name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	return true
}

// Clean 清屏
func Clean() {
	switch GetOs() {
	case constant.Darwin, constant.Linux:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	case constant.Window:
		cmd := exec.Command("main", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

// GetOs 获取操作系统
func GetOs() int {
	switch runtime.GOOS {
	case "darwin":
		return constant.Darwin
	case "windows":
		return constant.Window
	case "linux":
		return constant.Linux
	default:
		return constant.Unknown
	}
}