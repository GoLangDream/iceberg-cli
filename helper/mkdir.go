package helper

import (
	"os"
)

func MkDir(name string) {
	if _, err := os.Stat(name); os.IsExist(err) {
		os.MkdirAll(name, os.ModeDir)
		SuccessPuts("create", "创建 %s", name)
	} else {
		ErrorPuts("create", "%s 文件已存在 ", name)
	}
}
