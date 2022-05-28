package helper

import (
	"os"
)

func MkDir(name string) {
	if _, err := os.Stat(name); os.IsExist(err) {
		os.MkdirAll(name, os.ModeDir)
		SuccessPuts("创建目录成功", "创建 %s", name)
	} else {
		ErrorPuts("创建目录失败", "%s 文件已存在 ", name)
	}
}
