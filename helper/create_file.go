package helper

import (
	"github.com/flosch/pongo2/v5"
	"os"
)

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

func CreateFileFromTmpl(tmplContent, targetPath string, ctx pongo2.Context) {
	tmpl, err := pongo2.FromString(tmplContent)
	if err != nil {
		ErrorPuts("create", "%s\n            %s\n", targetPath, err.Error())
	}
	fileContent, err := tmpl.Execute(ctx)
	if err != nil {
		ErrorPuts("create", "%s\n            %s\n", targetPath, err.Error())
	}
	CreateFile(targetPath, fileContent)
}
