package helper

import (
	"os"
)

func MkDir(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		os.MkdirAll(name, os.ModePerm)
		SuccessPuts("create", "%s", name)
	} else {
		ErrorPuts("create", "%s 文件已存在 ", name)
	}
}
