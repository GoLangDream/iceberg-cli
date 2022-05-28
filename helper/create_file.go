package helper

import "os"

func CreateFile(name string, content ...string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, _ := os.Create(name)
		if len(content) == 1 {
			file.WriteString(content[0])
		}
		file.Close()
		SuccessPuts("创建文件成功", "%s", name)
	} else {
		ErrorPuts("创建文件失败", "%s 文件已存在 ", name)
	}
}
