package helper

import "os"

func CreateFile(name string, content ...string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, _ := os.Create(name)
		if len(content) == 1 {
			file.WriteString(content[0])
		}
		file.Close()
		SuccessPuts("create", "%s", name)
	} else {
		ErrorPuts("create", "%s 文件已存在 ", name)
	}
}
