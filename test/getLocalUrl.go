package test

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func Test() string {
	fmt.Println("File:", getCurrentFile())
	fmt.Println("Line:", getCurrentLine())

	//logFilePath := filepath.Join(parentDir, "path", "to", "log", "file.log")
	fmt.Println("父级别路径%v:", getParentPath())

	return "ok"
}

//获得当前文件url的路径
func getCurrentFile() string {
	_, file, _, _ := runtime.Caller(1)
	absPath, _ := filepath.Abs(file)
	return absPath
}

//获得当前文件所在行数
func getCurrentLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

//获得当前父目录的绝对路径
func getParentPath() string {
	_, file, _, _ := runtime.Caller(1)
	absPath, _ := filepath.Abs(filepath.Dir(file))
	return absPath
}

//构造完整的文件路径
func getFilePath(path string) string {
	//logFilePath := filepath.Join(parentDir, "path", "to", "log", "file.log")
	return filepath.Join(getParentPath(), "path", "to", "log", "file.log")
}

//获得当前项目的父目录
func parentDir() string {
	_, file, _, _ := runtime.Caller(1)
	absPath, _ := filepath.Abs(filepath.Dir(filepath.Dir(file)))
	return absPath
}
